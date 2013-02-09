all: lex.yy.go y.go
	go fmt
	go test
	go build
	go vet || true

lex.yy.go: cc.l
	golex cc.l

y.go: cc.y
	go tool yacc cc.y

todo:
	grep -n ^[[:space:]]_[[:space:]]*=[[:space:]][[:alnum:]] *.go || true
	grep -n TODO|println *.go || true

clean:
	rm -f lex.yy.go y.go parser.tab.c y.output cc.tab.c
