package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	clientset "github.com/aseara/k8scr/internal/client/clientset/versioned"
	informers "github.com/aseara/k8scr/internal/client/informers/externalversions"
	"github.com/aseara/k8scr/internal/signal"
)

var (
	_kubeconfig string
	_masterURL  string
)

func init() {
	flag.StringVar(&_kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&_masterURL, "master", "", "The address of the Kubernetes API server. Override any value in kubeconfig. Only required if out-of-cluster.")
}

func main() {
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signal.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(_masterURL, _kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	networkClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	networkInformerFactory := informers.NewSharedInformerFactory(networkClient, time.Second*30)

	controller := NewController(kubeClient, networkClient, networkInformerFactory.Samplecrd().V1().Networks())

	// Start the informer factories to begin populating the informer cache
	glog.Info("Starting Network control loop")
	go networkInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}

}
