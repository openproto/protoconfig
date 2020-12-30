include common.mk
include .bingo/Variables.mk

.PHONY: all
all: docs

.PHONY: docs
docs: $(MDOX) ## Generates config snippets and doc formatting.
	@echo ">> generating & formatting docs"
	@$(MDOX) fmt -l *.md **/*.md
