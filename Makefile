

GO_INSTALL := ./scripts/go_install.sh

CONTROLLER_GEN := $(abspath ./bin/controller-gen)
CONVERSION_GEN := $(abspath ./bin/conversion-gen)

.PHONY: test
test:
	go test ./...


.PHONY: generate
generate:
	$(MAKE) generate-go-deepcopy generate-go-conversions

.PHONY: generate-go-deepcopy
generate-go-deepcopy: $(CONTROLLER_GEN)
	$(CONTROLLER_GEN) object paths=./api/...


.PHONY: generate-go-conversions
generate-go-conversions: $(CONVERSION_GEN)
	$(CONVERSION_GEN) --output-file=zz_generated.conversion.go ./api/v1beta1 --v=3


# Build controller-gen & conversion-gen

.PHONY: controller-gen
controller-gen: $(CONTROLLER_GEN)

$(CONTROLLER_GEN):
	GOBIN=$(abspath ./bin) $(GO_INSTALL) sigs.k8s.io/controller-tools/cmd/controller-gen controller-gen v0.18.0

.PHONY: conversion-gen
conversion-gen: $(CONVERSION_GEN)

.PHONY: $(CONVERSION_GEN)
$(CONVERSION_GEN):
	GOBIN=$(abspath ./bin) $(GO_INSTALL) k8s.io/code-generator/cmd/conversion-gen conversion-gen v0.33.0
