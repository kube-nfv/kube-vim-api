package test

import (
	"testing"

	commonv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/common/v1"
	storagev1 "github.com/kube-nfv/kube-vim-api/pkg/apis/storage/v1"
)

func TestStorageGoldenJSON(t *testing.T) {
	for file, msg := range map[string]golden{
		"volume.json":                &commonv1.Volume{},
		"create_volume_request.json": &storagev1.CreateVolumeRequest{},
		"update_volume_request.json": &storagev1.UpdateVolumeRequest{},
		"volume_snapshot.json":       &storagev1.VolumeSnapshot{},
	} {
		t.Run(file, func(t *testing.T) { checkGolden(t, "testdata/storage/"+file, msg) })
	}
}

func TestStorageValidation(t *testing.T) {
	const id = `"volumeId":"a7b8c9d0-e1f2-4a3b-8c4d-5e6f7a8b9c0d"`
	volume := func(data string) string { return `{"volume":{"storageData":` + data + `}}` }
	for name, tc := range map[string]struct {
		json  string
		msg   golden
		valid bool
	}{
		"create golden": {
			json: readFile(t, "testdata/storage/create_volume_request.json"), msg: &storagev1.CreateVolumeRequest{}, valid: true,
		},
		"no storage data": {
			json: `{"volume":{"volumeName":"a"}}`, msg: &storagev1.CreateVolumeRequest{},
		},
		"ephemeral volume": {
			json: volume(`{"typeOfStorage":"TYPE_OF_STORAGE_EPHEMERAL","sizeOfStorage":"10Gi"}`), msg: &storagev1.CreateVolumeRequest{},
		},
		"boot volume": {
			json: volume(`{"typeOfStorage":"TYPE_OF_STORAGE_VOLUME","sizeOfStorage":"10Gi","isBoot":true}`), msg: &storagev1.CreateVolumeRequest{},
		},
		"size not a quantity": {
			json: volume(`{"typeOfStorage":"TYPE_OF_STORAGE_VOLUME","sizeOfStorage":"10GB"}`), msg: &storagev1.CreateVolumeRequest{},
		},
		"storage class not a name": {
			json: volume(`{"typeOfStorage":"TYPE_OF_STORAGE_VOLUME","sizeOfStorage":"10Gi","storageClass":"Fast_NVMe"}`), msg: &storagev1.CreateVolumeRequest{},
		},
		"scale": {
			json: `{` + id + `,"newSize":"200Gi"}`, msg: &storagev1.ScaleVolumeRequest{}, valid: true,
		},
		"scale size not a quantity": {
			json: `{` + id + `,"newSize":"lots"}`, msg: &storagev1.ScaleVolumeRequest{},
		},
		"migrate": {
			json: `{` + id + `,"storageClass":"ceph-rbd"}`, msg: &storagev1.MigrateVolumeRequest{}, valid: true,
		},
		"migrate without storage class": {
			json: `{` + id + `}`, msg: &storagev1.MigrateVolumeRequest{},
		},
		"update golden": {
			json: readFile(t, "testdata/storage/update_volume_request.json"), msg: &storagev1.UpdateVolumeRequest{}, valid: true,
		},
		"get id not a uuid": {
			json: `{"volumeId":"vol-1"}`, msg: &storagev1.GetVolumeRequest{},
		},
		"create snapshot": {
			json: `{` + id + `,"snapshot":{"snapshotName":"nightly"}}`, msg: &storagev1.CreateVolumeSnapshotRequest{}, valid: true,
		},
		"revert snapshot id not a uuid": {
			json: `{` + id + `,"snapshotId":"snap-1"}`, msg: &storagev1.RevertVolumeSnapshotRequest{},
		},
		"list snapshots of one volume": {
			json: `{` + id + `}`, msg: &storagev1.ListVolumeSnapshotsRequest{}, valid: true,
		},
		"list snapshots of all volumes": {
			json: `{"volumeId":"-"}`, msg: &storagev1.ListVolumeSnapshotsRequest{}, valid: true,
		},
		"list snapshots volume id not a uuid": {
			json: `{"volumeId":"vol-1"}`, msg: &storagev1.ListVolumeSnapshotsRequest{},
		},
	} {
		t.Run(name, func(t *testing.T) { checkValidation(t, tc.json, tc.msg, tc.valid) })
	}
}
