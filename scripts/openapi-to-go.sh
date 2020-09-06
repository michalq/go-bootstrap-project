 #!/bin/sh

OPENAPI_GENERATOR_VERSION=v5.0.0-beta
OPENAPI_GENERATOR=go-experimental

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:${OPENAPI_GENERATOR_VERSION} generate \
    -g ${OPENAPI_GENERATOR} \
    --global-property models,modelDocs=false \
    -i /local/api/openapi.yaml \
    -o /local/pkg/api \
    -t /local/scripts/openapi