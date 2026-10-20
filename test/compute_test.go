package test

import (
	"testing"

	computev1 "github.com/kube-nfv/kube-vim-api/pkg/apis/compute/v1"
)

func TestComputeGoldenJSON(t *testing.T) {
	for file, msg := range map[string]golden{
		"compute.json":                 &computev1.Compute{},
		"create_compute_request.json":  &computev1.CreateComputeRequest{},
		"operate_compute_request.json": &computev1.OperateComputeRequest{},
		"update_compute_request.json":  &computev1.UpdateComputeRequest{},
		"migrate_compute_request.json": &computev1.MigrateComputeRequest{},
	} {
		t.Run(file, func(t *testing.T) { checkGolden(t, "testdata/compute/"+file, msg) })
	}
}

func TestComputeValidation(t *testing.T) {
	const flavour = `"flavourId":"9a8b7c6d-5e4f-4a3b-9c2d-1e0f9a8b7c6d"`
	const network = `"networkId":"e5f6a7b8-c9d0-4e1f-8a2b-3c4d5e6f7a8b"`
	create := func(rest string) string { return `{"compute":{` + flavour + `}` + rest + `}` }
	for name, tc := range map[string]struct {
		json  string
		msg   golden
		valid bool
	}{
		"create golden": {
			json: readFile(t, "testdata/compute/create_compute_request.json"), msg: &computev1.CreateComputeRequest{}, valid: true,
		},
		"create minimal": {json: create(``), msg: &computev1.CreateComputeRequest{}, valid: true},
		"no flavour": {
			json: `{"compute":{"computeName":"a"}}`, msg: &computev1.CreateComputeRequest{},
		},
		"flavour id not a uuid": {
			json: `{"compute":{"flavourId":"small"}}`, msg: &computev1.CreateComputeRequest{},
		},
		"interface by subnet only": {
			json:  create(`,"interfaceData":[{"subnetId":"4f3e2d1c-0b9a-4f8e-9d7c-6b5a4f3e2d1c"}]`),
			msg:   &computev1.CreateComputeRequest{},
			valid: true,
		},
		"masquerade requested": {
			json: create(`,"interfaceData":[{` + network + `,"typeVirtualNic":"TYPE_VIRTUAL_NIC_MASQUERADE"}]`), msg: &computev1.CreateComputeRequest{},
		},
		"interface without network, subnet or port": {
			json: create(`,"interfaceData":[{"ipAssignment":"IP_ASSIGNMENT_AUTOMATIC"}]`), msg: &computev1.CreateComputeRequest{},
		},
		"static assignment without addresses": {
			json: create(`,"interfaceData":[{` + network + `,"ipAssignment":"IP_ASSIGNMENT_STATIC"}]`), msg: &computev1.CreateComputeRequest{},
		},
		"addresses without static assignment": {
			json: create(`,"interfaceData":[{` + network + `,"ipAddress":["10.0.0.5"]}]`), msg: &computev1.CreateComputeRequest{},
		},
		"bad ip address": {
			json: create(`,"interfaceData":[{` + network + `,"ipAssignment":"IP_ASSIGNMENT_STATIC","ipAddress":["10.0.0"]}]`), msg: &computev1.CreateComputeRequest{},
		},
		"bad mac address": {
			json: create(`,"interfaceData":[{` + network + `,"macAddress":"02-00-00-aa-bb-01"}]`), msg: &computev1.CreateComputeRequest{},
		},
		"empty user data": {
			json: create(`,"userData":{"content":""}`), msg: &computev1.CreateComputeRequest{},
		},
		"operate golden": {
			json: readFile(t, "testdata/compute/operate_compute_request.json"), msg: &computev1.OperateComputeRequest{}, valid: true,
		},
		"operate without operation": {
			json: `{"computeId":"c0ffee00-1234-4abc-8def-0123456789ab"}`, msg: &computev1.OperateComputeRequest{},
		},
		"update golden": {
			json: readFile(t, "testdata/compute/update_compute_request.json"), msg: &computev1.UpdateComputeRequest{}, valid: true,
		},
		"migrate without type": {
			json: `{"computeId":"c0ffee00-1234-4abc-8def-0123456789ab"}`, msg: &computev1.MigrateComputeRequest{},
		},
		"scale flavour not a uuid": {
			json: `{"computeId":"c0ffee00-1234-4abc-8def-0123456789ab","computeFlavourId":"large"}`, msg: &computev1.ScaleComputeRequest{},
		},
		"attach volume": {
			json:  `{"computeId":"c0ffee00-1234-4abc-8def-0123456789ab","storageId":"5a6b7c8d-9e0f-4a1b-8c2d-3e4f5a6b7c8d"}`,
			msg:   &computev1.AttachVolumeRequest{},
			valid: true,
		},
	} {
		t.Run(name, func(t *testing.T) { checkValidation(t, tc.json, tc.msg, tc.valid) })
	}
}
