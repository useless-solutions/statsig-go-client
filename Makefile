pull-spec:
	if [ ! -d .out ]; then mkdir .out; fi
	curl -o .out/spec.json https://docs.statsig.com/openapi

gen-spec: pull-spec
	go generate tools/tools.go
