# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=./calc

GOFILES=\
	calc.go\
	tokenizer.go\

CLEANFILES += tokenizer.go calc.go y.output

include $(GOROOT)/src/Make.cmd

tokenizer.go: tokenizer.l
	golex -t tokenizer.l | gofmt > tokenizer.go

calc.go: calc.y
	goyacc -o calc.go calc.y && gofmt -w calc.go
