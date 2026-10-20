LOCALBIN ?= $(shell pwd)/bin

OPENAPI_RAW_DIR ?= $(LOCALBIN)/openapi-raw
OPENAPI_DIR ?= generated/openapi

BUF ?= $(LOCALBIN)/buf
PROTOC_GEN_GO ?= $(LOCALBIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC ?= $(LOCALBIN)/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY ?= $(LOCALBIN)/protoc-gen-grpc-gateway
PROTOC_GEN_OPENAPIV3 ?= $(LOCALBIN)/protoc-gen-openapiv3

BUF_VERSION ?= v1.74.0
PROTOC_GEN_GO_VERSION ?= v1.36.12
PROTOC_GEN_GO_GRPC_VERSION ?= v1.6.2
GRPC_GATEWAY_VERSION ?= v2.32.0

.PHONY: all
all: generate

.PHONY: help
help: ## Display available commands.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: generate
generate: buf protoc-gen-go protoc-gen-go-grpc protoc-gen-grpc-gateway protoc-gen-openapiv3 ## Generate Go code and OpenAPI schemas from the protos.
	PATH="$(LOCALBIN):$$PATH" $(BUF) generate
	rm -rf $(OPENAPI_DIR)
	go run ./hack/openapi -in $(OPENAPI_RAW_DIR) -out $(OPENAPI_DIR)

.PHONY: lint
lint: buf ## Lint and format-check the protos.
	$(BUF) lint
	$(BUF) format --diff --exit-code

.PHONY: test
test: ## Run Go tests.
	go test ./...

.PHONY: clean
clean: ## Remove generated code and local tools.
	rm -rf pkg/apis generated $(LOCALBIN)

$(LOCALBIN):
	mkdir -p $(LOCALBIN)

.PHONY: buf
buf: $(LOCALBIN)
	@test -x $(BUF) && $(BUF) --version | grep -q $(BUF_VERSION:v%=%) || \
	GOBIN=$(LOCALBIN) go install github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION)

.PHONY: protoc-gen-go
protoc-gen-go: $(LOCALBIN)
	@test -x $(PROTOC_GEN_GO) && $(PROTOC_GEN_GO) --version | grep -q $(PROTOC_GEN_GO_VERSION) || \
	GOBIN=$(LOCALBIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@$(PROTOC_GEN_GO_VERSION)

.PHONY: protoc-gen-go-grpc
protoc-gen-go-grpc: $(LOCALBIN)
	@test -x $(PROTOC_GEN_GO_GRPC) && $(PROTOC_GEN_GO_GRPC) --version | grep -q $(PROTOC_GEN_GO_GRPC_VERSION:v%=%) || \
	GOBIN=$(LOCALBIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@$(PROTOC_GEN_GO_GRPC_VERSION)

.PHONY: protoc-gen-grpc-gateway
protoc-gen-grpc-gateway: $(LOCALBIN)
	@test -x $(PROTOC_GEN_GRPC_GATEWAY) && $(PROTOC_GEN_GRPC_GATEWAY) --version | grep -q $(GRPC_GATEWAY_VERSION) || \
	GOBIN=$(LOCALBIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@$(GRPC_GATEWAY_VERSION)

# protoc-gen-openapiv3 has no version flag; a marker file records the installed version.
.PHONY: protoc-gen-openapiv3
protoc-gen-openapiv3: $(LOCALBIN)
	@test -x $(PROTOC_GEN_OPENAPIV3) && test -f $(LOCALBIN)/.protoc-gen-openapiv3-$(GRPC_GATEWAY_VERSION) || \
	{ GOBIN=$(LOCALBIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv3@$(GRPC_GATEWAY_VERSION) && \
	rm -f $(LOCALBIN)/.protoc-gen-openapiv3-* && touch $(LOCALBIN)/.protoc-gen-openapiv3-$(GRPC_GATEWAY_VERSION); }
