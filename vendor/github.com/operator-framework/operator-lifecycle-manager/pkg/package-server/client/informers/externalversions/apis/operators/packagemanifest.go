/*
Copyright 2019 Red Hat, Inc.

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

package operators

import (
	time "time"

	apisoperators "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apis/operators"
	versioned "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/client/clientset/versioned"
	internalinterfaces "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/client/informers/externalversions/internalinterfaces"
	operators "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/client/listers/apis/operators"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PackageManifestInformer provides access to a shared informer and lister for
// PackageManifests.
type PackageManifestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() operators.PackageManifestLister
}

type packageManifestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPackageManifestInformer constructs a new informer for PackageManifest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPackageManifestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPackageManifestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPackageManifestInformer constructs a new informer for PackageManifest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPackageManifestInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Operators().PackageManifests(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Operators().PackageManifests(namespace).Watch(options)
			},
		},
		&apisoperators.PackageManifest{},
		resyncPeriod,
		indexers,
	)
}

func (f *packageManifestInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPackageManifestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *packageManifestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisoperators.PackageManifest{}, f.defaultInformer)
}

func (f *packageManifestInformer) Lister() operators.PackageManifestLister {
	return operators.NewPackageManifestLister(f.Informer().GetIndexer())
}
