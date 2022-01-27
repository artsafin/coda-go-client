.PHONY: all

all: build

build:
	docker run --rm -v "$(shell pwd):/data" ghcr.io/artsafin/docker-oapi-codegen:master \
	-o /data/api.generated.go -generate client,types \
	-exclude-tags Packs -package codaapi https://coda.io/apis/v1/openapi.yaml