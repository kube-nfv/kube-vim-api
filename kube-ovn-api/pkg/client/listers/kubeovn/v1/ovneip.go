// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// OvnEipLister helps list OvnEips.
// All objects returned here must be treated as read-only.
type OvnEipLister interface {
	// List lists all OvnEips in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.OvnEip, err error)
	// Get retrieves the OvnEip from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.OvnEip, error)
	OvnEipListerExpansion
}

// ovnEipLister implements the OvnEipLister interface.
type ovnEipLister struct {
	listers.ResourceIndexer[*kubeovnv1.OvnEip]
}

// NewOvnEipLister returns a new OvnEipLister.
func NewOvnEipLister(indexer cache.Indexer) OvnEipLister {
	return &ovnEipLister{listers.New[*kubeovnv1.OvnEip](indexer, kubeovnv1.Resource("ovneip"))}
}