.PHONY: clean build install

COMMIT?=${BUILDCOMMIT}
VERSION?=${BUILDTAG}

# enable cgo because it's required by OSX keychain library
CGO_ENABLED=1

# enable go modules
GO111MODULE=on

export CGO_ENABLED GO111MODULE

dep:
	go get ./...

test:
	go test ./...

lint:
	golangci-lint run

gitpathchanges: cmd/gitpathchanges/* pkg/*
	go build -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} ${BUILDLDFLAGS}" ${BUILDARGS} \
		-o ${BUILDOUTPREFIX}gitpathchanges cmd/gitpathchanges/main.go

clean:
	rm ${BUILDOUTPREFIX}gitpathchanges* 2> /dev/null || exit 0

build: gitpathchanges

install: build
	cp ${BUILDOUTPREFIX}gitpathchanges* /usr/local/bin
