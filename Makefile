GO_BIN := $(GOPATH)/bin
GOMETALINTER := $(GO_BIN)/gometalinter

# Build builds the api and place them in the projects level bin directory
.PHONY: build
build: clean
	bash ./scripts/build.sh

# Test runs tests on the
.PHONY: test
test: lint
	bash ./scripts/test.sh

.PHONY: lint
lint:
	# Check is gometalinter is installed
	if [ ! -f $GOMETALINTER ]; then
  	# If gometalinter is not installed then install it
		go get -u github.com/alecthomas/gometalinter
		gometalinter --install &> /dev/null
	fi

	bash ./scripts/lint.sh

.PHONY: clean
clean:
	bash ./scripts/clean.sh
