LOCAL_BIN:=$(CURDIR)/bin

run:
	go run ./cmd/cart

PHONY: .bin-mock
.bin-mock:
	$(info Installing mockery...)
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v3@v3.2.5

PHONY: test
test:
	go test -v -race ./...