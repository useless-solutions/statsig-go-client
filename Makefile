pull-spec:
	if [ ! -d .out ]; then mkdir .out; fi
	curl -o .out/spec.json https://api.statsig.com/openapi

gen-spec: pull-spec
	go generate tools/tools.go

gen-spec-2: pull-spec
	openapi-generator generate -i .out/spec.json -g go -o client --skip-validate-spec
