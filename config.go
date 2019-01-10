package main

import (
	"encoding/json"
	"os"
	"path"

	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/lijiansgit/go/libs/consul"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"
)

var (
	kubePath   = "kubernetes/v1"
	deployment *v1beta1.Deployment
	pod        *v1.Pod
)

func ReadConsul() (err error) {
	clt, err := consul.NewClient()
	if err != nil {
		return err
	}

	clt.SetBasePath(kubePath)

	deploys, err := clt.Get(path.Join("deploy"))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(deploys, &deployment); err != nil {
		return err
	}

	pods, err := clt.Get(path.Join("pod"))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(pods, &pod); err != nil {
		return err
	}

	GetEnv()

	deployment.Spec.Template.Spec = pod.Spec
	return nil
}

func GetEnv() {
	// deployment
	n := os.Getenv(DeployNamespace)
	if NoNull(n) {
		deployment.SetNamespace(n)
	}

	r := os.Getenv(DeployReplicas)
	if NoNull(r) {
		deployment.Spec.Replicas = strToInt32p(r)
	}

	mrs := os.Getenv(DeployMinReadySeconds)
	if NoNull(mrs) {
		deployment.Spec.MinReadySeconds = strToInt32(mrs)
	}

	rhl := os.Getenv(DeployRevisionHistoryLimit)
	if NoNull(rhl) {
		deployment.Spec.RevisionHistoryLimit = strToInt32p(rhl)
	}

	// pod
	ns := os.Getenv(PodNamespace)
	if NoNull(ns) {
		pod.Namespace = ns
	}

	tgps := os.Getenv(PodTerminationGracePeriodSeconds)
	if NoNull(tgps) {
		pod.Spec.TerminationGracePeriodSeconds = strToInt64p(tgps)
	}

	ap := os.Getenv(AppPort)
	if NoNull(ap) {
		pod.Spec.Containers[0].Ports[0].ContainerPort = strToInt32(ap)
	}

	lc := os.Getenv(AppLimitsCPU)
	if NoNull(lc) {
		pod.Spec.Containers[0].Resources.Limits[v1.ResourceCPU] = resource.MustParse(lc)
	}

	lm := os.Getenv(AppLimitsMemory)
	if NoNull(lm) {
		pod.Spec.Containers[0].Resources.Limits[v1.ResourceMemory] = resource.MustParse(lm)
	}

	rc := os.Getenv(AppRequestsCPU)
	if NoNull(rc) {
		pod.Spec.Containers[0].Resources.Requests[v1.ResourceCPU] = resource.MustParse(rc)
	}

	rm := os.Getenv(AppRequestsMemory)
	if NoNull(rm) {
		pod.Spec.Containers[0].Resources.Requests[v1.ResourceMemory] = resource.MustParse(rm)
	}

	lp := os.Getenv(AppLivenessPath)
	if NoNull(lp) {
		pod.Spec.Containers[0].LivenessProbe.HTTPGet.Path = lp
	}

	ld := os.Getenv(AppLivenessDelay)
	if NoNull(ld) {
		pod.Spec.Containers[0].LivenessProbe.InitialDelaySeconds = strToInt32(ld)
	}

	rp := os.Getenv(AppReadinessPath)
	if NoNull(rp) {
		pod.Spec.Containers[0].ReadinessProbe.HTTPGet.Path = rp
	}

	rd := os.Getenv(AppReadinessDelay)
	if NoNull(rd) {
		pod.Spec.Containers[0].ReadinessProbe.InitialDelaySeconds = strToInt32(rd)
	}
}

func NoNull(str string) bool {
	if str == "" {
		return false
	}

	return true
}
