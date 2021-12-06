PROJECT_ROOT            ?= $(PWD)
GO_PACKAGES             ?= $(shell go list ./... | grep -v vendor)
GO_TEST_FLAGS           ?= -v -cover

.PHONY: fmt
fmt:
	@go fmt $(GO_PACKAGES)

.PHONY: test
test: fmt
	@go test $(GO_TEST_FLAGS) $(GO_PACKAGES)

.PHONY: vendor
vendor:
	@go mod tidy
	@go mod vendor