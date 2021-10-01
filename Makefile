#!/usr/bin/make -f

SHELL:=$(shell which bash)

.PHONY: go-generate swag-generate generate-wire generate download test build image

install-tools:
	@go get -u github.com/google/wire/cmd/wire
	@go get -u github.com/swaggo/swag/cmd/swag

go-generate:
	@go generate

swag-generate:
	@swag init

generate-wire:
	@find cmd/ pkg/ -type d | while read dir ; do \
		pushd "$${dir}" ; \
		wire gen ; \
		popd ; \
	done

generate: go-generate swag-generate

download:
	@go mod download -x

test: download
	@go test ./... || \
		(echo 'ERROR: tests failed!' >&2 && exit 1) && \
		echo 'SUCCESS: all tests passed!'

build: generate download

image:
	@docker build -t naftulikay/golang-webapp:latest ./