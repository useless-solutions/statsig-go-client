//go:build tools
// +build tools

package main

//go:generate go run -v github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=../config.yaml ../.out/spec.json
import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)
