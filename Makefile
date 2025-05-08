LOCAL_BIN:=$(CURDIR)/bin

run:
	go run ./cmd/cart

PHONY: .bin-mock
.bin-mock:
	$(info Installing mockery...)
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v3@v3.2.5

PHONY: mocks
mocks: .bin-mock
	$(info Generate mocks...)
	$(LOCAL_BIN)/mockery

PHONY: test
test:
	$(info Run tests...)
	go test -v -race ./...