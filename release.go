package main

import (
	log "github.com/lijiansgit/go/libs/log4go"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	deleteGrayPod()
	log.Debug("deployment: %s", config.deployment.String())
	_, err := clientset.ExtensionsV1beta1().Deployments(
		config.deployment.Namespace).Get(config.deployment.Name, v1.GetOptions{})
	if err != nil {
		log.Warn("deployment: %s no exist, create...", config.deployment.Name)
		createDeployment()
	} else {
		updateDeployment()
	}
}

func createDeployment() {
	_, err := clientset.ExtensionsV1beta1().Deployments(
		config.deployment.Namespace).Create(config.deployment)
	if err != nil {
		log.Error("deployment: %s create err(%v)", config.deployment.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("deployment: %s create ok", config.deployment.Name)
}

func updateDeployment() {
	_, err := clientset.ExtensionsV1beta1().Deployments(
		config.deployment.Namespace).Update(config.deployment)
	if err != nil {
		log.Error("deployment: %s update err(%v)", config.deployment.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("deployment: %s update ok", config.deployment.Name)
}

func gray() {
	// 灰度 TODO
	deleteGrayPod()
	createGrayPod()
}

func createGrayPod() {
	log.Debug("gray pod: %s", config.grayPod.String())
	_, err := clientset.CoreV1().Pods(config.grayPod.Namespace).Create(config.grayPod)
	if err != nil {
		log.Error("gray pod: %s create err (%v)", config.grayPod.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("gray pod: %s create ok", config.grayPod.Name)
}

func deleteGrayPod() {
	err := clientset.CoreV1().Pods(config.grayPod.Namespace).Delete(
		config.grayPod.Name, new(metav1.DeleteOptions))
	if err != nil {
		log.Error("gray pod: %s delete err (%v)", config.grayPod.Name, err)
		panic(ErrKubeclt)
	}

	log.Info("gray pod: %s delete ok", config.grayPod.Name)
}

func rollback() {

}
