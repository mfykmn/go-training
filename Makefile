## Go
.PHONY: setup
setup:
	@echo "Start setup"
	@env GO111MODULE=on go mod vendor
	@go generate $(shell go list ./... | grep -v /vendor/)
