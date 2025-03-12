# Statsig Go Client
[![Release](https://img.shields.io/github/release/useless-solutions/statsig-go-client.svg?style=flat-square)](https://github.com/useless-solutions/statsig-go-client/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/useless-solutions/statsig-go-client)](https://goreportcard.com/report/github.com/useless-solutions/statsig-go-client)

This is a Go client for the Statsig API. It is code-generated from the official OpenAPI spec and intended for use primarily in the [useless-solutions/terraform-provider-statsig](https://github.com/useless-solutions/terraform-provider-statsig) project. The project uses the official [Statsig OpenAPI specification](https://docs.statsig.com/console-api/all-endpoints-generated). The current version is `20240601`.

> The Console API is versioned. Each version is guaranteed to not break existing usage; each new version introduces breaking changes. There is currently only one version: 20240601.

## Installation
To install the Statsig Go client, run the following command:

```bash
go get github.com/useless-solutions/statsig-go-client
```

## Hacks
As a workaround to get around codegen issues, we manually edited the OpenAPI schema to change the `passRate` and `rolloutRate` parameters to use a `string` type instead of an `array` type.
```json
{
  "name": "passRate",
  "required": false,
  "in": "query",
  "description": "Filter by pass rate of the gates, as determined by a sampling of overall true/false values returned: 0, 100, or INBETWEEN (pass rate greater than zero but less than 100)",
  "schema": {
    "type": "array",
    "items": {
      "type": "string",
      "enum": ["0", "100", "INBETWEEN"]
    }
  }
},
{
 "name": "rolloutRate",
 "required": false,
 "in": "query",
 "description": "Filter by rollout rate of the gates: 0 (all rules are set to pass 0%), 100 (all rules pass 100% including an \"everyone\" catch all rule), or INBETWEEN (at least one rule has a pass rate greater than 0 but less than 100)",
 "schema": {
   "type": "array",
   "items": {
     "type": "string",
     "enum": ["0", "100", "INBETWEEN"]
   }
 }
},
```

The repository is managed by [@httpoz](https://github.com/httpoz) and [@tamirarnesty](https://github.com/tamirarnesty).
