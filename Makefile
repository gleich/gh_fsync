##########
# Building
##########

build-docker-prod:
	docker build -t mattgleich/gh_fsync:latest .
build-docker-dev:
	docker build -f dev.Dockerfile -t mattgleich/gh_fsync:test .
build-docker-dev-lint:
	docker build -f dev.lint.Dockerfile -t mattgleich/gh_fsync:lint .
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
lint-goreleaser:
	goreleaser check
lint-hadolint:
	hadolint Dockerfile
	hadolint dev.Dockerfile
	hadolint dev.lint.Dockerfile
lint-in-docker: build-docker-dev-lint
	docker run mattgleich/gh_fsync:lint

#########
# Testing
#########

test-go:
	go get -v -t -d ./...
	go test ./...
test-in-docker: build-docker-dev
	docker run mattgleich/gh_fsync:test

##########
# Grouping
##########

# Testing
local-test: test-go
docker-test: test-in-docker
# Linting
local-lint: lint-golangci lint-goreleaser lint-hadolint lint-gomod
docker-lint: lint-in-docker
# Build
local-build: build-docker-prod build-docker-dev build-docker-dev-lint
