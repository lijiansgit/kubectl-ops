package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"path"
)

var (
	clientset *kubernetes.Clientset
	kubeConf string
)

func init() {
	defaultKubeConf := path.Join(os.Getenv("HOME"), ".kube/config")
	flag.StringVar(&kubeConf, "c", defaultKubeConf, "kubernetes client config file path")
}

func main() {
	flag.Parse()

	err := ReadConsul()
	if err != nil {
		panic(err)
	}

	conf, err := clientcmd.BuildConfigFromFlags("", kubeConf)
	if err != nil {
		panic(err)
	}

	clientset, err = kubernetes.NewForConfig(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println("connect kubernets ok")

	_, err = clientset.AppsV1beta1().Deployments("default").Create(deployment)
	if err != nil {
		panic(err)
	}

	//huidu
	pod.Name = pod.Name + "-huidu"
	_, err = clientset.CoreV1().Pods("default").Create(pod)
	if err != nil {
		panic(err)
	}


	fmt.Println("connect kubernets end")
}
