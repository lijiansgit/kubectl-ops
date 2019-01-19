package main

import (
	"flag"
	"os"
	"path"

	log "github.com/lijiansgit/go/libs/log4go"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	clientset *kubernetes.Clientset
	kubeConf  string
	verbose   bool
	action    string
	config    *Config
)

func init() {
	defaultKubeConf := path.Join(os.Getenv("HOME"), ".kube/config")
	flag.StringVar(&kubeConf, "c", defaultKubeConf, "kubernetes client config file path")
	flag.BoolVar(&verbose, "v", true, "log verbose")
	flag.StringVar(&action, "a", "deploy", "kubernetes client action: deploy/gray/rollback")
}

func main() {
	flag.Parse()

	if verbose == true {
		log.AddFilter("stdout", log.DEBUG, log.NewConsoleLogWriter())
	} else {
		log.AddFilter("stdout", log.INFO, log.NewConsoleLogWriter())
	}
	defer log.Close()

	var err error
	config, err = NewConfig()
	if err != nil {
		panic(err)
	}

	if err = build(); err != nil {
		panic(err)
	}

	kconf, err := clientcmd.BuildConfigFromFlags("", kubeConf)
	if err != nil {
		panic(err)
	}

	clientset, err = kubernetes.NewForConfig(kconf)
	if err != nil {
		panic(err)
	}

	log.Debug("Connect Kubernets: %s OK", kconf.Host)

	release()

	log.Debug("Connect Kubernets: %s END", kconf.Host)
}
