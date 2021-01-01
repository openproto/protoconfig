REPO_ROOT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

# Tools.
PROTO_DESCRIPTOR_FILE ?= $(REPO_ROOT_DIR)/proto/google/protobuf/descriptor.proto

# buf-v0.33.0 protoc --version prints 3.13.0-buf
PROTOC_VERSION        ?= "3.13.0"

GIT ?= $(shell which git)
TMP_PATH ?= /tmp

# Support gsed on OSX (installed via brew), falling back to sed. On Linux
# systems gsed won't be installed, so will use sed as expected.
SED ?= $(shell which gsed 2>/dev/null || which sed)

define require_clean_work_tree
	@git update-index -q --ignore-submodules --refresh

    @if ! git diff-files --quiet --ignore-submodules --; then \
        echo >&2 "cannot $1: you have unstaged changes."; \
        git diff-files --name-status -r --ignore-submodules -- >&2; \
        echo >&2 "Please commit or stash them."; \
        exit 1; \
    fi

    @if ! git diff-index --cached --quiet HEAD --ignore-submodules --; then \
        echo >&2 "cannot $1: your index contains uncommitted changes."; \
        git diff-index --cached --name-status -r --ignore-submodules HEAD -- >&2; \
        echo >&2 "Please commit or stash them."; \
        exit 1; \
    fi

endef

help: ## Displays help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

$(PROTO_DESCRIPTOR_FILE): $(REPO_ROOT_DIR)/common.mk
	@echo ">> fetching descriptor.proto that matches protoc version $(PROTOC_VERSION) as $(PROTO_DESCRIPTOR_FILE)"
	@wget --output-document="$(PROTO_DESCRIPTOR_FILE)" "https://raw.githubusercontent.com/protocolbuffers/protobuf/v$(PROTOC_VERSION)/src/google/protobuf/descriptor.proto"
