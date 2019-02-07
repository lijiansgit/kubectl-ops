package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"

	appsv1 "k8s.io/api/apps/v1"

	"github.com/lijiansgit/go/libs/consul"
	autov1 "k8s.io/api/autoscaling/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/util/intstr"
)

const (
	defaultConsulPath    = "kubernetes/v1"
	dockerFileName       = "Dockerfile"
	defaultAppBuildCmd   = "make"
	defaultAppBuildPath  = "./"
	defaultAppGitBranch  = "master"
	defaultAppEnv        = "dev"
	defaultAppLabelsKey  = "app"
	defaultDockerHub     = "test.hub.com/test"
	defaultNamespace     = "default"
	defaultReleaseCheck  = "1"
	defaultReleaseAction = "check"
	defaultHPAKind       = "Deployment"
)

type Config struct {
	// k8s
	deployment *appsv1.Deployment
	pod        *v1.Pod
	grayPod    *v1.Pod
	service    *v1.Service
	hpa        *autov1.HorizontalPodAutoscaler

	// docker
	dockerHub     string
	dockerFile    string
	appHPA        string
	appBuildCmd   string
	appBuildPath  string
	appGitBranch  string
	appEnv        string
	image         string
	releaseCheck  string
	releaseAction string
}

func NewConfig() (config *Config, err error) {
	config = new(Config)
	clt, err := consul.NewClient()
	if err != nil {
		return config, err
	}

	clt.SetBasePath(consulPath)

	dockerFiles, err := clt.Get("dockerfile")
	if err != nil {
		return config, err
	}
	config.dockerFile = string(dockerFiles)

	// deployment
	deploys, err := clt.Get("deploy")
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(deploys, &config.deployment); err != nil {
		return config, err
	}

	// pod
	pods, err := clt.Get("pod")
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(pods, &config.pod); err != nil {
		return config, err
	}

	// service
	service, err := clt.Get("service")
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(service, &config.service); err != nil {
		return config, err
	}

	// hpa
	hpa, err := clt.Get("hpa")
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(hpa, &config.hpa); err != nil {
		return config, err
	}

	// env value
	config.GetEnv()
	// deployment + pod
	config.deployment.Spec.Selector.MatchLabels[defaultAppLabelsKey] = config.deployment.Name
	config.deployment.Spec.Template.ObjectMeta.Labels[defaultAppLabelsKey] = config.deployment.Name
	config.pod.ObjectMeta.Labels[defaultAppLabelsKey] = config.pod.Name
	appTags := fmt.Sprintf("%s:%s", config.pod.Name, ReleaseTag(config.appGitBranch, config.appEnv))
	config.image = path.Join(config.dockerHub, appTags)
	config.pod.Spec.Containers[0].Image = config.image
	config.deployment.Spec.Template.Spec = config.pod.Spec
	// service
	config.service.Namespace = config.deployment.Namespace
	config.service.Name = config.deployment.Name
	ports := config.pod.Spec.Containers[0].Ports
	config.service.Spec.Ports = make([]v1.ServicePort, len(ports))
	for k, v := range ports {
		config.service.Spec.Ports[k].Name = strconv.Itoa(int(v.ContainerPort))
		config.service.Spec.Ports[k].Port = v.ContainerPort
		config.service.Spec.Ports[k].TargetPort = intstr.FromInt(int(v.ContainerPort))
	}
	config.service.Spec.Selector[defaultAppLabelsKey] = config.deployment.Name

	// hpa
	config.hpa.Name = config.pod.Name
	config.hpa.Namespace = config.pod.Namespace
	config.hpa.Spec.ScaleTargetRef.Name = config.pod.Name

	// same name
	if config.deployment.Namespace != config.pod.Namespace {
		return config, fmt.Errorf("deployment namespace: %s, pod namespace: %s",
			config.deployment.Namespace, config.pod.Namespace)
	}

	// gray pod
	config.grayPod = config.pod
	config.grayPod.Name = config.grayPod.Name + "-gray"

	return config, nil
}

