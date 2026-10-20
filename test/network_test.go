package test

import (
	"testing"

	commonv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/common/v1"
	networkv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/network/v1"
)

func TestNetworkGoldenJSON(t *testing.T) {
	for file, msg := range map[string]golden{
		"network.json":                &networkv1.Network{},
		"create_network_request.json": &networkv1.CreateNetworkRequest{},
		"subnet.json":                 &networkv1.Subnet{},
		"trunk.json":                  &networkv1.Trunk{},
		"update_trunk_request.json":   &networkv1.UpdateTrunkRequest{},
		"affinity_group.json":         &commonv1.AffinityGroup{},
	} {
		t.Run(file, func(t *testing.T) { checkGolden(t, "testdata/network/"+file, msg) })
	}
}

func TestNetworkValidation(t *testing.T) {
	network := func(n string) string { return `{"network":` + n + `}` }
	subnet := func(data string) string { return `{"subnet":{"networkId":"e5f6a7b8-c9d0-4e1f-8a2b-3c4d5e6f7a8b","data":` + data + `}}` }
	for name, tc := range map[string]struct {
		json  string
		msg   golden
		valid bool
	}{
		"create golden": {
			json: readFile(t, "testdata/network/create_network_request.json"), msg: &networkv1.CreateNetworkRequest{}, valid: true,
		},
		"vlan on provider network": {
			json:  network(`{"networkType":"NETWORK_TYPE_VLAN","providerNetwork":"physnet1","segmentationId":100}`),
			msg:   &networkv1.CreateNetworkRequest{},
			valid: true,
		},
		"flat on provider network": {
			json:  network(`{"networkType":"NETWORK_TYPE_FLAT","providerNetwork":"physnet1"}`),
			msg:   &networkv1.CreateNetworkRequest{},
			valid: true,
		},
		"no network type": {
			json: network(`{"networkResourceName":"x"}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"vlan without provider network": {
			json: network(`{"networkType":"NETWORK_TYPE_VLAN","segmentationId":100}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"vlan id out of range": {
			json: network(`{"networkType":"NETWORK_TYPE_VLAN","providerNetwork":"physnet1","segmentationId":4095}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"geneve with provider network": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE","providerNetwork":"physnet1"}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"geneve with segmentation id": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE","segmentationId":5000}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"flat with segmentation id": {
			json: network(`{"networkType":"NETWORK_TYPE_FLAT","providerNetwork":"physnet1","segmentationId":10}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"subnet ipv6": {
			json:  subnet(`{"ipVersion":"IP_VERSION_IPV6","cidr":"2001:db8::/64","gatewayIp":"2001:db8::1"}`),
			msg:   &networkv1.CreateSubnetRequest{},
			valid: true,
		},
		"subnet cidr not a prefix": {
			json: subnet(`{"ipVersion":"IP_VERSION_IPV4","cidr":"10.0.0.1"}`), msg: &networkv1.CreateSubnetRequest{},
		},
		"subnet cidr does not match version": {
			json: subnet(`{"ipVersion":"IP_VERSION_IPV6","cidr":"10.0.0.0/24"}`), msg: &networkv1.CreateSubnetRequest{},
		},
		"subnet gateway does not match version": {
			json: subnet(`{"ipVersion":"IP_VERSION_IPV4","cidr":"10.0.0.0/24","gatewayIp":"2001:db8::1"}`), msg: &networkv1.CreateSubnetRequest{},
		},
		"subnet network id not a uuid": {
			json: `{"subnet":{"networkId":"net-1","data":{"ipVersion":"IP_VERSION_IPV4","cidr":"10.0.0.0/24"}}}`, msg: &networkv1.CreateSubnetRequest{},
		},
		"affinity resource id not a uuid": {
			json: `{"port":{"networkId":"e5f6a7b8-c9d0-4e1f-8a2b-3c4d5e6f7a8b"},"affinityOrAntiAffinityConstraints":[{"type":"AFFINITY_TYPE_AFFINITY","affinityOrAntiAffinityResourceList":{"resourceId":["vm-1"]}}]}`,
			msg:  &networkv1.CreatePortRequest{},
		},
		"subnet without network": {
			json: `{"subnet":{"data":{"ipVersion":"IP_VERSION_IPV4","cidr":"10.0.0.0/24"}}}`, msg: &networkv1.CreateSubnetRequest{},
		},
		"affinity constraint without resources": {
			json: `{"port":{"networkId":"e5f6a7b8-c9d0-4e1f-8a2b-3c4d5e6f7a8b"},"affinityOrAntiAffinityConstraints":[{"type":"AFFINITY_TYPE_AFFINITY"}]}`,
			msg:  &networkv1.CreatePortRequest{},
		},
		"affinity constraint with group": {
			json:  `{"port":{"networkId":"e5f6a7b8-c9d0-4e1f-8a2b-3c4d5e6f7a8b"},"affinityOrAntiAffinityConstraints":[{"type":"AFFINITY_TYPE_AFFINITY","affinityOrAntiAffinityResourceGroupId":"d4e5f6a7-b8c9-4d0e-9f1a-2b3c4d5e6f7a"}]}`,
			msg:   &networkv1.CreatePortRequest{},
			valid: true,
		},
		"trunk update golden": {
			json: readFile(t, "testdata/network/update_trunk_request.json"), msg: &networkv1.UpdateTrunkRequest{}, valid: true,
		},
		"affinity group without type": {
			json: `{"affinityGroup":{"groupName":"g"}}`, msg: &networkv1.CreateAffinityGroupRequest{},
		},
	} {
		t.Run(name, func(t *testing.T) { checkValidation(t, tc.json, tc.msg, tc.valid) })
	}
}
