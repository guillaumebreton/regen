fmt:
	@go fmt $$(glide novendor)

lint:
	@glide novendor | sed "s|\./||g" | xargs -n1 golint

build:
	@go build -o regen regen.go


run:
	@./regen generate examples -t examples/template.html
	@./regen generate examples -t examples/template.txt
