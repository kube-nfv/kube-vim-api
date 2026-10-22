package test

import (
	"testing"

	imagev1 "github.com/kube-nfv/kube-vim-api/pkg/apis/image/v1"
)

func TestImageGoldenJSON(t *testing.T) {
	for file, msg := range map[string]golden{
		"image.json":                       &imagev1.Image{},
		"create_image_request.json":        &imagev1.CreateImageRequest{},
		"update_image_request.json":        &imagev1.UpdateImageRequest{},
		"list_images_response.json":        &imagev1.ListImagesResponse{},
		"list_image_sources_response.json": &imagev1.ListImageSourcesResponse{},
	} {
		t.Run(file, func(t *testing.T) { checkGolden(t, "testdata/image/"+file, msg) })
	}
}

func TestImageValidation(t *testing.T) {
	const id = `"imageId":"11111111-2222-4333-8444-555555555555"`
	for name, tc := range map[string]struct {
		json  string
		msg   golden
		valid bool
	}{
		"create golden": {
			json: readFile(t, "testdata/image/create_image_request.json"), msg: &imagev1.CreateImageRequest{}, valid: true,
		},
		"name and source only": {
			json: `{"image":{"name":"a","source":{"upload":{}}}}`, msg: &imagev1.CreateImageRequest{}, valid: true,
		},
		"http with inline credentials": {
			json:  `{"image":{"name":"a","source":{"http":{"url":"https://example.com/a.img","basicAuth":{"username":"u","password":"s"}}}}}`,
			msg:   &imagev1.CreateImageRequest{},
			valid: true,
		},
		"no source": {
			json: `{"image":{"name":"a"}}`, msg: &imagev1.CreateImageRequest{},
		},
		"empty source": {
			json: `{"image":{"name":"a","source":{}}}`, msg: &imagev1.CreateImageRequest{},
		},
		"missing name": {
			json: `{"image":{"source":{"upload":{}}}}`, msg: &imagev1.CreateImageRequest{},
		},
		"min disk not a quantity": {
			json: `{"image":{"name":"a","minDisk":"10GB","source":{"upload":{}}}}`, msg: &imagev1.CreateImageRequest{},
		},
		"blank size zero": {
			json: `{"image":{"name":"a","source":{"blank":{"size":"0Gi"}}}}`, msg: &imagev1.CreateImageRequest{},
		},
		"request id not a uuid": {
			json: `{"image":{"name":"a","source":{"upload":{}}},"requestId":"abc"}`, msg: &imagev1.CreateImageRequest{},
		},
		"http url not a uri": {
			json: `{"image":{"name":"a","source":{"http":{"url":"not a url"}}}}`, msg: &imagev1.CreateImageRequest{},
		},
		"unknown firmware": {
			json: `{"image":{"name":"a","source":{"upload":{}},"firmware":7}}`, msg: &imagev1.CreateImageRequest{},
		},
		"update golden": {
			json: readFile(t, "testdata/image/update_image_request.json"), msg: &imagev1.UpdateImageRequest{}, valid: true,
		},
		"update nothing": {
			json: `{` + id + `,"update":{}}`, msg: &imagev1.UpdateImageRequest{}, valid: true,
		},
		"update empty name": {
			json: `{` + id + `,"update":{"name":""}}`, msg: &imagev1.UpdateImageRequest{},
		},
		"update id not a uuid": {
			json: `{"imageId":"img-1","update":{"name":"a"}}`, msg: &imagev1.UpdateImageRequest{},
		},
	} {
		t.Run(name, func(t *testing.T) { checkValidation(t, tc.json, tc.msg, tc.valid) })
	}
}
