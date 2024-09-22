// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// VpcDnsLister helps list VpcDnses.
// All objects returned here must be treated as read-only.
type VpcDnsLister interface {
	// List lists all VpcDnses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.VpcDns, err error)
	// Get retrieves the VpcDns from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.VpcDns, error)
	VpcDnsListerExpansion
}

// vpcDnsLister implements the VpcDnsLister interface.
type vpcDnsLister struct {
	listers.ResourceIndexer[*kubeovnv1.VpcDns]
}

// NewVpcDnsLister returns a new VpcDnsLister.
func NewVpcDnsLister(indexer cache.Indexer) VpcDnsLister {
	return &vpcDnsLister{listers.New[*kubeovnv1.VpcDns](indexer, kubeovnv1.Resource("vpcdns"))}
}
