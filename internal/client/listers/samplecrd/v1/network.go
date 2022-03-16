/* test for k8s code generator */
// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/aseara/k8scr/internal/api/samplecrd/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkLister helps list Networks.
// All objects returned here must be treated as read-only.
type NetworkLister interface {
	// List lists all Networks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Network, err error)
	// Networks returns an object that can list and get Networks.
	Networks(namespace string) NetworkNamespaceLister
	NetworkListerExpansion
}

// networkLister implements the NetworkLister interface.
type networkLister struct {
	indexer cache.Indexer
}

// NewNetworkLister returns a new NetworkLister.
func NewNetworkLister(indexer cache.Indexer) NetworkLister {
	return &networkLister{indexer: indexer}
}

// List lists all Networks in the indexer.
func (s *networkLister) List(selector labels.Selector) (ret []*v1.Network, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Network))
	})
	return ret, err
}

// Networks returns an object that can list and get Networks.
func (s *networkLister) Networks(namespace string) NetworkNamespaceLister {
	return networkNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkNamespaceLister helps list and get Networks.
// All objects returned here must be treated as read-only.
type NetworkNamespaceLister interface {
	// List lists all Networks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Network, err error)
	// Get retrieves the Network from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Network, error)
	NetworkNamespaceListerExpansion
}

// networkNamespaceLister implements the NetworkNamespaceLister
// interface.
type networkNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Networks in the indexer for a given namespace.
func (s networkNamespaceLister) List(selector labels.Selector) (ret []*v1.Network, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Network))
	})
	return ret, err
}

// Get retrieves the Network from the indexer for a given namespace and name.
func (s networkNamespaceLister) Get(name string) (*v1.Network, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("network"), name)
	}
	return obj.(*v1.Network), nil
}
