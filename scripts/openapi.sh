#!/bin/bash
set -e

readonly service="$1"
readonly output_dir="$2"
readonly package="$3"
readonly external_client="$4" # true or false
# the rest of the arguments that are passed to oapi-codegen
readonly oapi_codegen_args="${@:5}"

# usage checks
if [ -z "$service" ]; then
  echo "Usage: $0 SERVICE_NAME OUTPUT_DIR PACKAGE_NAME [EXTERNAL_CLIENT]"
  exit 1
fi

# if not external client
if [ "$external_client" != "true" ]; then
    oapi-codegen -generate types -o "$output_dir/openapi_types.gen.go" -package "$package" "api/openapi/$service.yaml" "$oapi_codegen_args"
    oapi-codegen -generate chi-server -o "$output_dir/openapi_api.gen.go" -package "$package" "api/openapi/$service.yaml" "$oapi_codegen_args"
    oapi-codegen -generate types -o "internal/common/client/$service/openapi_types.gen.go" -package "$service" "api/openapi/$service.yaml" "$oapi_codegen_args"
    oapi-codegen -generate client -o "internal/common/client/$service/openapi_client_gen.go" -package "$service" "api/openapi/$service.yaml" "$oapi_codegen_args"
else
    oapi-codegen -generate types -o "$output_dir/$service/openapi_types.gen.go" -package "$package" "api/clients/$service.yaml" "$oapi_codegen_args"
    oapi-codegen -generate client -o "$output_dir/$service/openapi_client_gen.go" -package "$package" "api/clients/$service.yaml" "$oapi_codegen_args"
fi
