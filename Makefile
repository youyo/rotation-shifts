.DEFAULT_GOAL := help

## Run
run:
	go run *.go

## Run Debug mode
debug_run:
	DEBUG=true go run *.go

## GenerateSchema
generate_schema:
	mysqldump --defaults-file=my.cnf.mysqldump --column-statistics=0 --no-data rotation_shifts > db/schema.sql
	mysqldump --defaults-file=my.cnf.mysqldump --column-statistics=0 rotation_shifts > db/schema-include-sample-data.sql

## Show help
help:
	@make2help $(MAKEFILE_LIST)

.PHONY: help
.SILENT:
