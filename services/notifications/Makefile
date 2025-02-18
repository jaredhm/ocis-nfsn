SHELL := bash
NAME := notifications

# Where to write the files generated by this makefile.
OUTPUT_DIR = ./pkg/email/l10n
TEMPLATE_FILE = ./pkg/email/l10n/notifications.pot

include ../../.make/recursion.mk

############ tooling ############
ifneq (, $(shell command -v go 2> /dev/null)) # suppress `command not found warnings` for non go targets in CI
include ../../.bingo/Variables.mk
endif

############ go tooling ############
include ../../.make/go.mk

############ release ############
include ../../.make/release.mk

############ docs generate ############
include ../../.make/docs.mk

.PHONY: docs-generate
docs-generate: config-docs-generate

############ generate ############
include ../../.make/generate.mk

.PHONY: ci-go-generate
ci-go-generate: # CI runs ci-node-generate automatically before this target

.PHONY: ci-node-generate
ci-node-generate:

############ translations ########
.PHONY: l10n-pull
l10n-pull:
	cd $(OUTPUT_DIR) && tx pull --all --force --skip --minimum-perc=75

.PHONY: l10n-push
l10n-push:
	cd $(OUTPUT_DIR) && tx push -s --skip

.PHONY: l10n-read
l10n-read: $(GO_XGETTEXT)
	go-xgettext -o $(OUTPUT_DIR)/notifications.pot --keyword=Template --add-comments -s pkg/email/templates.go

.PHONY: l10n-write
l10n-write:

.PHONY: l10n-clean
l10n-clean:
	rm -f $(TEMPLATE_FILE);

############ licenses ############
.PHONY: ci-node-check-licenses
ci-node-check-licenses:

.PHONY: ci-node-save-licenses
ci-node-save-licenses:
