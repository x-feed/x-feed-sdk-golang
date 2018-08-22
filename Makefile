SHELL := /bin/sh

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(patsubst %/,%,$(dir $(MAKEFILE_PATH)))


define LINT
	@echo "Running code linters..."
	golangci-lint run -j 4 -v
endef

define GEN
	@echo "Generating go stubs from protobuf files"
	${CURRENT_DIR}/scripts/genproto.sh
endef


.PHONY: default
default: lint


.PHONY: lint
lint:
	@$(call LINT)

.PHONY: gen
gen:
	@$(call GEN)