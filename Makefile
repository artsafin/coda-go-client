GOPKG=codaapi

.PHONY: all build

all: build

build:
	@mkdir -pv build
	curl -s https://coda.io/apis/v1/openapi.json -o build/openapi.in.json
	jq -f patches/fix_nextPageLink.jq < build/openapi.in.json > build/openapi.final.json
	docker run --rm -v "$(shell pwd):/data" ghcr.io/artsafin/docker-oapi-codegen:master \
	-o /data/$(GOPKG)/api.generated.go -generate client,types \
	-exclude-tags Packs \
	-package $(GOPKG) /data/build/openapi.final.json
