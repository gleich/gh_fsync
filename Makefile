##########
# Building
##########

build-docker-prod:
	docker build -t mattgleich/gh_fsync:latest .
build-docker-dev:
	docker build -f docker/dev.Dockerfile -t mattgleich/gh_fsync:test .
build-docker-dev-lint:
	docker build -f docker/dev.lint.Dockerfile -t mattgleich/gh_fsync:lint .
build-docker-platform:
	docker build -f docker/platform.Dockerfile -t mattgleich/gh_fsync:platform .
build-go:
	go get -v -t -d ./...
	go build -v .
	rm gh_fsync

#########
# Linting
#########

lint-golangci:
	golangci-lint run
lint-gomod:
	go mod tidy
	git diff --exit-code go.mod
	git diff --exit-code go.sum
lint-hadolint:
	hadolint Dockerfile
	hadolint docker/dev.Dockerfile
	hadolint docker/dev.lint.Dockerfile
lint-in-docker: build-docker-dev-lint
	docker run mattgleich/gh_fsync:lint

#########
# Testing
#########

test-go:
	go get -v -t -d ./...
	go test -v ./...
test-in-docker: build-docker-dev
	docker run mattgleich/gh_fsync:test

###########
# Releasing
###########

release: build-docker-platform
	docker push mattgleich/gh_fsync:platform

##########
# Grouping
##########

# Testing
local-test: test-go
docker-test: test-in-docker
# Linting
local-lint: lint-golangci lint-hadolint lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-prod build-docker-dev build-docker-dev-lint build-docker-platform
