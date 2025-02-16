FROM golang:1.23

ARG PLATFORM
ARG PROTOC_VERSION
ARG K8S_APIMACHINERY_VERSION
ARG GRPC_GATEWAY_VERSION

RUN apt-get update && apt-get install -y unzip

# By default Intel chipset (x86_64) is assumed but if the host device is an Apple
# silicon (arm) chipset based then a relevant (aarch_64) release file is used.

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.1
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.3.0
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.25.0
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.25.0

# Download protoc and well-known types
RUN export ZIP=x86_64 && \
               if [ ${PLATFORM} = "arm64" ]; then export ZIP=aarch_64; fi && \
               wget --quiet https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-${ZIP}.zip && \
               unzip -o protoc-${PROTOC_VERSION}-linux-${ZIP}.zip -d /usr/local bin/protoc && \
               unzip -o protoc-${PROTOC_VERSION}-linux-${ZIP}.zip -d /usr/local 'include/*' && \
               rm -rf protoc-${PROTOC_VERSION}-linux-${ZIP}.zip
RUN wget --quiet https://github.com/kubernetes/apimachinery/archive/refs/tags/v${K8S_APIMACHINERY_VERSION}.zip && \
               mkdir -p /usr/local/include/k8s.io && \
               unzip -o v${K8S_APIMACHINERY_VERSION}.zip -d /usr/local/include/k8s.io '*.proto' && \
               mv /usr/local/include/k8s.io/apimachinery-${K8S_APIMACHINERY_VERSION} /usr/local/include/k8s.io/apimachinery && \
               rm -rf v${K8S_APIMACHINERY_VERSION}.zip
RUN wget --quiet -P /usr/local/include/google/api \
                https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto \
                https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto \
                https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/field_behavior.proto \
                https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/httpbody.proto \
                https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/visibility.proto
RUN wget --quiet https://github.com/grpc-ecosystem/grpc-gateway/archive/refs/tags/v${GRPC_GATEWAY_VERSION}.zip && \
                mkdir -p /usr/local/include/protoc-gen-openapiv2 && \
                unzip -o v${GRPC_GATEWAY_VERSION}.zip -d /usr/local/include/protoc-gen-openapiv2 '**/protoc-gen-openapiv2/options/*.proto' && \
                mv /usr/local/include/protoc-gen-openapiv2/grpc-gateway-${GRPC_GATEWAY_VERSION}/protoc-gen-openapiv2/* /usr/local/include/protoc-gen-openapiv2 && \
                rm -rf v${GRPC_GATEWAY_VERSION}.zip
