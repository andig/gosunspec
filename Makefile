# Convenience targets for working with the SunSpec model definitions.

GO            ?= go
SPEC_JSON_DIR ?= spec/json
SPEC_XML_DIR  ?= spec/smdx
MODELS_DIR    ?= models

.PHONY: spec spec-xml build test help

help:
	@echo "Targets:"
	@echo "  spec        Regenerate Go model packages from the JSON spec into $(MODELS_DIR)/"
	@echo "  spec-xml    Regenerate Go model packages from the legacy XML/SMDX spec into $(MODELS_DIR)/"
	@echo "  build       go build ./..."
	@echo "  test        go test ./..."

# Generate models from the JSON definitions (spec/json/*.json) into $(MODELS_DIR)/.
spec:
	@test -d $(SPEC_JSON_DIR) || (echo "Missing $(SPEC_JSON_DIR) — run 'git submodule update --init'"; exit 1)
	cd generators/json && $(GO) run . \
		-json-dir ../../$(SPEC_JSON_DIR)/ \
		-out-dir  ../../$(MODELS_DIR)
	$(GO) build ./$(MODELS_DIR)/...

# Regenerate models from the legacy XML SMDX definitions (spec/smdx/*.xml).
spec-xml:
	@test -d $(SPEC_XML_DIR) || (echo "Missing $(SPEC_XML_DIR) — run 'git submodule update --init'"; exit 1)
	cd generators/xml && $(GO) run . \
		-smdx-dir ../../$(SPEC_XML_DIR)/
	$(GO) build ./$(MODELS_DIR)/...

build:
	$(GO) build ./...

test:
	$(GO) test ./...
