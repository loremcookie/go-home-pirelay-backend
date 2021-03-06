# Build builds the api and place them in the projects level bin directory
.PHONY: build
build: clean
	bash ./scripts/build.sh

# Test runs tests on the
.PHONY: test
test:
	bash ./scripts/test.sh

.PHONY: lint
lint:
	bash ./scripts/lint.sh

.PHONY: clean
clean:
	bash ./scripts/clean.sh
