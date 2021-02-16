// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	flyteworkflowv1alpha1 "github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	versioned "github.com/lyft/flytepropeller/pkg/client/clientset/versioned"
	internalinterfaces "github.com/lyft/flytepropeller/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/lyft/flytepropeller/pkg/client/listers/flyteworkflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FlyteWorkflowInformer provides access to a shared informer and lister for
// FlyteWorkflows.
type FlyteWorkflowInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FlyteWorkflowLister
}

type flyteWorkflowInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFlyteWorkflowInformer constructs a new informer for FlyteWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFlyteWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFlyteWorkflowInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFlyteWorkflowInformer constructs a new informer for FlyteWorkflow type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFlyteWorkflowInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlyteworkflowV1alpha1().FlyteWorkflows(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.FlyteworkflowV1alpha1().FlyteWorkflows(namespace).Watch(context.TODO(), options)
			},
		},
		&flyteworkflowv1alpha1.FlyteWorkflow{},
		resyncPeriod,
		indexers,
	)
}

func (f *flyteWorkflowInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFlyteWorkflowInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *flyteWorkflowInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&flyteworkflowv1alpha1.FlyteWorkflow{}, f.defaultInformer)
}

func (f *flyteWorkflowInformer) Lister() v1alpha1.FlyteWorkflowLister {
	return v1alpha1.NewFlyteWorkflowLister(f.Informer().GetIndexer())
}
