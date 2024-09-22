// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	fmt "fmt"

	v1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
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
	// Group=kubeovn.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("ips"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().IPs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ippools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().IPPools().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("iptables-dnat-rules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().IptablesDnatRules().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("iptables-eips"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().IptablesEIPs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("iptables-fip-rules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().IptablesFIPRules().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("iptables-snat-rules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().IptablesSnatRules().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ovn-dnat-rules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().OvnDnatRules().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ovn-eips"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().OvnEips().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ovn-fips"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().OvnFips().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("ovn-snat-rules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().OvnSnatRules().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("provider-networks"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().ProviderNetworks().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("qos-policies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().QoSPolicies().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("security-groups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().SecurityGroups().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("subnets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().Subnets().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("switch-lb-rules"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().SwitchLBRules().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("vips"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().Vips().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("vlans"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().Vlans().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("vpcs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().Vpcs().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("vpc-dnses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().VpcDnses().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("vpc-nat-gateways"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubeovn().V1().VpcNatGateways().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