func (c *Config) GetEnv() {
	// deployment
	n := os.Getenv(DeployNamespace)
	if NoNull(n) {
		c.deployment.Namespace = n
	} else {
		c.deployment.Namespace = defaultNamespace
	}

	name := os.Getenv(DeployName)
	if NoNull(name) {
		c.deployment.Name = name
	}

	r := os.Getenv(DeployReplicas)
	if NoNull(r) {
		c.deployment.Spec.Replicas = strToInt32p(r)
	}

	mrs := os.Getenv(DeployMinReadySeconds)
	if NoNull(mrs) {
		c.deployment.Spec.MinReadySeconds = strToInt32(mrs)
	}

	rhl := os.Getenv(DeployRevisionHistoryLimit)
	if NoNull(rhl) {
		c.deployment.Spec.RevisionHistoryLimit = strToInt32p(rhl)
	}

	// pod
	pn := os.Getenv(PodNamespace)
	if NoNull(pn) {
		c.pod.Namespace = pn
	} else {
		c.pod.Namespace = defaultNamespace
	}

	podName := os.Getenv(PodName)
	if NoNull(podName) {
		c.pod.Name = podName
	}

	tgps := os.Getenv(PodTerminationGracePeriodSeconds)
	if NoNull(tgps) {
		c.pod.Spec.TerminationGracePeriodSeconds = strToInt64p(tgps)
	}

	ph := os.Getenv(PodHPA)
	if NoNull(ph) {
		c.appHPA = ph
	}

	phk := os.Getenv(PodHPAKind)
	if NoNull(phk) {
		c.hpa.Spec.ScaleTargetRef.Kind = phk
	} else {
		c.hpa.Spec.ScaleTargetRef.Kind = defaultHPAKind
	}

	phm := os.Getenv(PodHPAMin)
	if NoNull(phm) {
		c.hpa.Spec.MinReplicas = strToInt32p(phm)
	}

	phmm := os.Getenv(PodHPAMax)
	if NoNull(phmm) {
		c.hpa.Spec.MaxReplicas = strToInt32(phmm)
	}

	phc := os.Getenv(PodHPACPU)
	if NoNull(phc) {
		c.hpa.Spec.TargetCPUUtilizationPercentage = strToInt32p(phc)
	}

	ap := os.Getenv(AppPort)
	if NoNull(ap) {
		c.pod.Spec.Containers[0].Ports = strToContainerPorts(ap)
	}

	lc := os.Getenv(AppLimitsCPU)
	if NoNull(lc) {
		c.pod.Spec.Containers[0].Resources.Limits[v1.ResourceCPU] = resource.MustParse(lc)
	}

	lm := os.Getenv(AppLimitsMemory)
	if NoNull(lm) {
		c.pod.Spec.Containers[0].Resources.Limits[v1.ResourceMemory] = resource.MustParse(lm)
	}

	rc := os.Getenv(AppRequestsCPU)
	if NoNull(rc) {
		c.pod.Spec.Containers[0].Resources.Requests[v1.ResourceCPU] = resource.MustParse(rc)
	}

	rm := os.Getenv(AppRequestsMemory)
	if NoNull(rm) {
		c.pod.Spec.Containers[0].Resources.Requests[v1.ResourceMemory] = resource.MustParse(rm)
	}

	lp := os.Getenv(AppLivenessPath)
	if NoNull(lp) {
		c.pod.Spec.Containers[0].LivenessProbe.HTTPGet.Path = lp
	}

	ld := os.Getenv(AppLivenessDelay)
	if NoNull(ld) {
		c.pod.Spec.Containers[0].LivenessProbe.InitialDelaySeconds = strToInt32(ld)
	}

	rp := os.Getenv(AppReadinessPath)
	if NoNull(rp) {
		c.pod.Spec.Containers[0].ReadinessProbe.HTTPGet.Path = rp
	}

	rd := os.Getenv(AppReadinessDelay)
	if NoNull(rd) {
		c.pod.Spec.Containers[0].ReadinessProbe.InitialDelaySeconds = strToInt32(rd)
	}

	hcc := os.Getenv(AppHealthCheck)
	if NoNull(hcc) && hcc == "0" {
		c.pod.Spec.Containers[0].LivenessProbe = nil
		c.pod.Spec.Containers[0].ReadinessProbe = nil
	}

	ac := os.Getenv(AppCmd)
	if NoNull(ac) {
		c.pod.Spec.Containers[0].Command = strings.Split(ac, " ")
	}

	acp := os.Getenv(AppCmdPath)
	if NoNull(acp) {
		c.pod.Spec.Containers[0].WorkingDir = acp
	}

	abc := os.Getenv(AppBuildCmd)
	if NoNull(abc) {
		c.appBuildCmd = abc
	} else {
		c.appBuildCmd = defaultAppBuildCmd
	}

	abp := os.Getenv(AppBuildPath)
	if NoNull(abp) {
		c.appBuildPath = abp
	} else {
		c.appBuildPath = defaultAppBuildPath
	}

	hub := os.Getenv(DockerHub)
	if NoNull(hub) {
		c.dockerHub = hub
	} else {
		c.dockerHub = defaultDockerHub
	}

	agb := os.Getenv(AppGitBranch)
	if NoNull(agb) {
		c.appGitBranch = agb
	} else {
		c.appGitBranch = defaultAppGitBranch
	}

	env := os.Getenv(AppEnv)
	if NoNull(env) {
		c.appEnv = env
	} else {
		c.appEnv = defaultAppEnv
	}

	c.releaseCheck = os.Getenv(ReleaseCheck)
	if !NoNull(c.releaseCheck) {
		c.releaseCheck = defaultReleaseCheck
	}

	c.releaseAction = os.Getenv(ReleaseAction)
	if !NoNull(c.releaseAction) {
		c.releaseAction = defaultReleaseAction
	}
}

func NoNull(str string) bool {
	if str == "" {
		return false
	}

	return true
}

func ReleaseTag(str, env string) string {
	for _, v := range "/\\_" {
		str = strings.Replace(str, string(v), "-", -1)
	}

	if env == "prd" {
		//random := rand.Intn(10000000000000000)
		//times := time.Now().UnixNano()
		str = fmt.Sprintf("%s-%s", env, str)
		return str
	}

	return env
}
