fmt:
	@go fmt $$(glide novendor)

build:
	@go build $$(glide novendor)

run: build
	@./resgen generate
