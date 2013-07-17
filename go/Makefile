.PHONY: all clean

all: editor
	go vet
	go install
	make todo

editor: scanner.go
	go fmt
	go test -i
	go test
	go build

scanner.go: scanner.l
	golex -o $@ $<

todo:
	@grep -n ^[[:space:]]*_[[:space:]]*=[[:space:]][[:alpha:]][[:alnum:]]* *.go *.l || true
	@grep -n TODO *.go *.l || true
	@grep -n BUG *.go *.l || true
	@grep -n println *.go *.l || true

clean:
	@go clean
