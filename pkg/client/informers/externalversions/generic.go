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

package externalversions

import (
	"fmt"

	v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/core/v1alpha1"
	multiclusterdns_v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/multiclusterdns/v1alpha1"
	scheduling_v1alpha1 "github.com/kubernetes-sigs/federation-v2/pkg/apis/scheduling/v1alpha1"
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
	// Group=core.federation.k8s.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("federatedclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedClusters().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedconfigmaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedConfigMaps().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedconfigmapoverrides"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedConfigMapOverrides().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedconfigmapplacements"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedConfigMapPlacements().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federateddeployments"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedDeployments().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federateddeploymentoverrides"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedDeploymentOverrides().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federateddeploymentplacements"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedDeploymentPlacements().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedjobs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedJobs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedjoboverrides"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedJobOverrides().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedjobplacements"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedJobPlacements().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatednamespaceplacements"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedNamespacePlacements().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedreplicasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedReplicaSets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedreplicasetoverrides"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedReplicaSetOverrides().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedreplicasetplacements"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedReplicaSetPlacements().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedsecrets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedSecrets().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedsecretoverrides"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedSecretOverrides().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedsecretplacements"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedSecretPlacements().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedServices().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedserviceplacements"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedServicePlacements().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("federatedtypeconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().FederatedTypeConfigs().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("propagatedversions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Core().V1alpha1().PropagatedVersions().Informer()}, nil

		// Group=multiclusterdns.federation.k8s.io, Version=v1alpha1
	case multiclusterdns_v1alpha1.SchemeGroupVersion.WithResource("dnsendpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().DNSEndpoints().Informer()}, nil
	case multiclusterdns_v1alpha1.SchemeGroupVersion.WithResource("multiclusterservicednsrecords"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Multiclusterdns().V1alpha1().MultiClusterServiceDNSRecords().Informer()}, nil

		// Group=scheduling.federation.k8s.io, Version=v1alpha1
	case scheduling_v1alpha1.SchemeGroupVersion.WithResource("replicaschedulingpreferences"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Scheduling().V1alpha1().ReplicaSchedulingPreferences().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
