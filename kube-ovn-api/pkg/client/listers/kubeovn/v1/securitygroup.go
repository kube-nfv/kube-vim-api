// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/DiMalovanyy/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// SecurityGroupLister helps list SecurityGroups.
// All objects returned here must be treated as read-only.
type SecurityGroupLister interface {
	// List lists all SecurityGroups in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.SecurityGroup, err error)
	// Get retrieves the SecurityGroup from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.SecurityGroup, error)
	SecurityGroupListerExpansion
}

// securityGroupLister implements the SecurityGroupLister interface.
type securityGroupLister struct {
	listers.ResourceIndexer[*kubeovnv1.SecurityGroup]
}

// NewSecurityGroupLister returns a new SecurityGroupLister.
func NewSecurityGroupLister(indexer cache.Indexer) SecurityGroupLister {
	return &securityGroupLister{listers.New[*kubeovnv1.SecurityGroup](indexer, kubeovnv1.Resource("securitygroup"))}
}
