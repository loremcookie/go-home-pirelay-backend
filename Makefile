# Build builds the api and place them in the projects level bin directory
.PHONY: build
build:
	bash ./scripts/build.sh

# Test runs tests on the
.PHONY: test
test: lint
	bash ./scripts/test.sh

.PHONY: lint
lint:
	bash ./scripts/lint.sh

.PHONY: clean
clean:
	bash ./scripts/clean.sh