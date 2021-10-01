#!/usr/bin/make -f

.PHONY: download install

download:
	@go mod download

get:
	@go get -u github.com/naftulikay/golang-snakes/cmd/x/modname

build: get download
	@go build

install: build
	@go install