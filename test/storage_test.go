package test

import (
	"testing"

	commonv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/common/v1"
	storagev1 "github.com/kube-nfv/kube-vim-api/pkg/apis/storage/v1"
)

func TestStorageGoldenJSON(t *testing.T) {
	for file, msg := range map[string]golden{
		"volume.json":                  &commonv1.Volume{},
		"create_volume_request.json":   &storagev1.CreateVolumeRequest{},
		"operate_volume_response.json": &storagev1.OperateVolumeResponse{},
	} {
		t.Run(file, func(t *testing.T) { checkGolden(t, "testdata/storage/"+file, msg) })
	}
}

func TestStorageValidation(t *testing.T) {
	const id = `"storageId":"a7b8c9d0-e1f2-4a3b-8c4d-5e6f7a8b9c0d"`
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
		"scale": {
			json: `{` + id + `,"newSize":"200Gi"}`, msg: &storagev1.ScaleVolumeRequest{}, valid: true,
		},
		"scale size not a quantity": {
			json: `{` + id + `,"newSize":"lots"}`, msg: &storagev1.ScaleVolumeRequest{},
		},
		"operate without operation": {
			json: `{` + id + `}`, msg: &storagev1.OperateVolumeRequest{},
		},
		"migrate without constraint": {
			json: `{` + id + `}`, msg: &storagev1.MigrateVolumeRequest{},
		},
		"update name": {
			json: `{` + id + `,"update":{"storageName":"renamed"}}`, msg: &storagev1.UpdateVolumeRequest{}, valid: true,
		},
		"get id not a uuid": {
			json: `{"storageId":"vol-1"}`, msg: &storagev1.GetVolumeRequest{},
		},
	} {
		t.Run(name, func(t *testing.T) { checkValidation(t, tc.json, tc.msg, tc.valid) })
	}
}
