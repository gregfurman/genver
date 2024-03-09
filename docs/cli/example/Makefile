all: gen run

gen:
	go generate ./...

genver:
	genver --package=main "$$(go list -m -json all)"

run:
	go run *.go

.PHONY: all gen genver run