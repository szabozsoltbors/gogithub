# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

SOURCE_PATH=./src/api
BINARY_PATH=./build
BINARY_NAME=app

.PHONY: all ./build clean


build:
	$(GOBUILD) -o $(BINARY_PATH)/$(BINARY_NAME) -i $(SOURCE_PATH)
	cp -r $(SOURCE_PATH)/templates $(BINARY_PATH)/templates

run:
	cd $(BINARY_PATH); ./$(BINARY_NAME)

clean:
	$(GOCLEAN)
	rm -r $(BINARY_PATH)
	