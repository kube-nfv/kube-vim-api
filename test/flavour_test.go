package test

import (
	"testing"

	flavourv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/flavour/v1"
)

func TestFlavourGoldenJSON(t *testing.T) {
	for file, msg := range map[string]golden{
		"flavour.json":                &flavourv1.Flavour{},
		"create_flavour_request.json": &flavourv1.CreateFlavourRequest{},
		"list_flavours_response.json": &flavourv1.ListFlavoursResponse{},
	} {
		t.Run(file, func(t *testing.T) { checkGolden(t, "testdata/flavour/"+file, msg) })
	}
}

func TestFlavourValidation(t *testing.T) {
	create := func(flavour string) string { return `{"flavour":` + flavour + `}` }
	for name, tc := range map[string]struct {
		json  string
		valid bool
	}{
		"create golden": {json: readFile(t, "testdata/flavour/create_flavour_request.json"), valid: true},
		"full flavour":  {json: create(readFile(t, "testdata/flavour/flavour.json")), valid: true},
		"no cpu":        {json: create(`{"virtualMemory":{"virtualMemSize":"2Gi"}}`)},
		"no memory":     {json: create(`{"virtualCpu":{"numVirtualCpu":1}}`)},
		"zero cpus":     {json: create(`{"virtualMemory":{"virtualMemSize":"2Gi"},"virtualCpu":{"numVirtualCpu":0}}`)},
		"memory not a quantity": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2GB"},"virtualCpu":{"numVirtualCpu":1}}`),
		},
		"hugepage size not allowed": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2Gi","hugepageSize":"4Mi"},"virtualCpu":{"numVirtualCpu":1}}`),
		},
		"topology does not match cpus": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2Gi"},"virtualCpu":{"numVirtualCpu":4,"topology":{"sockets":1,"cores":2,"threads":1}}}`),
		},
		"bad oversubscription": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2Gi","virtualMemOversubscriptionPolicy":"two"},"virtualCpu":{"numVirtualCpu":1}}`),
		},
		"pinning without policy": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2Gi"},"virtualCpu":{"numVirtualCpu":1,"virtualCpuPinning":{}}}`),
		},
		"storage without type": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2Gi"},"virtualCpu":{"numVirtualCpu":1},"storageAttributes":[{"sizeOfStorage":"10Gi"}]}`),
		},
		"numa without pinning": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2Gi","numaEnabled":true,"hugepageSize":"1Gi"},"virtualCpu":{"numVirtualCpu":1}}`),
		},
		"numa without hugepages": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2Gi","numaEnabled":true},"virtualCpu":{"numVirtualCpu":1,"virtualCpuPinning":{"virtualCpuPinningPolicy":"VIRTUAL_CPU_PINNING_POLICY_DYNAMIC"}}}`),
		},
		"numa disabled without pinning": {
			json:  create(`{"virtualMemory":{"virtualMemSize":"2Gi","numaEnabled":false},"virtualCpu":{"numVirtualCpu":1}}`),
			valid: true,
		},
		"oversubscription 2:1": {
			json:  create(`{"virtualMemory":{"virtualMemSize":"2Gi","virtualMemOversubscriptionPolicy":"2:1"},"virtualCpu":{"numVirtualCpu":1}}`),
			valid: true,
		},
		"oversubscription below 1:1": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2Gi","virtualMemOversubscriptionPolicy":"1:2"},"virtualCpu":{"numVirtualCpu":1}}`),
		},
		"reserved metadata key": {
			json:  create(`{"virtualMemory":{"virtualMemSize":"2Gi"},"virtualCpu":{"numVirtualCpu":1},"metadata":{"kubevim.kubenfv.io/tier":"gold"}}`),
			valid: true,
		},
		"two boot storages": {
			json: create(`{"virtualMemory":{"virtualMemSize":"2Gi"},"virtualCpu":{"numVirtualCpu":1},"storageAttributes":[` +
				`{"typeOfStorage":"TYPE_OF_STORAGE_VOLUME","sizeOfStorage":"10Gi","isBoot":true},` +
				`{"typeOfStorage":"TYPE_OF_STORAGE_VOLUME","sizeOfStorage":"10Gi","isBoot":true}]}`),
		},
	} {
		t.Run(name, func(t *testing.T) { checkValidation(t, tc.json, &flavourv1.CreateFlavourRequest{}, tc.valid) })
	}
}
