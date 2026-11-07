FROM golang:1.25

ARG PLATFORM
ARG PROTOC_VERSION

RUN apt-get update && apt-get install -y unzip

RUN mkdir -p ${GOPATH}/src/k8s.io/code-generator
RUN git clone --depth 1 --branch v0.35.0-alpha.1 https://github.com/kubernetes/code-generator.git ${GOPATH}/src/k8s.io/code-generator
