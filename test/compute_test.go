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
		"compute_snapshot.json":        &computev1.ComputeSnapshot{},
		"placement_group.json":         &computev1.PlacementGroup{},
	} {
		t.Run(file, func(t *testing.T) { checkGolden(t, "testdata/compute/"+file, msg) })
	}
}

func TestComputeValidation(t *testing.T) {
	const ids = `"flavourId":"9a8b7c6d-5e4f-4a3b-9c2d-1e0f9a8b7c6d","imageId":"5f0c9d2e-3a4b-4c1d-9e8f-0a1b2c3d4e5f"`
	const network = `"networkId":"e5f6a7b8-c9d0-4e1f-8a2b-3c4d5e6f7a8b"`
	const compute = `"computeId":"c0ffee00-1234-4abc-8def-0123456789ab"`
	create := func(rest string) string { return `{"compute":{` + ids + `}` + rest + `}` }
	iface := func(fields string) string { return create(`,"networkInterfaces":[{` + network + fields + `}]`) }
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
			json: `{"compute":{"imageId":"5f0c9d2e-3a4b-4c1d-9e8f-0a1b2c3d4e5f"}}`, msg: &computev1.CreateComputeRequest{},
		},
		"no image": {
			json: `{"compute":{"flavourId":"9a8b7c6d-5e4f-4a3b-9c2d-1e0f9a8b7c6d"}}`, msg: &computev1.CreateComputeRequest{},
		},
		"flavour id not a uuid": {
			json: `{"compute":{"flavourId":"small","imageId":"5f0c9d2e-3a4b-4c1d-9e8f-0a1b2c3d4e5f"}}`, msg: &computev1.CreateComputeRequest{},
		},
		"placement group not a uuid": {
			json: `{"compute":{` + ids + `,"placementGroupIds":["ha"]}}`, msg: &computev1.CreateComputeRequest{},
		},
		"interface without network": {
			json: create(`,"networkInterfaces":[{"name":"data0"}]`), msg: &computev1.CreateComputeRequest{},
		},
		"masquerade requested": {
			json: iface(`,"typeVirtualNic":"TYPE_VIRTUAL_NIC_MASQUERADE"`), msg: &computev1.CreateComputeRequest{},
		},
		"interface name not a label": {
			json: iface(`,"name":"Data_0"`), msg: &computev1.CreateComputeRequest{},
		},
		"duplicate interface names": {
			json: create(`,"networkInterfaces":[{` + network + `,"name":"data0"},{` + network + `,"name":"data0"}]`), msg: &computev1.CreateComputeRequest{},
		},
		"two unnamed interfaces on one network": {
			json:  create(`,"networkInterfaces":[{` + network + `},{` + network + `}]`),
			msg:   &computev1.CreateComputeRequest{},
			valid: true,
		},
		"static assignment without addresses": {
			json: iface(`,"ipAssignment":"IP_ASSIGNMENT_STATIC"`), msg: &computev1.CreateComputeRequest{},
		},
		"addresses without static assignment": {
			json: iface(`,"ipAddresses":["10.0.0.5"]`), msg: &computev1.CreateComputeRequest{},
		},
		"two ipv4 addresses": {
			json: iface(`,"ipAssignment":"IP_ASSIGNMENT_STATIC","ipAddresses":["10.0.0.5","10.0.0.6"]`), msg: &computev1.CreateComputeRequest{},
		},
		"bad ip address": {
			json: iface(`,"ipAssignment":"IP_ASSIGNMENT_STATIC","ipAddresses":["10.0.0"]`), msg: &computev1.CreateComputeRequest{},
		},
		"bad mac address": {
			json: iface(`,"macAddress":"02-00-00-aa-bb-01"`), msg: &computev1.CreateComputeRequest{},
		},
		"empty user data": {
			json: create(`,"userData":{"content":""}`), msg: &computev1.CreateComputeRequest{},
		},
		"operate golden": {
			json: readFile(t, "testdata/compute/operate_compute_request.json"), msg: &computev1.OperateComputeRequest{}, valid: true,
		},
		"operate without operation": {
			json: `{` + compute + `}`, msg: &computev1.OperateComputeRequest{},
		},
		"update golden": {
			json: readFile(t, "testdata/compute/update_compute_request.json"), msg: &computev1.UpdateComputeRequest{}, valid: true,
		},
		"scale flavour not a uuid": {
			json: `{` + compute + `,"flavourId":"large"}`, msg: &computev1.ScaleComputeRequest{},
		},
		"attach volume": {
			json:  `{` + compute + `,"volumeId":"5a6b7c8d-9e0f-4a1b-8c2d-3e4f5a6b7c8d"}`,
			msg:   &computev1.AttachVolumeRequest{},
			valid: true,
		},
		"create snapshot": {
			json:  `{` + compute + `,"snapshot":{"snapshotName":"before-upgrade"}}`,
			msg:   &computev1.CreateComputeSnapshotRequest{},
			valid: true,
		},
		"revert snapshot id not a uuid": {
			json: `{` + compute + `,"snapshotId":"snap-1"}`, msg: &computev1.RevertComputeSnapshotRequest{},
		},
		"list snapshots of one compute": {
			json: `{` + compute + `}`, msg: &computev1.ListComputeSnapshotsRequest{}, valid: true,
		},
		"list snapshots of all computes": {
			json: `{"computeId":"-"}`, msg: &computev1.ListComputeSnapshotsRequest{}, valid: true,
		},
		"list snapshots compute id not a uuid": {
			json: `{"computeId":"vm-1"}`, msg: &computev1.ListComputeSnapshotsRequest{},
		},
		"get snapshot of all computes": {
			json: `{"computeId":"-","snapshotId":"5a6b7c8d-9e0f-4a1b-8c2d-3e4f5a6b7c8d"}`, msg: &computev1.GetComputeSnapshotRequest{},
		},
		"update adds interfaces with the same name": {
			json: `{` + compute + `,"update":{"networkInterfacesToAdd":[{"name":"data","networkId":"3e2d1c0b-9a8f-4e7d-8c6b-5a4f3e2d1c0b"},{"name":"data","networkId":"3e2d1c0b-9a8f-4e7d-8c6b-5a4f3e2d1c0b"}]}}`,
			msg:  &computev1.UpdateComputeRequest{},
		},
		"delete with volumes": {
			json: `{` + compute + `,"deleteVolumes":true}`, msg: &computev1.DeleteComputeRequest{}, valid: true,
		},
		"placement group": {
			json:  `{"placementGroup":{"type":"PLACEMENT_TYPE_ANTI_AFFINITY","enforcement":"PLACEMENT_ENFORCEMENT_PREFERRED"}}`,
			msg:   &computev1.CreatePlacementGroupRequest{},
			valid: true,
		},
		"placement group without type": {
			json: `{"placementGroup":{"placementGroupName":"g"}}`, msg: &computev1.CreatePlacementGroupRequest{},
		},
	} {
		t.Run(name, func(t *testing.T) { checkValidation(t, tc.json, tc.msg, tc.valid) })
	}
}
