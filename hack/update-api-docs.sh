#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

# Setup at https://github.com/ahmetb/gen-crd-api-reference-docs

source $(dirname $0)/library.sh
header "updating api docs"

ensure_pandoc
ensure_vendor
make_fake_paths

export GOPATH="${FAKE_GOPATH}"
export GO111MODULE="off"

cd "${FAKE_REPOPATH}"

# EventBus
go run ${FAKE_REPOPATH}/vendor/github.com/ahmetb/gen-crd-api-reference-docs/main.go \
 -config "${FAKE_REPOPATH}/vendor/github.com/ahmetb/gen-crd-api-reference-docs/example-config.json" \
 -api-dir "github.com/argoproj-labs/argo-eventbus/pkg/apis/v1alpha1" \
 -out-file "${FAKE_REPOPATH}/api/event-bus.html" \
 -template-dir "${FAKE_REPOPATH}/hack/api-docs-template"

# Setup at https://pandoc.org/installing.html

pandoc --from markdown --to gfm ${FAKE_REPOPATH}/api/event-bus.html > ${FAKE_REPOPATH}/api/event-bus.md

