GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_CLEAN=$(GO_CMD) clean
GO_DEPS=$(GO_CMD) get -d -v
GO_DEPS_UPDATE=$(GO_CMD) get -d -v -u
GO_VET=$(GO_CMD) vet
GO_FMT=$(GO_CMD) fmt
GO_LINT=golint

TOP_PACKAGE_DIR := github.com/jalateras
PACKAGE_LIST := fileserver

.PHONY: all build

all: build

build: vet
	@for p in $(PACKAGE_LIST); do \
		echo "==> Build $$p ..."; \
		$(GO_BUILD) $(TOP_PACKAGE_DIR)/$$p || exit 1; \
	done

fmt:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Formatting $$p ..."; \
		$(GO_FMT) $(TOP_PACKAGE_DIR)/$$p || exit 1; \
	done

vet:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Vet $$p ..."; \
		$(GO_VET) $(TOP_PACKAGE_DIR)/$$p; \
	done

lint:
	@for p in $(PACKAGE_LIST); do \
		echo "==> Lint $$p ..."; \
		$(GO_LINT) src/$(TOP_PACKAGE_DIR)/$$p/main.go; \
	done

