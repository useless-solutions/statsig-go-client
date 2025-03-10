pull-spec:
	if [ ! -d .out ]; then mkdir .out; fi
	curl -o .out/spec.json https://api.statsig.com/openapi

gen-spec:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config config.yaml .out/spec.json

build: pull-spec gen-spec
