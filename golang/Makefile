include ../common.mk
include ../.bingo/Variables.mk

MODULE_HELLO_WORLD := examples/helloworld
MODULE_GEN_GO     :=protoc-gen-go-openconfig

FILES_TO_FMT      ?= $(shell find . -path ./vendor -prune -o -name '*.go' -print)
TMP_GOBIN ?= $(TMP_PATH)/gobin

GO111MODULE       ?= on
export GO111MODULE

GOBIN ?= $(firstword $(subst :, ,${GOPATH}))/bin

.PHONY: all
all: format build

.PHONY: build
build: ## Build protoc-gen-go-openconfig.
	@echo ">> building $(MODULE_GEN_GO)"
	@cd $(MODULE_GEN_GO) && GOBIN=$(GOBIN) go install github.com/thanos-io/OpenConfig/golang/protoc-gen-go-openconfig

.PHONY: deps
deps: ## Ensures fresh go.mod and go.sum.
	@go mod tidy && go mod verify
	@cd $(MODULE_HELLO_WORLD) && go mod tidy && go mod verify
	@cd $(MODULE_GEN_GO) && go mod tidy && go mod verify

.PHONY: format
format: ## Formats Go code.
format: $(GOIMPORTS)
	@echo ">> formatting code"
	@$(GOIMPORTS) -w $(FILES_TO_FMT)

.PHONY: test
test: ## Runs all Go unit tests.
export GOCACHE=/tmp/cache
test:
	@echo ">> running unit tests"
	@go mod tidy && go test -v -timeout=30m ./...
	@cd $(MODULE_HELLO_WORLD) && go test -v -timeout=30m ./...
	@cd $(MODULE_GEN_GO) && go test -v -timeout=30m ./...

# For protoc naming matters.
PROTOC_GEN_GO_CURRENT := $(TMP_GOBIN)/protoc-gen-go

.PHONY: proto
proto: ## Generated golang protobuf
proto: build $(PROTOC) $(PROTOC_GEN_GO)
	@mkdir -p $(TMP_GOBIN)
	@cp $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_CURRENT)
	@echo ">> generating $(REPO_ROOT_DIR)/examples/helloworld/helloworld.proto in $(REPO_ROOT_DIR)/golang/examples/helloworld/"
	@PATH=$(GOBIN):$(TMP_GOBIN) $(PROTOC) \
		-I $(REPO_ROOT_DIR)/examples/helloworld \
		-I $(REPO_ROOT_DIR) \
		--go_out=./examples/helloworld/ --go_opt=paths=source_relative \
	   	--go-openconfig_out=./examples/helloworld/ --go-openconfig_opt=paths=source_relative \
	    $(REPO_ROOT_DIR)/examples/helloworld/helloworld.proto

.PHONY: check-git
check-git:
ifneq ($(GIT),)
	@test -x $(GIT) || (echo >&2 "No git executable binary found at $(GIT)."; exit 1)
else
	@echo >&2 "No git binary found."; exit 1
endif

.PHONY: lint
lint: ## Runs various static analysis against our code.
lint:
	@DIR="." $(MAKE) lint_module
	@DIR=$(MODULE_HELLO_WORLD) $(MAKE) lint_module
	@DIR=$(MODULE_GEN_GO) $(MAKE) lint_module
	@echo ">> detecting misspells"
	@find . -type f | grep -v vendor/ | grep -vE '\./\..*' | xargs $(MISSPELL) -error
	$(call require_clean_work_tree,"detected files without copyright - run make lint and commit changes.")

.PHONY: lint_module
# PROTIP:
# Add
#      --cpu-profile-path string   Path to CPU profile output file
#      --mem-profile-path string   Path to memory profile output file
# to debug big allocations during linting.
lint_module: ## Runs various static analysis against our code.
lint_module: $(FAILLINT) $(GOLANGCI_LINT) $(MISSPELL) build format docs check-git deps
	$(call require_clean_work_tree,"detected not clean master before running lint - run make lint and commit changes.")
	@echo ">> verifying modules being imported"
	@$(FAILLINT) -paths "errors=github.com/pkg/errors" $(DIR)/...
	@$(FAILLINT) -paths "fmt.{Print,Printf,Println}" -ignore-tests $(DIR)/...
	@echo ">> examining all of the Go files"
	@go vet -stdmethods=false $(DIR)/...
	@echo ">> linting all of the Go files GOGC=${GOGC}"
	@cd $(DIR) && $(GOLANGCI_LINT) run
