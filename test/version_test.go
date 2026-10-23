package test

import (
	"encoding/json"
	"testing"

	computev1 "github.com/kube-nfv/kube-vim-api/pkg/apis/compute/v1"
	flavourv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/flavour/v1"
	imagev1 "github.com/kube-nfv/kube-vim-api/pkg/apis/image/v1"
	networkv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/network/v1"
	optionsv1 "github.com/kube-nfv/kube-vim-api/pkg/apis/options/v1"
	storagev1 "github.com/kube-nfv/kube-vim-api/pkg/apis/storage/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// TestApiVersion requires the API identity option and the published schema to agree.
func TestApiVersion(t *testing.T) {
	for name, fd := range map[string]protoreflect.FileDescriptor{
		"compute": computev1.File_kubenfv_kubevim_compute_v1_compute_proto,
		"flavour": flavourv1.File_kubenfv_kubevim_flavour_v1_flavour_proto,
		"image":   imagev1.File_kubenfv_kubevim_image_v1_image_proto,
		"network": networkv1.File_kubenfv_kubevim_network_v1_network_proto,
		"storage": storagev1.File_kubenfv_kubevim_storage_v1_storage_proto,
	} {
		t.Run(name, func(t *testing.T) {
			api := proto.GetExtension(fd.Options(), optionsv1.E_Api).(*optionsv1.Api)
			if api.GetName() != name {
				t.Errorf("api name = %q, want %q", api.GetName(), name)
			}
			var doc struct {
				Info struct {
					Version string `json:"version"`
				} `json:"info"`
			}
			if err := json.Unmarshal([]byte(readFile(t, "../generated/openapi/"+name+"/v1/"+name+".openapi.json")), &doc); err != nil {
				t.Fatal(err)
			}
			if doc.Info.Version != api.GetVersion() {
				t.Errorf("schema info.version = %q, api option version = %q", doc.Info.Version, api.GetVersion())
			}
		})
	}
}
