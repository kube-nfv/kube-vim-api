PROTOC_VERSION ?= 25.1

.PHONY: help
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: proto-clean
proto-clean: ## Clean generated proto.
	rm -rf pb/gs-nfv-ifa

.PHONY: proto-compile
proto-compile: ## Compile message protobuf and gRPC service files.
	docker build \
	  --build-arg PLATFORM=$(shell uname -m) \
	  --build-arg PROTOC_VERSION=$(PROTOC_VERSION) \
	  -t protogen-image \
	  -f docker/Dockerfile docker/
	docker run --rm \
	  -v "$(PWD)/pb:/source" \
	  -w "/source" \
	  protogen-image \
	  bash -c "protoc *.proto --proto_path=. \
	  --go_out=. --go_opt=module=github.com/DiMalovanyy/kube-vim-api/pb \
	  --go-grpc_out=. --go-grpc_opt=module=github.com/DiMalovanyy/kube-vim-api/pb"
