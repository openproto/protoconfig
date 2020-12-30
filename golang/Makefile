include ../common.mk
include ../.bingo/Variables.mk

FILES_TO_FMT      ?= $(shell find . -path ./vendor -prune -o -name '*.go' -print)
TMP_GOBIN ?= $(TMP_PATH)/gobin

GO111MODULE       ?= on
export GO111MODULE

GOBIN ?= $(firstword $(subst :, ,${GOPATH}))/bin

.PHONY: all
all: format build

.PHONY: build
build: ## Build protoc-gen-go-openconfig.
	@echo ">> building protoc-gen-go-openconfig"
	@cd protoc-gen-go-openconfig && GOBIN=$(GOBIN) go install github.com/thanos-io/OpenConfig/go/protoc-gen-go-openconfig

.PHONY: deps
deps: ## Ensures fresh go.mod and go.sum.
	@cd protoc-gen-go-openconfig && go mod tidy && go mod verify

.PHONY: format
format: ## Formats Go code.
format: $(GOIMPORTS)
	@echo ">> formatting code"
	@$(GOIMPORTS) -w $(FILES_TO_FMT)

.PHONY: test
test: ## Runs all Go unit tests.
export GOCACHE=/tmp/cache
test:
	@echo ">> running unit tests (without cache)"
	@rm -rf $(GOCACHE)
	@go test -v -timeout=30m $(shell go list ./... | grep -v /vendor/);

# For protoc naming matters.
PROTOC_GEN_GO_CURRENT := $(TMP_GOBIN)/protoc-gen-go

.PHONY: proto
proto: ## Generated golang protobuf
proto: build $(PROTOC) $(PROTOC_GEN_GO)
	@mkdir -p $(TMP_GOBIN)
	@cp $(PROTOC_GEN_GO) $(PROTOC_GEN_GO_CURRENT)
	@echo ">> generating $(REPO_ROOT_DIR)/examples/helloworld/helloworld.proto in $(REPO_ROOT_DIR)/go/examples/helloworld/"
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

# PROTIP:
# Add
#      --cpu-profile-path string   Path to CPU profile output file
#      --mem-profile-path string   Path to memory profile output file
# to debug big allocations during linting.
lint: ## Runs various static analysis against our code.
lint: $(FAILLINT) $(GOLANGCI_LINT) $(MISSPELL) build format docs check-git deps
	$(call require_clean_work_tree,"detected not clean master before running lint")
	@echo ">> verifying modules being imported"
	@$(FAILLINT) -paths "errors=github.com/pkg/errors" ./...
	@$(FAILLINT) -paths "fmt.{Print,Printf,Println}" -ignore-tests ./...
	@echo ">> examining all of the Go files"
	@go vet -stdmethods=false ./...
	@echo ">> linting all of the Go files GOGC=${GOGC}"
	@$(GOLANGCI_LINT) run
	@echo ">> detecting misspells"
	@find . -type f | grep -v vendor/ | grep -vE '\./\..*' | xargs $(MISSPELL) -error
	@echo ">> ensuring Copyright headers"
	@go run ./scripts/copyright/...
	$(call require_clean_work_tree,"detected files without copyright; run make lint file and commit changes.")
