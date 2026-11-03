package test

import (
	"testing"

	networkv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/network/v1"
)

func TestNetworkGoldenJSON(t *testing.T) {
	for file, msg := range map[string]golden{
		"network.json":                &networkv1.Network{},
		"create_network_request.json": &networkv1.CreateNetworkRequest{},
		"subnet.json":                 &networkv1.Subnet{},
		"update_network_request.json": &networkv1.UpdateNetworkRequest{},
	} {
		t.Run(file, func(t *testing.T) { checkGolden(t, "testdata/network/"+file, msg) })
	}
}

func TestNetworkValidation(t *testing.T) {
	network := func(n string) string { return `{"network":` + n + `}` }
	subnet := func(data string) string {
		return `{"subnet":{"networkId":"e5f6a7b8-c9d0-4e1f-8a2b-3c4d5e6f7a8b","data":` + data + `}}`
	}
	const v4 = `{"ipVersion":"IP_VERSION_IPV4","cidr":"10.0.0.0/24"}`
	const v6 = `{"ipVersion":"IP_VERSION_IPV6","cidr":"fd00::/64"}`
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
		"vlan with sriov pool": {
			json:  network(`{"networkType":"NETWORK_TYPE_VLAN","providerNetwork":"physnet1","segmentationId":100,"sriov":{"pool":"physnet1_vfio"}}`),
			msg:   &networkv1.CreateNetworkRequest{},
			valid: true,
		},
		"sriov without pool": {
			json: network(`{"networkType":"NETWORK_TYPE_FLAT","providerNetwork":"physnet1","sriov":{}}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"vlan with an external gateway": {
			json:  network(`{"networkType":"NETWORK_TYPE_VLAN","providerNetwork":"physnet1","segmentationId":100,"layer3Attributes":[{"ipVersion":"IP_VERSION_IPV4","cidr":"10.0.0.0/24","gatewayMode":"GATEWAY_MODE_EXTERNAL"}]}`),
			msg:   &networkv1.CreateNetworkRequest{},
			valid: true,
		},
		"geneve with a gateway mode": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE","layer3Attributes":[{"ipVersion":"IP_VERSION_IPV4","cidr":"10.0.0.0/24","gatewayMode":"GATEWAY_MODE_VIM"}]}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"geneve with sriov": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE","sriov":{"pool":"physnet1_vfio"},"layer3Attributes":[` + v4 + `]}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"no network type": {
			json: network(`{"networkName":"x"}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"vlan without provider network": {
			json: network(`{"networkType":"NETWORK_TYPE_VLAN","segmentationId":100}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"vlan id out of range": {
			json: network(`{"networkType":"NETWORK_TYPE_VLAN","providerNetwork":"physnet1","segmentationId":4095}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"geneve with provider network": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE","providerNetwork":"physnet1","layer3Attributes":[` + v4 + `]}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"geneve with segmentation id": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE","segmentationId":5000,"layer3Attributes":[` + v4 + `]}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"geneve without subnet": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE"}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"flat with segmentation id": {
			json: network(`{"networkType":"NETWORK_TYPE_FLAT","providerNetwork":"physnet1","segmentationId":10}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"dual stack": {
			json:  network(`{"networkType":"NETWORK_TYPE_GENEVE","layer3Attributes":[` + v4 + `,` + v6 + `]}`),
			msg:   &networkv1.CreateNetworkRequest{},
			valid: true,
		},
		"two subnets of one version": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE","layer3Attributes":[` + v4 + `,{"ipVersion":"IP_VERSION_IPV4","cidr":"10.1.0.0/24"}]}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"three subnets": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE","layer3Attributes":[` + v4 + `,` + v6 + `,` + v4 + `]}`), msg: &networkv1.CreateNetworkRequest{},
		},
		"mtu too small": {
			json: network(`{"networkType":"NETWORK_TYPE_GENEVE","mtu":60,"layer3Attributes":[` + v4 + `]}`), msg: &networkv1.CreateNetworkRequest{},
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
		"subnet dns server not an ip": {
			json: subnet(`{"ipVersion":"IP_VERSION_IPV4","cidr":"10.0.0.0/24","dnsServers":["dns.example.com"]}`), msg: &networkv1.CreateSubnetRequest{},
		},
		"subnet network id not a uuid": {
			json: `{"subnet":{"networkId":"net-1","data":` + v4 + `}}`, msg: &networkv1.CreateSubnetRequest{},
		},
		"subnet without network": {
			json: `{"subnet":{"data":` + v4 + `}}`, msg: &networkv1.CreateSubnetRequest{},
		},
		"update golden": {
			json: readFile(t, "testdata/network/update_network_request.json"), msg: &networkv1.UpdateNetworkRequest{}, valid: true,
		},
		"update empty name": {
			json: `{"networkId":"3e2d1c0b-9a8f-4e7d-8c6b-5a4f3e2d1c0b","update":{"networkName":""}}`, msg: &networkv1.UpdateNetworkRequest{},
		},
	} {
		t.Run(name, func(t *testing.T) { checkValidation(t, tc.json, tc.msg, tc.valid) })
	}
}
