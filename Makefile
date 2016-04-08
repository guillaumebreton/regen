fmt:
	@go fmt $$(glide novendor)

build:
	@go build -o regen regen.go

run:
	@./regen generate examples -t examples/template.html
