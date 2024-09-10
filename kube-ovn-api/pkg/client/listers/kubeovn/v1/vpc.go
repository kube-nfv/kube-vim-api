// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// VpcLister helps list Vpcs.
// All objects returned here must be treated as read-only.
type VpcLister interface {
	// List lists all Vpcs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.Vpc, err error)
	// Get retrieves the Vpc from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.Vpc, error)
	VpcListerExpansion
}

// vpcLister implements the VpcLister interface.
type vpcLister struct {
	listers.ResourceIndexer[*kubeovnv1.Vpc]
}

// NewVpcLister returns a new VpcLister.
func NewVpcLister(indexer cache.Indexer) VpcLister {
	return &vpcLister{listers.New[*kubeovnv1.Vpc](indexer, kubeovnv1.Resource("vpc"))}
}
