fmt:
	@go fmt

build:
	@go build .

run: build
	@./resgen generate
