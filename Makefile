.PHONY: build

build:
	sam build

start:
	sam build
	sam local start-api
