.DEFAULT_GOAL := help

## Setup
setup:
	go get -v -u github.com/golang/dep/cmd/dep

## Install dependencies
deps:
	dep ensure

## Run
run:
	go run *.go

## Run Debug mode
debug_run:
	DEBUG=true BASIC_AUTH=false go run *.go

## Build docker dev image
build_docker_image:
	docker image build -t youyo/rotation-shifts:dev .

## Start docker-compose
docker_compose:
	docker-compose up

## GenerateSchema
generate_schema:
	mysqldump --defaults-file=my.cnf.mysqldump --column-statistics=0 --no-data rotation_shifts > db/schema.sql
	mysqldump --defaults-file=my.cnf.mysqldump --column-statistics=0 rotation_shifts > db/schema-include-sample-data.sql

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: help
.SILENT:
