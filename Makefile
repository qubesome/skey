TARGET_BIN ?= build/skey

.PHONY: build
build:
	go build -o $(TARGET_BIN) cmd/skey/main.go

test:
	go test ./...
