# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

SOURCE_PATH=./src
BINARY_PATH=./build
BINARY_NAME=app

.PHONY: all ./build clean

all: clean build run

build:
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) -i $(SOURCE_PATH)
	cp -r $(SOURCE_PATH)/web/templates $(BINARY_PATH)/templates

run:
	cd $(BINARY_PATH); ./$(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -r $(BINARY_PATH)
