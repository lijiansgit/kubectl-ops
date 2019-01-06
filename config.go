package main

import (
	"encoding/json"
	"github.com/lijiansgit/go/libs/consul"
	"k8s.io/api/apps/v1beta1"
	"k8s.io/api/core/v1"
	"path"
)

var (
	consulPath = "kubernetes/v1"
	deployment *v1beta1.Deployment
	pod *v1.Pod
)

func ReadConsul() (err error) {
	clt, err := consul.NewClient()
	if err != nil {
		return err
	}

	deploys, err := clt.Get(path.Join(consulPath, "deploy"))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(deploys, &deployment); err != nil {
		return err
	}


	pods, err := clt.Get(path.Join(consulPath, "pod"))
	if err != nil {
		return err
	}

	if err = json.Unmarshal(pods, &pod); err != nil {
		return err
	}

	deployment.Spec.Template.Spec = pod.Spec
	return nil
}

func NoNull(str string) bool {
	if str == "" {
		return false
	}

	return true
}
