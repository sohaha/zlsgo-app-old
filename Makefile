BUILD_NAME=$(shell basename "$(PWD)")
GOBASE=$(shell pwd)
GOBIN=$(GOBASE)/bin

default: build

.PHONY: run
run:
	go build -o tmpApp
	./tmpApp

.PHONY: test
test:
	go test ./... -v

.PHONY: clean
clean:
	go clean

.PHONY: dev
dev:
	zzz w

.PHONY: build
build:test
	zzz build --os mac,win,linux -P

.PHONY: buildSqlite3
buildSqlite3:test
	zzz build --os mac,win,linux --cgo -P -- --tags sqlite3