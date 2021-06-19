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

gitpathchanges: cmd/git-path-changes/* pkg/*
	go build -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} ${BUILDLDFLAGS}" ${BUILDARGS} \
		-o ${BUILDOUTPREFIX}git-path-changes cmd/git-path-changes/main.go

clean:
	rm ${BUILDOUTPREFIX}git-path-changes* 2> /dev/null || exit 0

build: gitpathchanges

install: build
	cp ${BUILDOUTPREFIX}git-path-changes* /usr/local/bin
