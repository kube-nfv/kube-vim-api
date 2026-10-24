package test

import (
	"strings"
	"testing"

	commonv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/common/v1"
)

func TestQuantityRule(t *testing.T) {
	for size, valid := range map[string]bool{
		"10Gi":  true,
		"500M":  true,
		"1.5Gi": true,
		"1e3":   true,
		"0.5Gi": true,
		"-1Gi":  false,
		"+1Gi":  false,
		"0":     false,
		"0Gi":   false,
		"0.0":   false,
		"10GB":  false,
		"":      false,
	} {
		t.Run(size, func(t *testing.T) {
			data := `{"typeOfStorage":"TYPE_OF_STORAGE_VOLUME","sizeOfStorage":"` + size + `"}`
			checkValidation(t, data, &commonv1.VirtualStorageData{}, valid)
		})
	}
}

func TestMetadataRule(t *testing.T) {
	volume := func(metadata string) string {
		return `{"storageData":{"typeOfStorage":"TYPE_OF_STORAGE_VOLUME","sizeOfStorage":"1Gi"},"metadata":` + metadata + `}`
	}
	many := make([]string, 65)
	for i := range many {
		many[i] = `"k` + strings.Repeat("x", i) + `":"v"`
	}
	for name, tc := range map[string]struct {
		metadata string
		valid    bool
	}{
		"consumer keys":      {`{"osm.etsi.org/vdu":"router","owner":"team-a"}`, true},
		"kube-vim prefix":    {`{"kubevim.kubenfv.io/role":"edge"}`, true},
		"too many entries":   {`{` + strings.Join(many, ",") + `}`, false},
		"empty key":          {`{"":"v"}`, false},
		"key too long":       {`{"` + strings.Repeat("k", 257) + `":"v"}`, false},
		"value too long":     {`{"k":"` + strings.Repeat("v", 2049) + `"}`, false},
		"value at the limit": {`{"k":"` + strings.Repeat("v", 2048) + `"}`, true},
	} {
		t.Run(name, func(t *testing.T) { checkValidation(t, volume(tc.metadata), &commonv1.Volume{}, tc.valid) })
	}
}
