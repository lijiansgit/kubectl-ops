package main

import (
	log "github.com/lijiansgit/go/libs/log4go"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/kubernetes/typed/extensions/v1beta1"
)

var (
	ErrKubeclt = errors.New("kubectl exec fail!")
)

func release() {
	if action == "deploy" {
		deploy()
	}

	if action == "gray" {
		gray()
	}

	if action == "rollback" {
		rollback()
	}
}

func deploy() {
	r := NewRelease()
	r.DeletePod()
	//log.Debug("deployment: %s", config.deployment.String())
	err := r.GetDeployment()
	return
	if err != nil {
		log.Warn("deployment: %s no exist, create...", config.deployment.Name)
		r.CreateDeployment()
	} else {
		r.UpdateDeployment()
	}

	err = r.GetService()
	if err != nil {
		log.Warn("service: %s no exist, create...", config.service.Name)
		r.CreateService()
	} else {
		log.Warn("service: %s exist, skip...", config.service.Name)
	}
}

type Release struct {
	deploymentClt v1beta1.DeploymentInterface
	podClt        corev1.PodInterface
	serviceClt    corev1.ServiceInterface
}

func NewRelease() *Release {
	release := new(Release)
	release.deploymentClt = clientset.ExtensionsV1beta1().Deployments(config.deployment.Namespace)
	release.podClt = clientset.CoreV1().Pods(config.pod.Namespace)
	release.serviceClt = clientset.CoreV1().Services(config.service.Namespace)
	return release
}

func (r *Release) GetDeployment() (err error) {
	_, err = r.deploymentClt.Get(config.deployment.Name, v1.GetOptions{})
	return err
}

func (r *Release) StatusDeployment() (err error) {
	dl, err := r.deploymentClt.Watch(v1.ListOptions{})
	log.Info("%v", dl)
	return err
}

func (r *Release) CreateDeployment() {
	_, err := r.deploymentClt.Create(config.deployment)
	if err != nil {
		log.Error("deployment: %s create err(%v)", config.deployment.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("deployment: %s create ok", config.deployment.Name)
}

func (r *Release) UpdateDeployment() {
	_, err := r.deploymentClt.Update(config.deployment)
	if err != nil {
		log.Error("deployment: %s update err(%v)", config.deployment.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("deployment: %s update ok", config.deployment.Name)
}

func (r *Release) CreatePod() {
	log.Debug("gray pod: %s", config.grayPod.String())
	_, err := r.podClt.Create(config.grayPod)
	if err != nil {
		log.Error("pod: %s create err (%v)", config.grayPod.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("pod: %s create ok", config.grayPod.Name)
}

func (r *Release) DeletePod() {
	return
	err := r.podClt.Delete(config.grayPod.Name, new(metav1.DeleteOptions))
	if err != nil {
		log.Error("pod: %s delete err (%v)", config.grayPod.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("pod: %s delete ok", config.grayPod.Name)
}

func (r *Release) GetService() (err error) {
	_, err = r.serviceClt.Get(config.service.Name, v1.GetOptions{})
	return err
}

func (r *Release) CreateService() {
	_, err := r.serviceClt.Create(config.service)
	if err != nil {
		log.Error("service: %s create err(%v)", config.service.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("service: %s create ok", config.service.Name)
}

func (r *Release) UpdateService() {
	_, err := r.serviceClt.Update(config.service)
	if err != nil {
		log.Error("service: %s update err(%v)", config.service.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("service: %s update ok", config.service.Name)
}

func gray() {
	// 灰度 TODO
}

func rollback() {

}
