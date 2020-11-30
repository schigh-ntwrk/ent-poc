.DELETE_ON_ERROR:
.DEFAULT_GOAL := help
_YELLOW=\033[0;33m
_NC=\033[0m

MODULE = github.com/schigh-ntwrk/ent-poc

TOOLS_DIR         = $(CURDIR)/_tools
BIN_DIR           = $(CURDIR)/_bin
COVER_DIR         = $(CURDIR)/_coverage
BUILD_DIR         = $(CURDIR)/_build
BUILD_ROOT        = $(CURDIR)/cmd/app
SCHEMA_ROOT       = $(CURDIR)/internal/ent/schema
COVER_OUT         = $(wildcard $(COVER_DIR)/*.out)
COVER_INTERCHANGE = $(COVER_OUT:.out=.interchange)
COVER_HTML        = $(COVER_OUT:.out=.html)
COVER_XML         = $(COVER_OUT:.out=.xml)
COVER_COMBINED    = $(COVER_DIR)/combined.out

# these are _tools for building/testing/codegen, etc
TOOLS := github.com/golang/mock/mockgen
TOOLS := $(TOOLS) golang.org/x/tools/cmd/goimports
TOOLS := $(TOOLS) github.com/golangci/golangci-lint/cmd/golangci-lint
TOOLS := $(TOOLS) github.com/axw/gocov/gocov
TOOLS := $(TOOLS) github.com/matm/gocov-html
TOOLS := $(TOOLS) github.com/AlekSi/gocov-xml
TOOLS := $(TOOLS) github.com/wadey/gocovmerge
TOOLS := $(TOOLS) mvdan.cc/sh/v3/cmd/shfmt
TOOLS := $(TOOLS) github.com/google/wire/cmd/wire
TOOLS := $(TOOLS) github.com/facebook/ent/cmd/entc

#	###########################################################################
#	NON-PATH TARGETS
#	###########################################################################

.PHONY: help
help: ## prints this help
	@ grep -hE '^[\.a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "${_YELLOW}%-16s${_NC} %s\n", $$1, $$2}'

.PHONY: tools
tools: $(BIN_DIR) ## install tools
	@ $(MAKE) $(BIN_DIR)

.PHONY: fmt
fmt: $(BIN_DIR) ## apply formatting and fix imports
	@ GO111MODULE=on \
	  GOFLAGS="$(GO_FLAGS)" \
	  $(BIN_DIR)/goimports -w \
	  	-local $(MODULE) \
		$(shell find . -type f -name '*.go' -not -path "./vendor/*")

.PHONY: lint
lint: $(BIN_DIR) ## run linter
	@ GO111MODULE=on \
	  GOFLAGS="$(GO_FLAGS)" \
	  $(BIN_DIR)/golangci-lint run \
		--config $(CURDIR)/.golangci.yaml \
		--print-resources-usage \
		--verbose

.PHONY: build
build: $(BUILD_DIR)
	@ go build -a -o $(BUILD_DIR)/app $(BUILD_ROOT)

.PHONY: wire
wire: $(BIN_DIR) ## run wire dependency injector
	@{ \
  		cd $(CURDIR)/internal/bootstrap ;\
  		$(BIN_DIR)/wire ;\
	}

.PHONY: ent
ent: $(BIN_DIR) ## run ent schema generator
	@ $(BIN_DIR)/entc generate $(SCHEMA_ROOT)

.PHONY: ent-describe
ent-describe:
	@ $(BIN_DIR)/entc describe $(SCHEMA_ROOT)

.PHONY: gen
gen: wire ent ## run all the generators

#	###########################################################################
#	PATH TARGETS
#	###########################################################################

$(BIN_DIR): $(TOOLS_DIR)
	@ cd $(TOOLS_DIR) && GOBIN=$(BIN_DIR) GOFLAGS="-mod=" go install $(TOOLS)

$(COVER_DIR):
	@ mkdir -p $(CURDIR)/_coverage

$(BUILD_DIR):
	@ mkdir -p $(CURDIR)/_build
