/*
SPDX-FileCopyrightText: 2020 SAP SE or an SAP affiliate company and Gardener contributors

SPDX-License-Identifier: Apache-2.0
*/

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=extensions.gardener.cloud, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("backupbuckets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().BackupBuckets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("backupentries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().BackupEntries().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().Clusters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("containerruntimes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().ContainerRuntimes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("controlplanes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().ControlPlanes().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("extensions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().Extensions().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("infrastructures"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().Infrastructures().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("networks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().Networks().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("operatingsystemconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().OperatingSystemConfigs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("workers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Extensions().V1alpha1().Workers().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
