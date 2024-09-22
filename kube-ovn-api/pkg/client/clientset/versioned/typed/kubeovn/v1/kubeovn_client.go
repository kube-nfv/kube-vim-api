// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	http "net/http"

	kubeovnv1 "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/apis/kubeovn/v1"
	scheme "github.com/kube-nfv/kube-vim-api/kube-ovn-api/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type KubeovnV1Interface interface {
	RESTClient() rest.Interface
	IPsGetter
	IPPoolsGetter
	IptablesDnatRulesGetter
	IptablesEIPsGetter
	IptablesFIPRulesGetter
	IptablesSnatRulesGetter
	OvnDnatRulesGetter
	OvnEipsGetter
	OvnFipsGetter
	OvnSnatRulesGetter
	ProviderNetworksGetter
	QoSPoliciesGetter
	SecurityGroupsGetter
	SubnetsGetter
	SwitchLBRulesGetter
	VipsGetter
	VlansGetter
	VpcsGetter
	VpcDnsesGetter
	VpcNatGatewaysGetter
}

// KubeovnV1Client is used to interact with features provided by the kubeovn.io group.
type KubeovnV1Client struct {
	restClient rest.Interface
}

func (c *KubeovnV1Client) IPs() IPInterface {
	return newIPs(c)
}

func (c *KubeovnV1Client) IPPools() IPPoolInterface {
	return newIPPools(c)
}

func (c *KubeovnV1Client) IptablesDnatRules() IptablesDnatRuleInterface {
	return newIptablesDnatRules(c)
}

func (c *KubeovnV1Client) IptablesEIPs() IptablesEIPInterface {
	return newIptablesEIPs(c)
}

func (c *KubeovnV1Client) IptablesFIPRules() IptablesFIPRuleInterface {
	return newIptablesFIPRules(c)
}

func (c *KubeovnV1Client) IptablesSnatRules() IptablesSnatRuleInterface {
	return newIptablesSnatRules(c)
}

func (c *KubeovnV1Client) OvnDnatRules() OvnDnatRuleInterface {
	return newOvnDnatRules(c)
}

func (c *KubeovnV1Client) OvnEips() OvnEipInterface {
	return newOvnEips(c)
}

func (c *KubeovnV1Client) OvnFips() OvnFipInterface {
	return newOvnFips(c)
}

func (c *KubeovnV1Client) OvnSnatRules() OvnSnatRuleInterface {
	return newOvnSnatRules(c)
}

func (c *KubeovnV1Client) ProviderNetworks() ProviderNetworkInterface {
	return newProviderNetworks(c)
}

func (c *KubeovnV1Client) QoSPolicies() QoSPolicyInterface {
	return newQoSPolicies(c)
}

func (c *KubeovnV1Client) SecurityGroups() SecurityGroupInterface {
	return newSecurityGroups(c)
}

func (c *KubeovnV1Client) Subnets() SubnetInterface {
	return newSubnets(c)
}

func (c *KubeovnV1Client) SwitchLBRules() SwitchLBRuleInterface {
	return newSwitchLBRules(c)
}

func (c *KubeovnV1Client) Vips() VipInterface {
	return newVips(c)
}

func (c *KubeovnV1Client) Vlans() VlanInterface {
	return newVlans(c)
}

func (c *KubeovnV1Client) Vpcs() VpcInterface {
	return newVpcs(c)
}

func (c *KubeovnV1Client) VpcDnses() VpcDnsInterface {
	return newVpcDnses(c)
}

func (c *KubeovnV1Client) VpcNatGateways() VpcNatGatewayInterface {
	return newVpcNatGateways(c)
}

// NewForConfig creates a new KubeovnV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*KubeovnV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new KubeovnV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*KubeovnV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &KubeovnV1Client{client}, nil
}

// NewForConfigOrDie creates a new KubeovnV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KubeovnV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KubeovnV1Client for the given RESTClient.
func New(c rest.Interface) *KubeovnV1Client {
	return &KubeovnV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := kubeovnv1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KubeovnV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
