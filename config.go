package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/lijiansgit/go/libs/consul"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
)

const (
	consulPath          = "kubernetes/v1"
	dockerFileName      = "Dockerfile"
	defaultAppBuildCmd  = "make"
	defaultAppBuildPath = "./"
	defaultDockerHub    = "test.hub.com"
	defaultNamespace    = "default"
)

type Config struct {
	// k8s
	deployment *v1beta1.Deployment
	pod        *v1.Pod
	grayPod    *v1.Pod
	image      string
	service    *v1.Service
	ingress    *v1beta1.Ingress

	// docker
	dockerHub    string
	dockerFile   string
	appBuildCmd  string
	appBuildPath string
	appGitBranch string
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

	deploys, err := clt.Get("deploy")
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(deploys, &config.deployment); err != nil {
		return config, err
	}

	pods, err := clt.Get("pod")
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(pods, &config.pod); err != nil {
		return config, err
	}

	service, err := clt.Get("service")
	if err != nil {
		return config, err
	}

	if err = json.Unmarshal(service, &config.service); err != nil {
		return config, err
	}

	config.GetEnv()

	appTags := fmt.Sprintf("%s:%s", config.pod.Name, config.appGitBranch)
	config.image = path.Join(config.dockerHub, appTags)
	config.pod.Spec.Containers[0].Image = config.image
	config.deployment.Spec.Template.Spec = config.pod.Spec

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
	ns := os.Getenv(PodNamespace)
	if NoNull(ns) {
		c.pod.Namespace = ns
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

	ap := os.Getenv(AppPort)
	if NoNull(ap) {
		c.pod.Spec.Containers[0].Ports[0].ContainerPort = strToInt32(ap)
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
	}
}

func NoNull(str string) bool {
	if str == "" {
		return false
	}

	return true
}
