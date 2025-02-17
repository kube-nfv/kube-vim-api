KUBE_VIM_API_URL=github.com/kube-nfv/kube-vim-api
# TODO: Get from tag
KUBE_VIM_API_VERSION=0.0.4

PROTOC_VERSION ?= 25.1
K8S_APIMACHINERY_VERSION ?= 0.31.0
GRPC_GATEWAY_VERSION ?= 2.26.1

OPENAPI_GEN_IMAGE ?= openapitools/openapi-generator-cli
OPENAPI_GEN_VERSION ?= v7.11.0

OPENAPIV2_DIR ?= openapiv2
OPENAPIV2_GEN_DIR = $(OPENAPIV2_DIR)/gen
OPENAPIV2_PYTHON_GEN_DIR = $(OPENAPIV2_GEN_DIR)/python

.PHONY: help
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: proto-clean
proto-clean: ## Clean generated proto.
	rm -rf pb/nfv

.PHONY: proto-compile
proto-compile: proto-image-build $(OPENAPIV2_DIR) ## Compile message protobuf and gRPC service files.
	docker run --rm \
	  -v "$(PWD)/pb:/source" \
	  -v "$(PWD)/$(OPENAPIV2_DIR):/api" \
	  -w "/source" \
	  protogen-image \
	  bash -c "protoc *.proto --proto_path=. \
	  --go_out=. --go_opt=module=github.com/kube-nfv/kube-vim-api/pb \
	  --go-grpc_out=. --go-grpc_opt=module=github.com/kube-nfv/kube-vim-api/pb \
	  --grpc-gateway_out=./nfv --grpc-gateway_opt paths=source_relative \
	  --openapiv2_out=/api"

.PHONY: proto-image-build
proto-image-build: ## Build docker image with all nececary dependencies to compile message protobuf and gRPC service files.
	docker build \
	  --build-arg PLATFORM=$(shell uname -m) \
	  --build-arg PROTOC_VERSION=$(PROTOC_VERSION) \
	  --build-arg K8S_APIMACHINERY_VERSION=$(K8S_APIMACHINERY_VERSION) \
	  --build-arg GRPC_GATEWAY_VERSION=$(GRPC_GATEWAY_VERSION) \
	  -t protogen-image \
	  -f docker/Dockerfile docker/

.PHONY: openapiv2-python-gen
openapiv2-python-gen: $(OPENAPIV2_PYTHON_GEN_DIR) ## Generate python stubs from the openapiv2 schema.

$(OPENAPIV2_PYTHON_GEN_DIR): $(OPENAPIV2_GEN_DIR) proto-compile
	mkdir -p $@
	docker run --rm \
	-v "$(PWD)/$(OPENAPIV2_DIR):/$(OPENAPIV2_DIR)" \
	$(OPENAPI_GEN_IMAGE):$(OPENAPI_GEN_VERSION) \
	generate -i /$(OPENAPIV2_DIR)/vi-vnfm.swagger.json \
	-g python -o /$(OPENAPIV2_PYTHON_GEN_DIR) \
	--additional-properties=packageUrl=$(KUBE_VIM_API_URL),packageVersion=$(KUBE_VIM_API_VERSION),packageName=kubevim_vivnfm_client
	bash ./hack/prepare_oapi_py_package.sh $(OPENAPIV2_PYTHON_GEN_DIR)

$(OPENAPIV2_GEN_DIR): $(OPENAPIV2_DIR)
	mkdir -p $@
