include common.mk
include .bingo/Variables.mk

.PHONY: all
all: docs

.PHONY: docs
docs: $(MDOX) ## Generates config snippets and doc formatting.
	@echo ">> generating & formatting docs"
	#@$(MDOX) fmt -l $(shell find . -name "*.md" -type f | xargs)

.PHONY: lint
lint: $(BUF) docs
	@echo ">> lint proto files"
	@$(BUF) check lint
	$(call require_clean_work_tree,"detected changed files - run make lint and commit changes.")
	@echo ">> lint golang files"
	@$(MAKE) -C golang lint
	$(call require_clean_work_tree,"detected changed files - run make lint and commit changes.")

