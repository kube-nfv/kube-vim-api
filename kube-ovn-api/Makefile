
.PHONY: generate-client
generate-client:
	docker build \
	  -t kubeovngen-image \
	  -f docker/Dockerfile docker/
	docker run --rm \
	  -v "$(PWD):/source" \
	  -w "/source" \
	  kubeovngen-image \
	  bash ./hack/update_codegen.sh
	go mod tidy

