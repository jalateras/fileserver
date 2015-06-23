.PHONY: all clean echo test fmt install bench run

EXECUTABLE = fileserver
GDFLAGS ?= $(GDFLAGS:)
ARGS ?= $(ARGS:)

all: test

clean:
	@echo "===> Cleaning"
	@godep go clean $(GDFLAGS) -i ./...

build:
	@echo "===> Building"
	@godep go build $(GDFLAGS) -o $(EXECUTABLE) ./...

fmt:
	@echo "===> Formatting"
	@godep go fmt $(GDFLAGS) ./...

install:
	@echo "===> Installing"
	@godep go install $(GDFLAGS)

test:
	@echo "===> Testing"
	@godep go test $(GDFLAGS) ./...

bench:
	@echo "===> Benchmarking"
	@godep go test -run=NONE -bench=. $(GDFLAGS) ./...

start: build
	@echo "===> Starting Server"
	@./$(EXECUTABLE) $(ARGS)

