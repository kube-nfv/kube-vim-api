// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	listers "k8s.io/client-go/listers"
	cache "k8s.io/client-go/tools/cache"
)

// VpcNatGatewayLister helps list VpcNatGateways.
// All objects returned here must be treated as read-only.
type VpcNatGatewayLister interface {
	// List lists all VpcNatGateways in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*kubeovnv1.VpcNatGateway, err error)
	// Get retrieves the VpcNatGateway from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*kubeovnv1.VpcNatGateway, error)
	VpcNatGatewayListerExpansion
}

// vpcNatGatewayLister implements the VpcNatGatewayLister interface.
type vpcNatGatewayLister struct {
	listers.ResourceIndexer[*kubeovnv1.VpcNatGateway]
}

// NewVpcNatGatewayLister returns a new VpcNatGatewayLister.
func NewVpcNatGatewayLister(indexer cache.Indexer) VpcNatGatewayLister {
	return &vpcNatGatewayLister{listers.New[*kubeovnv1.VpcNatGateway](indexer, kubeovnv1.Resource("vpcnatgateway"))}
}
