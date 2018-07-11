HAVE_VGO:=$(shell which vgo)

## Go
.PHONY: setup
setup: vgo
	@echo "Start setup"
	@vgo mod -vendor
	@go generate $(shell go list ./... | grep -v /vendor/)

## Install package
.PHONY: vgo
vgo:
ifndef HAVE_VGO
	@echo "Installing vgo"
	@go get -u golang.org/x/vgo
endif
