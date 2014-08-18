# Copyright (c) 2014 The golex Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

all: lex.yy.go y.go
	go fmt
	go test -i
	go test
	go build
	#TODO go vet || true

lex.yy.go: cc.l
	golex cc.l

y.go: cc.y
	go tool yacc cc.y

todo:
	@grep -n ^[[:space:]]_[[:space:]]*=[[:space:]][[:alnum:]] *.go || true
	@grep -n 'TODO|println' *.go || true

clean:
	rm -f parser.tab.c y.output cc.tab.c
