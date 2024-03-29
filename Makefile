DOCKER_PORT ?= 80
HOST_PORT ?= 6060
DOCKER_IMAGE ?= golang:1.12

DOCKER_VOLUME_PATH = /go/src/go-alg
BASE_COMMAND = docker run -it -v $(PWD):$(DOCKER_VOLUME_PATH) --rm -p $(HOST_PORT):$(DOCKER_PORT) $(DOCKER_IMAGE)

.PHONY: bash
bash:
	$(BASE_COMMAND) bash

.PHONY: go-version
go-version:
	$(BASE_COMMAND) go version

.PHONY: go-test
go-test:
	$(BASE_COMMAND) bash -c "cd $(DOCKER_VOLUME_PATH) && go test ./... --race"

.PHONY: go-list
go-list:
	$(BASE_COMMAND) bash -c "cd $(DOCKER_VOLUME_PATH) && go list ./..."

.PHONY: go-doc
go-doc:
	$(BASE_COMMAND) godoc -http=0.0.0.0:80