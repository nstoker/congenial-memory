OUT := congenial-memory
IMG := gcr.io/pixelate-199308/${OUT}
PKG := github.com/nstoker/congenial-memory
VERSION=$(shell git describe --always --long --dirty)
TAG=$(shell git describe --always | cut -d- -f1)
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/)

all: run

docker: docker-check server
	@echo "building version $(DOCKER_NAME):$(TAG)"
	@docker build -t ${IMG}:${TAG} -t${IMG}:latest --build-arg VERSION=${VERSION} --build-arg GITLAB_TOKEN=${GITLAB_TOKEN} --build-arg GITLAB_NAME=${GITLAB_NAME} .

push:
	@echo Version : $(VERSION).
	@echo "DockerV : $(TAG)"
	@echo ${IMG} $(TAG).
	docker push ${IMG}:${TAG}
	docker push $(IMG):latest

docker-check:
ifndef GITLAB_NAME
	echo $(error The GITLAB_NAME environment variable has not been set)
endif
ifndef GITLAB_TOKEN
	echo $(error The GITLAB_TOKEN environment variable has not been set)
endif

server:
	go build -i -v -o ${OUT} -ldflags="-X main.version=${VERSION}" ${PKG}

test:
	@go test -short ${PKG_LIST}

vet:
	@go vet ${PKG_LIST}

lint:
	@for file in ${GO_FILES} ;  do \
		golint $$file ; \
	done

static: vet lint
	go build -i -v -o ${OUT}-v${VERSION} -tags netgo -ldflags="-extldflags \"-static\" -w -s -X main.version=${VERSION}" ${PKG}

run: server
	./${OUT}

clean:
	-@rm ${OUT} ${OUT}-v*

.PHONY: run server static vet lint
