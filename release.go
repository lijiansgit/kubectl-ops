package main

import (
	"time"

	"github.com/pkg/errors"
	"k8s.io/api/core/v1"

	log "github.com/lijiansgit/go/libs/log4go"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	cltappsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	cltcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
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
	log.Debug("deployment: %s", config.deployment.String())

	if _, err := r.GetDeployment(); err != nil {
		log.Warn("deployment: %s no exist, create...", config.deployment.Name)
		r.CreateDeployment()
	} else {
		r.UpdateDeployment()
	}

	if _, err := r.GetService(); err != nil {
		log.Warn("service: %s no exist, create...", config.service.Name)
		r.CreateService()
	} else {
		log.Warn("service: %s exist, skip...", config.service.Name)
	}

	r.CheckDeployment()
}

func gray() {
	// 灰度 todo
	r := NewRelease()
	log.Debug("gray pod: %s", config.grayPod.String())

	r.DeletePod()
	r.CreatePod()

	r.CheckPod()
}

func rollback() {
	// todo
}

type Release struct {
	deploymentClt cltappsv1.DeploymentInterface
	podClt        cltcorev1.PodInterface
	serviceClt    cltcorev1.ServiceInterface
}

func NewRelease() *Release {
	release := new(Release)
	release.deploymentClt = clientset.AppsV1().Deployments(config.deployment.Namespace)
	release.podClt = clientset.CoreV1().Pods(config.pod.Namespace)
	release.serviceClt = clientset.CoreV1().Services(config.service.Namespace)

	return release
}

func (r *Release) GetDeployment() (d *appsv1.Deployment, err error) {
	d, err = r.deploymentClt.Get(config.deployment.Name, metav1.GetOptions{})
	return d, err
}

func (r *Release) CheckDeployment() {
	for {
		time.Sleep(1e9)

		deployment, err := r.GetDeployment()
		if err != nil {
			log.Error("Get deployment err(%v)", err)
			continue
		}

		status := deployment.Status
		log.Info("Replicas: %d, Ready: %d, Updated: %d", status.Replicas,
			status.ReadyReplicas, status.UpdatedReplicas)

		if status.Replicas == status.ReadyReplicas && status.Replicas == status.UpdatedReplicas {
			break
		}
	}
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

func (r *Release) GetPod() (p *v1.Pod, err error) {
	p, err = r.podClt.Get(config.grayPod.Name, metav1.GetOptions{})
	return p, err

}

func (r *Release) CheckPod() {
	for {
		time.Sleep(1e9)

		pod, err := r.podClt.Get(config.grayPod.Name, metav1.GetOptions{})
		if err != nil {
			log.Error("Get pod err(%v)", err)
			continue
		}

		ready := true
		status := pod.Status
		for _, v := range status.ContainerStatuses {
			if !v.Ready {
				ready = false
			}

			log.Info("Container Name: %s, Ready: %v", v.Name, v.Ready)
		}

		if ready {
			break
		}
	}
}

func (r *Release) CreatePod() {
	_, err := r.podClt.Create(config.grayPod)
	if err != nil {
		log.Error("pod: %s create err (%v)", config.grayPod.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("pod: %s create ok", config.grayPod.Name)
}

func (r *Release) DeletePod() {
	if _, err := r.GetPod(); err != nil {
		log.Info("no pod: %s,skip delete pod", config.grayPod.Name)
		return
	}

	err := r.podClt.Delete(config.grayPod.Name, new(metav1.DeleteOptions))
	if err != nil {
		log.Error("pod: %s delete err (%v)", config.grayPod.Name, err)
		panic(ErrKubeclt)
	}

	for {
		if _, err := r.GetPod(); err != nil {
			break
		} else {
			log.Info("pod: %s delete ing...", config.grayPod.Name)
		}

		time.Sleep(1e9)
	}

	log.Info("pod: %s delete ok", config.grayPod.Name)
}

func (r *Release) GetService() (s *v1.Service, err error) {
	s, err = r.serviceClt.Get(config.service.Name, metav1.GetOptions{})
	return s, err
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
