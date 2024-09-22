// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// OvnSnatRuleLister helps list OvnSnatRules.
// All objects returned here must be treated as read-only.
type OvnSnatRuleLister interface {
	// List lists all OvnSnatRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.OvnSnatRule, err error)
	// Get retrieves the OvnSnatRule from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.OvnSnatRule, error)
	OvnSnatRuleListerExpansion
}

// ovnSnatRuleLister implements the OvnSnatRuleLister interface.
type ovnSnatRuleLister struct {
	listers.ResourceIndexer[*kubeovnv1.OvnSnatRule]
}

// NewOvnSnatRuleLister returns a new OvnSnatRuleLister.
func NewOvnSnatRuleLister(indexer cache.Indexer) OvnSnatRuleLister {
	return &ovnSnatRuleLister{listers.New[*kubeovnv1.OvnSnatRule](indexer, kubeovnv1.Resource("ovnsnatrule"))}
}
