all: lexer.go
	go fmt
	go test

lexer.go: lexer.l
	golex -o lexer.go lexer.l

todo:
	@grep -n ^[[:space:]]_[[:space:]]*=[[:space:]][[:alnum:]] *.go || true
	@grep -n TODO *.go || true

clean:
	rm -f lexer.go
