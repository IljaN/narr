VERSION := $(shell git describe --tags --dirty --always)
LDFLAGS = -ldflags "-X main.Version=${VERSION}"
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' ! -name '*_test.go')

build:
	CGO_ENABLED=0 go build ${LDFLAGS} .

build_all: bin/narr-amd64-linux bin/narr-amd64-darwin bin/narr-amd64-windows.exe bin/narrl-arm64-darwin

bin/narr-amd64-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o bin/narr-amd64-linux $(SOURCES)

bin/narr-amd64-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o bin/narr-amd64-darwin $(SOURCES)

bin/narr-amd64-windows.exe:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o bin/narr-amd64-windows.exe $(SOURCES)

bin/narrl-arm64-darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build ${LDFLAGS} -o bin/narr-arm64-darwin $(SOURCES)

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
coverage:
	go test -coverprofile cover.out -v ./...

.PHONY: show-coverage
show-coverage: coverage
	go tool cover -html=cover.out

.PHONY: clean
clean:
	rm -rf ./bin narr cover.out