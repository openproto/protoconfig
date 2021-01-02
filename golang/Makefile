include ../common.mk
include ../.bingo/Variables.mk

MODULE_EXAMPLES   := examples
MODULE_GEN_GO     := protoc-gen-go-openconfig
MODULE_KINGPINV2  := kingpinv2

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
	@cd $(MODULE_EXAMPLES) && go mod tidy && go mod verify
	@cd $(MODULE_GEN_GO) && go mod tidy && go mod verify
	@cd $(MODULE_KINGPINV2) && go mod tidy && go mod verify

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
	@cd $(MODULE_EXAMPLES) && go test -v -timeout=30m ./...
	@cd $(MODULE_GEN_GO) && go test -v -timeout=30m ./...
	@cd $(MODULE_KINGPINV2) && go test -v -timeout=30m ./...

# For protoc naming matters.
PROTOC_GEN_GO_CURRENT := $(TMP_GOBIN)/protoc-gen-go

.PHONY: proto
proto: ## Generated golang protobuf
proto: build $(BUF) $(PROTOC_GEN_GO)
	@mkdir -p $(TMP_GOBIN)
	@cp $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_CURRENT)
	@echo ">> generating $(REPO_ROOT_DIR)/proto/openconfig/v1/extensions.proto in $(REPO_ROOT_DIR)/golang"
	@PATH=$(GOBIN):$(TMP_GOBIN) $(BUF) protoc \
		-I $(REPO_ROOT_DIR)/proto \
		--go_out=$(REPO_ROOT_DIR)/golang --go_opt=module="github.com/thanos-io/OpenConfig/golang" \
	    $(REPO_ROOT_DIR)/proto/openconfig/v1/extensions.proto
	@echo ">> generating $(REPO_ROOT_DIR)/proto/openconfig/golang/kingpinv2/v1/extensions.proto in $(REPO_ROOT_DIR)/golang/kingpinv2"
	@PATH=$(GOBIN):$(TMP_GOBIN) $(BUF) protoc \
		-I $(REPO_ROOT_DIR)/proto \
		--go_out=$(REPO_ROOT_DIR)/golang/kingpinv2 --go_opt=module="github.com/thanos-io/OpenConfig/golang/kingpinv2" \
		$(REPO_ROOT_DIR)/proto/openconfig/golang/kingpinv2/v1/extensions.proto
	@echo ">> generating $(REPO_ROOT_DIR)/proto/examples/helloworld/v1/helloworld.proto in $(REPO_ROOT_DIR)/golang/examples/helloworld/"
	@PATH=$(GOBIN):$(TMP_GOBIN) $(BUF) protoc \
		-I $(REPO_ROOT_DIR)/proto \
		--go_out=$(REPO_ROOT_DIR)/golang/examples/helloworld/ --go_opt=module="github.com/thanos-io/OpenConfig/golang/examples/helloworld" \
	   	--go-openconfig_out=$(REPO_ROOT_DIR)/golang/examples/helloworld/ --go-openconfig_opt=module="github.com/thanos-io/OpenConfig/golang/examples/helloworld" \
	    $(REPO_ROOT_DIR)/proto/examples/helloworld/v1/helloworld.proto

.PHONY: check-git
check-git:
ifneq ($(GIT),)
	@test -x $(GIT) || (echo >&2 "No git executable binary found at $(GIT)."; exit 1)
else
	@echo >&2 "No git binary found."; exit 1
endif

.PHONY: lint
lint: ## Runs various static analysis against our code.
lint: $(FAILLINT) $(GOLANGCI_LINT) $(MISSPELL) build format proto check-git deps
	@DIR="." $(MAKE) lint_module
	@DIR=$(MODULE_EXAMPLES) $(MAKE) lint_module
	@DIR=$(MODULE_GEN_GO) $(MAKE) lint_module
	@DIR=$(MODULE_KINGPINV2) $(MAKE) lint_module
	@echo ">> detecting misspells"
	@find . -type f | grep -v vendor/ | grep -vE '\./\..*' | xargs $(MISSPELL) -error
	$(call require_clean_work_tree,"detected files change during formar/lint process - run make lint and commit changes.")

.PHONY: lint_module
# PROTIP:
# Add
#      --cpu-profile-path string   Path to CPU profile output file
#      --mem-profile-path string   Path to memory profile output file
# to debug big allocations during linting.
lint_module: ## Runs various static analysis against our code.
lint_module:
	@echo ">> verifying modules being imported"
	@cd $(DIR) && $(FAILLINT) -paths "errors=github.com/pkg/errors" ./...
	@cd $(DIR) && $(FAILLINT) -paths "fmt.{Print,Printf,Println}" -ignore-tests ./...
	@echo ">> examining all of the Go files"
	@cd $(DIR) && go vet -stdmethods=false ./...
	@echo ">> linting all of the Go files GOGC=${GOGC}"
	@cd $(DIR) && $(GOLANGCI_LINT) run
