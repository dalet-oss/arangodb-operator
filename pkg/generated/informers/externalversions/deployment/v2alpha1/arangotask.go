//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//

// Code generated by informer-gen. DO NOT EDIT.

package v2alpha1

import (
	"context"
	time "time"

	deploymentv2alpha1 "github.com/arangodb/kube-arangodb/pkg/apis/deployment/v2alpha1"
	versioned "github.com/arangodb/kube-arangodb/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/arangodb/kube-arangodb/pkg/generated/informers/externalversions/internalinterfaces"
	v2alpha1 "github.com/arangodb/kube-arangodb/pkg/generated/listers/deployment/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ArangoTaskInformer provides access to a shared informer and lister for
// ArangoTasks.
type ArangoTaskInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2alpha1.ArangoTaskLister
}

type arangoTaskInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewArangoTaskInformer constructs a new informer for ArangoTask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewArangoTaskInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredArangoTaskInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredArangoTaskInformer constructs a new informer for ArangoTask type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredArangoTaskInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DatabaseV2alpha1().ArangoTasks(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.DatabaseV2alpha1().ArangoTasks(namespace).Watch(context.TODO(), options)
			},
		},
		&deploymentv2alpha1.ArangoTask{},
		resyncPeriod,
		indexers,
	)
}

func (f *arangoTaskInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredArangoTaskInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *arangoTaskInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&deploymentv2alpha1.ArangoTask{}, f.defaultInformer)
}

func (f *arangoTaskInformer) Lister() v2alpha1.ArangoTaskLister {
	return v2alpha1.NewArangoTaskLister(f.Informer().GetIndexer())
}
