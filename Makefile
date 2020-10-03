GOOS := "darwin"

.PHONY: build
build:
	@go build -o leto

.PHONY: clean
clean:
	rm leto

.PHONY: test
test:
	@go test --cover ./...

.PHONY: fmt
fmt:
	@go fmt

.PHONY: validate
validate: build test fmt clean