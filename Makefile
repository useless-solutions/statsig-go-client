pull-spec:
	if [ ! -d .out ]; then mkdir .out; fi
	curl -o .out/statsig-spec.yaml https://docs.statsig.com/openapi

gen-spec: pull-spec
	go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
	--config=config.yaml .out/statsig-spec.yaml
