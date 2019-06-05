/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	time "time"

	servicecatalogv1beta1 "github.com/kubernetes-sigs/service-catalog/pkg/apis/servicecatalog/v1beta1"
	clientset "github.com/kubernetes-sigs/service-catalog/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/kubernetes-sigs/service-catalog/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1beta1 "github.com/kubernetes-sigs/service-catalog/pkg/client/listers_generated/servicecatalog/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ServiceBrokerInformer provides access to a shared informer and lister for
// ServiceBrokers.
type ServiceBrokerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.ServiceBrokerLister
}

type serviceBrokerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewServiceBrokerInformer constructs a new informer for ServiceBroker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewServiceBrokerInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredServiceBrokerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredServiceBrokerInformer constructs a new informer for ServiceBroker type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredServiceBrokerInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecatalogV1beta1().ServiceBrokers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServicecatalogV1beta1().ServiceBrokers(namespace).Watch(options)
			},
		},
		&servicecatalogv1beta1.ServiceBroker{},
		resyncPeriod,
		indexers,
	)
}

func (f *serviceBrokerInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredServiceBrokerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *serviceBrokerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&servicecatalogv1beta1.ServiceBroker{}, f.defaultInformer)
}

func (f *serviceBrokerInformer) Lister() v1beta1.ServiceBrokerLister {
	return v1beta1.NewServiceBrokerLister(f.Informer().GetIndexer())
}
