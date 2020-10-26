OUTPUT = main
PKG_LIST := $(shell go list ./... | grep -v /vendor/)

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: install
install:
	go get -u golang.org/x/lint/golint
	go get ./...

main: main.go
	goreleaser build --snapshot --rm-dist

.PHONY: clean
clean:
	rm -rf api/*.go

.PHONY: lint
lint:
	@golint -set_exit_status ${PKG_LIST}

.PHONY: proto
proto:
	protoc --go_out=plugins=grpc:. --proto_path=. api/*.proto
