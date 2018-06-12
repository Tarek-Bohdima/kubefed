/*
Copyright 2018 The Kubernetes Authors.

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

package v1alpha1

import (
	time "time"

	core_v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	versioned "github.com/kubernetes-sigs/federation-v2/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubernetes-sigs/federation-v2/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/client/listers/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FederatedServiceInformer provides access to a shared informer and lister for
// FederatedServices.
type FederatedServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FederatedServiceLister
}

type federatedServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedServiceInformer constructs a new informer for FederatedService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedServiceInformer constructs a new informer for FederatedService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().FederatedServices(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CoreV1alpha1().FederatedServices(namespace).Watch(options)
			},
		},
		&core_v1alpha1.FederatedService{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&core_v1alpha1.FederatedService{}, f.defaultInformer)
}

func (f *federatedServiceInformer) Lister() v1alpha1.FederatedServiceLister {
	return v1alpha1.NewFederatedServiceLister(f.Informer().GetIndexer())
}
