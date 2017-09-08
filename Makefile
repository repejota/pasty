VERSION=`cat VERSION`
BUILD=`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:%h -1`
PACKAGES = $(shell go list ./...)

# Setup the -ldflags option for go build here, interpolate the variable
# values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

# Build & Install

install:
	go install $(LDFLAGS) -v $(PACKAGES)

build:
	go build $(LDFLAGS) -v $(PACKAGES)

.PHONY: version
version:
	@echo $(VERSION)-$(BUILD)

# Testing

.PHONY: test
test:
	go test -v $(PACKAGES)

.PHONY: cover
cover:
	go test -cover $(PACKAGES)

.PHONY: cover-html
cover-html:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	rm -rf coverage.out
	go tool cover -html=coverage-all.out

# Lint

lint:
	gometalinter --tests ./...

# Dependencies

deps:
	go get -u "github.com/getlantern/systray"

dev-deps:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

# Pacakge and Distribution

dist: dist-linux dist-darwin dist-windows

dist-linux:
	GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o pasty-${VERSION}-linux-amd64
	zip pasty-${VERSION}-linux-amd64.zip pasty-${VERSION}-linux-amd64 README.md LICENSE
	GOOS=linux GOARCH=386 go build ${LDFLAGS} -o pasty-${VERSION}-linux-386
	zip pasty-${VERSION}-linux-386.zip pasty-${VERSION}-linux-386 README.md LICENSE

dist-darwin:
	GOOS=darwin GOARCH=amd64 go build ${LDFLAGS} -o pasty-${VERSION}-darwin-amd64
	zip pasty-${VERSION}-darwin-amd64.zip pasty-${VERSION}-darwin-amd64 README.md LICENSE
	GOOS=darwin GOARCH=386 go build ${LDFLAGS} -o pasty-${VERSION}-darwin-386
	zip pasty-${VERSION}-darwin-386.zip pasty-${VERSION}-darwin-386 README.md LICENSE

dist-windows:
	GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o pasty-${VERSION}-windows-amd64.exe
	zip pasty-${VERSION}-windows-amd64.zip pasty-${VERSION}-windows-amd64.exe README.md LICENSE
	GOOS=windows GOARCH=386 go build ${LDFLAGS} -o pasty-${VERSION}-windows-386.exe
	zip pasty-${VERSION}-windows-386.zip pasty-${VERSION}-windows-386.exe README.md LICENSE

# Cleaning up

.PHONY: clean
clean:
	go clean
	rm -rf coverage-all.out
	rn -rf pasty-*
