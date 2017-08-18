#! /usr/bin/make
ifeq ($(OS),Windows_NT)
	BUILD_TARGET_FILES = beerbox.exe main.go
else
	BUILD_TARGET_FILES ?= beerbox main.go
endif

.DEFAULT_GOAL := prepare

all: cleandep depend clean mkdir precompile gen precompile build build-console

prepare: cleandep depend clean mkdir precompile gen precompile

depend:
	@glide install

cleandep:
	@rm -rf vendor

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf swagger
	@rm -rf assets

mkdir:
	@mkdir swagger

confinit:
	@cp config/dev.toml.example config/dev.toml 
	@cp config/local.toml.example config/local.toml 
	@cp config/test.toml.example config/test.toml 

gen:
	@goagen app -d goa-study/design
	@goagen swagger -d goa-study/design
	@goagen controller -d goa-study/design --pkg controllers -o controllers
	@touch app/.gitkeep swagger/.gitkeep

precompile:
	@go-bindata -pkg=swaggerassets -o=swaggerassets/swagger.go swagger/... swaggerui/...
	@go-bindata -pkg=assets -o=assets/bindata.go config/...

build:
	@go build -o $(BUILD_TARGET_FILES)

build-console:
	@go build -o $(BUILD_TARGET_FILES_CONSOLE)

go-run:
	@go run main.go

run: precompile go-run

test:
	@go test -tags=test goa-study/models/...  goa-study/controllers/... goa-study/concepts/... -cover