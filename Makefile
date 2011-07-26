# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=golex

GOFILES=\
	main.go\
	render.go\

include $(GOROOT)/src/Make.cmd

CLEANFILES+=example1 example1.go example2 example2.go lex.yy.go *~

example1: example1.l $(GOFILES)
	make install clean && make -f mkex1

example2: example2.l $(GOFILES)
	make install clean && make -f mkex2
