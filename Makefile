pull-spec:
	if [ ! -d docs ]; then mkdir docs; fi
	curl -o docs/spec.json https://api.statsig.com/openapi/20240601.json

gen-spec:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config config.yaml docs/spec.json

build: pull-spec gen-spec
