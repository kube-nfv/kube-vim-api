KUBE_VIM_API_URL=github.com/kube-nfv/kube-vim-api

KUBE_VIM_API_VERSION ?= $(shell git describe --tags --abbrev=0)

PROTOC_VERSION ?= 25.1
K8S_APIMACHINERY_VERSION ?= 0.31.0
GRPC_GATEWAY_VERSION ?= 2.26.1

OPENAPI_GEN_IMAGE ?= openapitools/openapi-generator-cli
OPENAPI_GEN_VERSION ?= v7.11.0

OPENAPI_DIR ?= generated/openapi
GENERATED_DIR ?= generated
PYTHON_GEN_DIR = $(GENERATED_DIR)/python

LICENSE_FILE ?= LICENSE

.PHONY: all
all: proto-compile python-gen

.PHONY: help
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: proto-clean
proto-clean: ## Clean generated proto.
	rm -rf pkg/apis $(GENERATED_DIR)

.PHONY: proto-dirs
proto-dirs: ## Create proto output directories.
	mkdir -p pkg/apis/vivnfm pkg/apis/admin $(OPENAPI_DIR)

.PHONY: proto-compile
$(OPENAPI_DIR):
	mkdir -p $(OPENAPI_DIR)

proto-compile: proto-dirs proto-image-build $(OPENAPI_DIR) ## Compile message protobuf and gRPC service files.
	docker run --rm \
	  -v "$(PWD):/workspace" \
	  -v "$(PWD)/$(OPENAPI_DIR):/api" \
	  -w "/workspace/api" \
	  protogen-image \
	  bash -c "protoc common/common.proto common/types.proto --proto_path=. \
	  --go_out=/workspace --go_opt=module=$(KUBE_VIM_API_URL) && \
	  protoc vivnfm/vi-vnfm.proto --proto_path=. \
	  --go_out=/workspace --go_opt=module=$(KUBE_VIM_API_URL) \
	  --go-grpc_out=/workspace --go-grpc_opt=module=$(KUBE_VIM_API_URL) \
	  --grpc-gateway_out=/workspace --grpc-gateway_opt=module=$(KUBE_VIM_API_URL) \
	  --grpc-gateway_opt=Mvivnfm/vi-vnfm.proto=$(KUBE_VIM_API_URL)/pkg/apis/vivnfm \
	  --openapiv2_out=/api && \
	  protoc admin/kubevim-admin.proto --proto_path=. \
	  --go_out=/workspace --go_opt=module=$(KUBE_VIM_API_URL) \
	  --go-grpc_out=/workspace --go-grpc_opt=module=$(KUBE_VIM_API_URL) \
	  --grpc-gateway_out=/workspace --grpc-gateway_opt=module=$(KUBE_VIM_API_URL) \
	  --grpc-gateway_opt=Madmin/kubevim-admin.proto=$(KUBE_VIM_API_URL)/pkg/apis/admin"

.PHONY: proto-image-build
proto-image-build: ## Build docker image with all nececary dependencies to compile message protobuf and gRPC service files.
	docker build \
	  --build-arg PLATFORM=$(shell uname -m) \
	  --build-arg PROTOC_VERSION=$(PROTOC_VERSION) \
	  --build-arg K8S_APIMACHINERY_VERSION=$(K8S_APIMACHINERY_VERSION) \
	  --build-arg GRPC_GATEWAY_VERSION=$(GRPC_GATEWAY_VERSION) \
	  -t protogen-image \
	  -f docker/Dockerfile docker/

.PHONY: python-gen
python-gen: $(PYTHON_GEN_DIR) ## Generate python stubs from the openapi schema.

$(PYTHON_GEN_DIR): $(OPENAPI_DIR) proto-compile
	mkdir -p $@
	docker run --rm \
	-v "$(PWD)/$(OPENAPI_DIR):/$(OPENAPI_DIR)" \
	-v "$(PWD)/$(PYTHON_GEN_DIR):/$(PYTHON_GEN_DIR)" \
	$(OPENAPI_GEN_IMAGE):$(OPENAPI_GEN_VERSION) \
	generate -i /$(OPENAPI_DIR)/vivnfm/vi-vnfm.swagger.json \
	-g python -o /$(PYTHON_GEN_DIR) \
	--additional-properties=\
	packageUrl=$(KUBE_VIM_API_URL),\
	packageVersion=$(KUBE_VIM_API_VERSION),\
	packageName=kubevim_vivnfm_client,\
	projectDescription="Python client for VI-VNFM API"
	python3 ./hack/prepare_oapi_py_package.py $(PYTHON_GEN_DIR) $(LICENSE_FILE)
