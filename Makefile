# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

SOURCE_PATH=./src
BINARY_PATH=./build
BINARY_NAME=app

profile="prod"

.PHONY: all ./build clean

all: clean build test run

build:
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) -i $(SOURCE_PATH)
	cp -r $(SOURCE_PATH)/web/templates $(BINARY_PATH)/templates
	cp -r $(SOURCE_PATH)/api/config/environments $(BINARY_PATH)/environments

test:
	$(GOTEST) $(SOURCE_PATH)/util -v

run:
	cd $(BINARY_PATH); ./$(BINARY_NAME) -profile=${profile}

clean:
	$(GOCLEAN)
	rm -r $(BINARY_PATH)
	mkdir build
