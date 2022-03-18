package main

import (
	corev1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/kubernetes"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	"k8s.io/client-go/kubernetes/scheme"

	samplecrdv1 "github.com/aseara/k8scr/internal/api/samplecrd/v1"
	clientset "github.com/aseara/k8scr/internal/client/clientset/versioned"
	networkscheme "github.com/aseara/k8scr/internal/client/clientset/versioned/scheme"
	informers "github.com/aseara/k8scr/internal/client/informers/externalversions/samplecrd/v1"
	listers "github.com/aseara/k8scr/internal/client/listers/samplecrd/v1"
	"github.com/golang/glog"
)

const _controllerAgentName = "network-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when a network is synced
	SuccessSynced = "Synced"

	// MessageResourceSynced is the message used for an Event fired when a Network
	// is synced successfully
	MessageResourceSynced = "Network synced successfully"
)

// Controller is the controller implementation for Network resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// networkclientset is a clientset for our own API group
	networkclientset clientset.Interface

	networksLister listers.NetworkLister
	networkSynced  cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	record record.EventRecorder
}

func NewController(
	kubeclientset kubernetes.Interface,
	networkclientset clientset.Interface,
	networkInformer informers.NetworkInformer) *Controller {

	// Create event broadcaster
	// Add sample-controller types to the default kubernetes Scheme so Events can be
	// logged for sample-controller types.
	utilruntime.Must(networkscheme.AddToScheme(scheme.Scheme))
	glog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{
		Interface: kubeclientset.CoreV1().Events(""),
	})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{
		Component: _controllerAgentName,
	})

	wq := workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Networks")

	controller := &Controller{
		kubeclientset:    kubeclientset,
		networkclientset: networkclientset,
		networksLister:   networkInformer.Lister(),
		networkSynced:    networkInformer.Informer().HasSynced,
		workqueue:        wq,
		record:           recorder,
	}

	glog.Info("Setting up event handlers")

	// Set up an event handler for when Network resources change
	networkInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueNetwork,
		UpdateFunc: func(old, new interface{}) {
			oldNetwork := old.(*samplecrdv1.Network)
			newNetwork := new.(*samplecrdv1.Network)
			if oldNetwork.ResourceVersion == newNetwork.ResourceVersion {
				// Periodic resync will send update events for all known Networks.
				// Two different versions of the same Network will always have different RVs.
				return
			}
			controller.enqueueNetwork(new)
		},
		DeleteFunc: controller.enqueueNetworkForDelete,
	})

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
//
func (c *Controller) Run(threadness int, stopCh chan<- struct{}) error {

	return nil
}

func (c *Controller) enqueueNetwork(obj interface{}) {

}

func (c *Controller) enqueueNetworkForDelete(obj interface{}) {

}
