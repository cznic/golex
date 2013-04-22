/*

Copyright (c) 2013 jnml. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Substantial parts of this file are an adaption of

	http://www.lysator.liu.se/c/ANSI-C-grammar-l.html

There is no copyright declared there as of 2013-01-28.

---

ANSI C grammar, Lex specification

In 1985, Jeff Lee published this Lex specification together with a Yacc grammar
for the April 30, 1985 ANSI C draft.  Tom Stockfisch reposted both to
net.sources in 1987; that original, as mentioned in the answer to question
17.25 of the comp.lang.c FAQ, can be ftp'ed from ftp.uu.net, file
usenet/net.sources/ansi.c.grammar.Z.

I intend to keep this version as close to the current C Standard grammar as
possible; please let me know if you discover discrepancies.

Jutta Degener, 1995

*/
// Copyright (c) 2013 jnml. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CAUTION: If this file is 'lex.yy.go', it was generated
// automatically from 'cc.l' - DO NOT EDIT in that case!

//TODO http://en.wikipedia.org/wiki/The_lexer_hack
//TODO remove single char tokens from the lexical grammar

package cc

import (
	"fmt"
	"io"
	"strconv"

	"github.com/cznic/strutil"
)

// Errf is an error reporting function. If it returns false then further
// procesing (scanning, parsing, compiling) should be aborted.
type Errf func(file string, line, col int, msg string, args ...interface{}) bool

type lexer struct {
	dict  *strutil.Dict
	file  int
	line  int
	col   int
	r     io.Reader
	buf   []byte
	token []byte
	id    int
	errf  Errf
	err   error
	prev  int
}

// newLexer returns a new `lexer`. `buf` is the scanner buffer to use, which
// may be nil.
func newLexer(dict *strutil.Dict, file string, r io.Reader, errf Errf, buf []byte) (l *lexer) {
	const bufSize = 1 << 16

	l = &lexer{
		dict: dict,
		file: dict.Id(file),
		line: 1,
		col:  1,
		r:    r,
		buf:  buf,
		errf: errf,
	}
	if cap(buf) == 0 {
		l.buf = make([]byte, bufSize)
	}
	l.buf = l.buf[:0]
	return
}

func (l *lexer) s(id int) (str string) {
	var ok bool
	if str, ok = l.dict.S(id); !ok {
		str = "<MISSING:string from id>"
	}
	return
}

func (l *lexer) error(msg string, args ...interface{}) {
	if l.errf(l.s(l.file), l.line, l.col, msg, args...) {
		l.err = io.EOF
	}
}

func (l *lexer) peek() (c int) {
	// defer func() { println("peek", c) }()
	if len(l.buf) == 0 {
		return l.read()
	}

	return int(l.buf[0])
}

func (l *lexer) read() (c int) {
	// defer func() { println("read", c) }()
	if l.err != nil {
		return 0
	}

	var n int
	if n, l.err = l.r.Read(l.buf[:cap(l.buf)]); n == 0 {
		switch {
		case l.err == nil:
			l.err = io.EOF
		case l.err != io.EOF:
			l.error(l.err.Error())
		}
		l.buf = l.buf[:0]
		return 0
	}
	l.buf = l.buf[:n]
	return int(l.buf[0])
}

func (l *lexer) next(curr int) (c int) {
	// defer func() { println("next", c) }()
	l.prev = curr
	switch curr {
	default:
		l.col++
	case '\n':
		if curr == '\n' {
			l.line++
			l.col = 1
		}
	case 0:
	}
	l.token = append(l.token, byte(curr))
	if len(l.buf) > 1 {
		l.buf = l.buf[1:]
		return int(l.buf[0])
	}
	return l.read()
}

func (l *lexer) scan() (ret int) {
	const (
		INITIAL = iota
		LINE
		FILE
		FN
		EOL
	)
	sc := INITIAL
	c := l.peek()
	var line int

yystate0:

	if ret != 0 {
		return
	}

	l.token = l.token[:0]

	switch yyt, yyb := sc, l.prev == '\n' || l.prev == 0; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		if yyb {
			goto yystart261
		}
		goto yystart1
	case 1: // start condition: LINE
		if yyb {
			goto yystart263
		}
		goto yystart241
	case 2: // start condition: FILE
		if yyb {
			goto yystart264
		}
		goto yystart247
	case 3: // start condition: FN
		if yyb {
			goto yystart265
		}
		goto yystart252
	case 4: // start condition: EOL
		if yyb {
			goto yystart266
		}
		goto yystart259
	}

	goto yystate1 // silence unused label error
yystate1:
	c = l.next(c)
yystart1:
	switch {
	default:
		goto yystate3 // c >= '\x01' && c <= '\b' || c >= '\r' && c <= '\x1f' || c == '#' || c == '$' || c == '@' || c == '\\' || c == '`' || c >= '\u007f' && c <= 'ÿ'
	case c == '!':
		goto yystate6
	case c == '"':
		goto yystate8
	case c == '%':
		goto yystate12
	case c == '&':
		goto yystate15
	case c == '(':
		goto yystate22
	case c == ')':
		goto yystate23
	case c == '*':
		goto yystate24
	case c == '+':
		goto yystate26
	case c == ',':
		goto yystate29
	case c == '-':
		goto yystate30
	case c == '.':
		goto yystate34
	case c == '/':
		goto yystate42
	case c == '0':
		goto yystate47
	case c == ':':
		goto yystate70
	case c == ';':
		goto yystate72
	case c == '<':
		goto yystate73
	case c == '=':
		goto yystate79
	case c == '>':
		goto yystate81
	case c == '?':
		goto yystate85
	case c == 'L':
		goto yystate88
	case c == '[':
		goto yystate90
	case c == '\'':
		goto yystate18
	case c == '\n':
		goto yystate5
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c == ']':
		goto yystate91
	case c == '^':
		goto yystate92
	case c == 'a':
		goto yystate94
	case c == 'b':
		goto yystate98
	case c == 'c':
		goto yystate103
	case c == 'd':
		goto yystate119
	case c == 'e':
		goto yystate131
	case c == 'f':
		goto yystate143
	case c == 'g':
		goto yystate150
	case c == 'i':
		goto yystate154
	case c == 'l':
		goto yystate158
	case c == 'r':
		goto yystate162
	case c == 's':
		goto yystate174
	case c == 't':
		goto yystate202
	case c == 'u':
		goto yystate209
	case c == 'v':
		goto yystate220
	case c == 'w':
		goto yystate230
	case c == '{':
		goto yystate235
	case c == '|':
		goto yystate236
	case c == '}':
		goto yystate239
	case c == '~':
		goto yystate240
	case c >= '1' && c <= '9':
		goto yystate68
	case c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c == 'h' || c == 'j' || c == 'k' || c >= 'm' && c <= 'q' || c >= 'x' && c <= 'z':
		goto yystate86
	}

yystate2:
	c = l.next(c)
	goto yyrule1

yystate3:
	c = l.next(c)
	goto yyrule101

yystate4:
	c = l.next(c)
	switch {
	default:
		goto yyrule100
	case c >= '\t' && c <= '\f' || c == ' ':
		goto yystate5
	}

yystate5:
	c = l.next(c)
	switch {
	default:
		goto yyrule100
	case c >= '\t' && c <= '\f' || c == ' ':
		goto yystate5
	}

yystate6:
	c = l.next(c)
	switch {
	default:
		goto yyrule88
	case c == '=':
		goto yystate7
	}

yystate7:
	c = l.next(c)
	goto yyrule75

yystate8:
	c = l.next(c)
	switch {
	default:
		goto yyrule101
	case c == '"':
		goto yystate10
	case c == '\\':
		goto yystate11
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate9
	}

yystate9:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate10
	case c == '\\':
		goto yystate11
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate9
	}

yystate10:
	c = l.next(c)
	goto yyrule53

yystate11:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate9
	}

yystate12:
	c = l.next(c)
	switch {
	default:
		goto yyrule94
	case c == '=':
		goto yystate13
	case c == '>':
		goto yystate14
	}

yystate13:
	c = l.next(c)
	goto yyrule61

yystate14:
	c = l.next(c)
	goto yyrule78

yystate15:
	c = l.next(c)
	switch {
	default:
		goto yyrule87
	case c == '&':
		goto yystate16
	case c == '=':
		goto yystate17
	}

yystate16:
	c = l.next(c)
	goto yyrule70

yystate17:
	c = l.next(c)
	goto yyrule62

yystate18:
	c = l.next(c)
	switch {
	default:
		goto yyrule101
	case c == '\\':
		goto yystate21
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate19
	}

yystate19:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate20
	case c == '\\':
		goto yystate21
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate19
	}

yystate20:
	c = l.next(c)
	goto yyrule49

yystate21:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate19
	}

yystate22:
	c = l.next(c)
	goto yyrule82

yystate23:
	c = l.next(c)
	goto yyrule83

yystate24:
	c = l.next(c)
	switch {
	default:
		goto yyrule92
	case c == '=':
		goto yystate25
	}

yystate25:
	c = l.next(c)
	goto yyrule59

yystate26:
	c = l.next(c)
	switch {
	default:
		goto yyrule91
	case c == '+':
		goto yystate27
	case c == '=':
		goto yystate28
	}

yystate27:
	c = l.next(c)
	goto yyrule67

yystate28:
	c = l.next(c)
	goto yyrule57

yystate29:
	c = l.next(c)
	goto yyrule79

yystate30:
	c = l.next(c)
	switch {
	default:
		goto yyrule90
	case c == '-':
		goto yystate31
	case c == '=':
		goto yystate32
	case c == '>':
		goto yystate33
	}

yystate31:
	c = l.next(c)
	goto yyrule68

yystate32:
	c = l.next(c)
	goto yyrule58

yystate33:
	c = l.next(c)
	goto yyrule69

yystate34:
	c = l.next(c)
	switch {
	default:
		goto yyrule86
	case c == '.':
		goto yystate35
	case c >= '0' && c <= '9':
		goto yystate37
	}

yystate35:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate36
	}

yystate36:
	c = l.next(c)
	goto yyrule54

yystate37:
	c = l.next(c)
	switch {
	default:
		goto yyrule51
	case c == 'E' || c == 'e':
		goto yystate38
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate41
	case c >= '0' && c <= '9':
		goto yystate37
	}

yystate38:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate39
	case c >= '0' && c <= '9':
		goto yystate40
	}

yystate39:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate40
	}

yystate40:
	c = l.next(c)
	switch {
	default:
		goto yyrule51
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate41
	case c >= '0' && c <= '9':
		goto yystate40
	}

yystate41:
	c = l.next(c)
	goto yyrule51

yystate42:
	c = l.next(c)
	switch {
	default:
		goto yyrule93
	case c == '*':
		goto yystate43
	case c == '=':
		goto yystate46
	}

yystate43:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate44
	case c >= '\x01' && c <= ')' || c >= '+' && c <= 'ÿ':
		goto yystate43
	}

yystate44:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate44
	case c == '/':
		goto yystate45
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '.' || c >= '0' && c <= 'ÿ':
		goto yystate43
	}

yystate45:
	c = l.next(c)
	goto yyrule12

yystate46:
	c = l.next(c)
	goto yyrule60

yystate47:
	c = l.next(c)
	switch {
	default:
		goto yyrule48
	case c == '.':
		goto yystate48
	case c == 'E' || c == 'e':
		goto yystate59
	case c == 'L' || c == 'U' || c == 'l' || c == 'u':
		goto yystate64
	case c == 'X' || c == 'x':
		goto yystate65
	case c >= '0' && c <= '9':
		goto yystate58
	}

yystate48:
	c = l.next(c)
	switch {
	default:
		goto yyrule52
	case c == 'E' || c == 'e':
		goto yystate54
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate57
	case c >= '0' && c <= '9':
		goto yystate49
	}

yystate49:
	c = l.next(c)
	switch {
	default:
		goto yyrule51
	case c == 'E' || c == 'e':
		goto yystate50
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate53
	case c >= '0' && c <= '9':
		goto yystate49
	}

yystate50:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate51
	case c >= '0' && c <= '9':
		goto yystate52
	}

yystate51:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate52
	}

yystate52:
	c = l.next(c)
	switch {
	default:
		goto yyrule51
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate53
	case c >= '0' && c <= '9':
		goto yystate52
	}

yystate53:
	c = l.next(c)
	goto yyrule51

yystate54:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate55
	case c >= '0' && c <= '9':
		goto yystate56
	}

yystate55:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate56
	}

yystate56:
	c = l.next(c)
	switch {
	default:
		goto yyrule52
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate57
	case c >= '0' && c <= '9':
		goto yystate56
	}

yystate57:
	c = l.next(c)
	goto yyrule52

yystate58:
	c = l.next(c)
	switch {
	default:
		goto yyrule47
	case c == '.':
		goto yystate48
	case c == 'E' || c == 'e':
		goto yystate59
	case c == 'L' || c == 'U' || c == 'l' || c == 'u':
		goto yystate63
	case c >= '0' && c <= '9':
		goto yystate58
	}

yystate59:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate60
	case c >= '0' && c <= '9':
		goto yystate61
	}

yystate60:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate61
	}

yystate61:
	c = l.next(c)
	switch {
	default:
		goto yyrule50
	case c == 'F' || c == 'L' || c == 'f' || c == 'l':
		goto yystate62
	case c >= '0' && c <= '9':
		goto yystate61
	}

yystate62:
	c = l.next(c)
	goto yyrule50

yystate63:
	c = l.next(c)
	switch {
	default:
		goto yyrule47
	case c == 'L' || c == 'U' || c == 'l' || c == 'u':
		goto yystate63
	}

yystate64:
	c = l.next(c)
	switch {
	default:
		goto yyrule48
	case c == 'L' || c == 'U' || c == 'l' || c == 'u':
		goto yystate64
	}

yystate65:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate66
	}

yystate66:
	c = l.next(c)
	switch {
	default:
		goto yyrule46
	case c == 'L' || c == 'U' || c == 'l' || c == 'u':
		goto yystate67
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate66
	}

yystate67:
	c = l.next(c)
	switch {
	default:
		goto yyrule46
	case c == 'L' || c == 'U' || c == 'l' || c == 'u':
		goto yystate67
	}

yystate68:
	c = l.next(c)
	switch {
	default:
		goto yyrule48
	case c == '.':
		goto yystate48
	case c == 'E' || c == 'e':
		goto yystate59
	case c == 'L' || c == 'U' || c == 'l' || c == 'u':
		goto yystate64
	case c >= '0' && c <= '9':
		goto yystate69
	}

yystate69:
	c = l.next(c)
	switch {
	default:
		goto yyrule48
	case c == '.':
		goto yystate48
	case c == 'E' || c == 'e':
		goto yystate59
	case c == 'L' || c == 'U' || c == 'l' || c == 'u':
		goto yystate64
	case c >= '0' && c <= '9':
		goto yystate69
	}

yystate70:
	c = l.next(c)
	switch {
	default:
		goto yyrule80
	case c == '>':
		goto yystate71
	}

yystate71:
	c = l.next(c)
	goto yyrule85

yystate72:
	c = l.next(c)
	goto yyrule76

yystate73:
	c = l.next(c)
	switch {
	default:
		goto yyrule95
	case c == '%':
		goto yystate74
	case c == ':':
		goto yystate75
	case c == '<':
		goto yystate76
	case c == '=':
		goto yystate78
	}

yystate74:
	c = l.next(c)
	goto yyrule77

yystate75:
	c = l.next(c)
	goto yyrule84

yystate76:
	c = l.next(c)
	switch {
	default:
		goto yyrule66
	case c == '=':
		goto yystate77
	}

yystate77:
	c = l.next(c)
	goto yyrule56

yystate78:
	c = l.next(c)
	goto yyrule72

yystate79:
	c = l.next(c)
	switch {
	default:
		goto yyrule81
	case c == '=':
		goto yystate80
	}

yystate80:
	c = l.next(c)
	goto yyrule74

yystate81:
	c = l.next(c)
	switch {
	default:
		goto yyrule96
	case c == '=':
		goto yystate82
	case c == '>':
		goto yystate83
	}

yystate82:
	c = l.next(c)
	goto yyrule73

yystate83:
	c = l.next(c)
	switch {
	default:
		goto yyrule65
	case c == '=':
		goto yystate84
	}

yystate84:
	c = l.next(c)
	goto yyrule55

yystate85:
	c = l.next(c)
	goto yyrule99

yystate86:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate87:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate88:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == '"':
		goto yystate9
	case c == '\'':
		goto yystate89
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate89:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '\\':
		goto yystate21
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate19
	}

yystate90:
	c = l.next(c)
	goto yyrule84

yystate91:
	c = l.next(c)
	goto yyrule85

yystate92:
	c = l.next(c)
	switch {
	default:
		goto yyrule97
	case c == '=':
		goto yystate93
	}

yystate93:
	c = l.next(c)
	goto yyrule63

yystate94:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'u':
		goto yystate95
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate87
	}

yystate95:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate96
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate96:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'o':
		goto yystate97
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate97:
	c = l.next(c)
	switch {
	default:
		goto yyrule13
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate98:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'r':
		goto yystate99
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate87
	}

yystate99:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate100
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate100:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'a':
		goto yystate101
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate87
	}

yystate101:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'k':
		goto yystate102
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate87
	}

yystate102:
	c = l.next(c)
	switch {
	default:
		goto yyrule14
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate103:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'a':
		goto yystate104
	case c == 'h':
		goto yystate107
	case c == 'o':
		goto yystate110
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'g' || c >= 'i' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate104:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 's':
		goto yystate105
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate87
	}

yystate105:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate106
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate106:
	c = l.next(c)
	switch {
	default:
		goto yyrule15
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate107:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'a':
		goto yystate108
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate87
	}

yystate108:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'r':
		goto yystate109
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate87
	}

yystate109:
	c = l.next(c)
	switch {
	default:
		goto yyrule16
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate110:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'n':
		goto yystate111
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate111:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 's':
		goto yystate112
	case c == 't':
		goto yystate114
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate112:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate113
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate113:
	c = l.next(c)
	switch {
	default:
		goto yyrule17
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate114:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'i':
		goto yystate115
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate87
	}

yystate115:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'n':
		goto yystate116
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate116:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'u':
		goto yystate117
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate87
	}

yystate117:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate118
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate118:
	c = l.next(c)
	switch {
	default:
		goto yyrule18
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate119:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate120
	case c == 'o':
		goto yystate126
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate120:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'f':
		goto yystate121
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate87
	}

yystate121:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'a':
		goto yystate122
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate87
	}

yystate122:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'u':
		goto yystate123
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate87
	}

yystate123:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'l':
		goto yystate124
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate87
	}

yystate124:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate125
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate125:
	c = l.next(c)
	switch {
	default:
		goto yyrule19
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate126:
	c = l.next(c)
	switch {
	default:
		goto yyrule20
	case c == 'u':
		goto yystate127
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate87
	}

yystate127:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'b':
		goto yystate128
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c >= 'c' && c <= 'z':
		goto yystate87
	}

yystate128:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'l':
		goto yystate129
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate87
	}

yystate129:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate130
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate130:
	c = l.next(c)
	switch {
	default:
		goto yyrule21
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate131:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'l':
		goto yystate132
	case c == 'n':
		goto yystate135
	case c == 'x':
		goto yystate138
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c == 'm' || c >= 'o' && c <= 'w' || c == 'y' || c == 'z':
		goto yystate87
	}

yystate132:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 's':
		goto yystate133
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate87
	}

yystate133:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate134
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate134:
	c = l.next(c)
	switch {
	default:
		goto yyrule22
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate135:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'u':
		goto yystate136
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate87
	}

yystate136:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'm':
		goto yystate137
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'l' || c >= 'n' && c <= 'z':
		goto yystate87
	}

yystate137:
	c = l.next(c)
	switch {
	default:
		goto yyrule23
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate138:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate139
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate139:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate140
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate140:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'r':
		goto yystate141
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate87
	}

yystate141:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'n':
		goto yystate142
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate142:
	c = l.next(c)
	switch {
	default:
		goto yyrule24
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate143:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'l':
		goto yystate144
	case c == 'o':
		goto yystate148
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c == 'm' || c == 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate144:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'o':
		goto yystate145
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate145:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'a':
		goto yystate146
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate87
	}

yystate146:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate147
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate147:
	c = l.next(c)
	switch {
	default:
		goto yyrule25
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate148:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'r':
		goto yystate149
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate87
	}

yystate149:
	c = l.next(c)
	switch {
	default:
		goto yyrule26
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate150:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'o':
		goto yystate151
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate151:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate152
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate152:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'o':
		goto yystate153
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate153:
	c = l.next(c)
	switch {
	default:
		goto yyrule27
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate154:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'f':
		goto yystate155
	case c == 'n':
		goto yystate156
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate155:
	c = l.next(c)
	switch {
	default:
		goto yyrule28
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate156:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate157
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate157:
	c = l.next(c)
	switch {
	default:
		goto yyrule29
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate158:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'o':
		goto yystate159
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate159:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'n':
		goto yystate160
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate160:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'g':
		goto yystate161
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate87
	}

yystate161:
	c = l.next(c)
	switch {
	default:
		goto yyrule30
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate162:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate163
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate163:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'g':
		goto yystate164
	case c == 't':
		goto yystate170
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate164:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'i':
		goto yystate165
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate87
	}

yystate165:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 's':
		goto yystate166
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate87
	}

yystate166:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate167
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate167:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate168
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate168:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'r':
		goto yystate169
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate87
	}

yystate169:
	c = l.next(c)
	switch {
	default:
		goto yyrule31
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate170:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'u':
		goto yystate171
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate87
	}

yystate171:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'r':
		goto yystate172
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate87
	}

yystate172:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'n':
		goto yystate173
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate173:
	c = l.next(c)
	switch {
	default:
		goto yyrule32
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate174:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'h':
		goto yystate175
	case c == 'i':
		goto yystate179
	case c == 't':
		goto yystate188
	case c == 'w':
		goto yystate197
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'j' && c <= 's' || c == 'u' || c == 'v' || c >= 'x' && c <= 'z':
		goto yystate87
	}

yystate175:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'o':
		goto yystate176
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate176:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'r':
		goto yystate177
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate87
	}

yystate177:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate178
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate178:
	c = l.next(c)
	switch {
	default:
		goto yyrule33
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate179:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'g':
		goto yystate180
	case c == 'z':
		goto yystate184
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'y':
		goto yystate87
	}

yystate180:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'n':
		goto yystate181
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate181:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate182
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate182:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'd':
		goto yystate183
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate87
	}

yystate183:
	c = l.next(c)
	switch {
	default:
		goto yyrule34
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate184:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate185
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate185:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'o':
		goto yystate186
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate186:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'f':
		goto yystate187
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate87
	}

yystate187:
	c = l.next(c)
	switch {
	default:
		goto yyrule35
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate188:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'a':
		goto yystate189
	case c == 'r':
		goto yystate193
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate87
	}

yystate189:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate190
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate190:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'i':
		goto yystate191
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate87
	}

yystate191:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'c':
		goto yystate192
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate87
	}

yystate192:
	c = l.next(c)
	switch {
	default:
		goto yyrule36
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate193:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'u':
		goto yystate194
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate87
	}

yystate194:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'c':
		goto yystate195
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate87
	}

yystate195:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate196
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate196:
	c = l.next(c)
	switch {
	default:
		goto yyrule37
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate197:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'i':
		goto yystate198
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate87
	}

yystate198:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate199
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate199:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'c':
		goto yystate200
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate87
	}

yystate200:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'h':
		goto yystate201
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate87
	}

yystate201:
	c = l.next(c)
	switch {
	default:
		goto yyrule38
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate202:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'y':
		goto yystate203
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate87
	}

yystate203:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'p':
		goto yystate204
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate87
	}

yystate204:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate205
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate205:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'd':
		goto yystate206
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate87
	}

yystate206:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate207
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate207:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'f':
		goto yystate208
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate87
	}

yystate208:
	c = l.next(c)
	switch {
	default:
		goto yyrule39
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate209:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'n':
		goto yystate210
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate210:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'i':
		goto yystate211
	case c == 's':
		goto yystate214
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate87
	}

yystate211:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'o':
		goto yystate212
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate212:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'n':
		goto yystate213
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate213:
	c = l.next(c)
	switch {
	default:
		goto yyrule40
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate214:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'i':
		goto yystate215
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate87
	}

yystate215:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'g':
		goto yystate216
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate87
	}

yystate216:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'n':
		goto yystate217
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate87
	}

yystate217:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate218
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate218:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'd':
		goto yystate219
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate87
	}

yystate219:
	c = l.next(c)
	switch {
	default:
		goto yyrule41
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate220:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'o':
		goto yystate221
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate87
	}

yystate221:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'i':
		goto yystate222
	case c == 'l':
		goto yystate224
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c == 'j' || c == 'k' || c >= 'm' && c <= 'z':
		goto yystate87
	}

yystate222:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'd':
		goto yystate223
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'c' || c >= 'e' && c <= 'z':
		goto yystate87
	}

yystate223:
	c = l.next(c)
	switch {
	default:
		goto yyrule42
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate224:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'a':
		goto yystate225
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate87
	}

yystate225:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 't':
		goto yystate226
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate87
	}

yystate226:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'i':
		goto yystate227
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate87
	}

yystate227:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'l':
		goto yystate228
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate87
	}

yystate228:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate229
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate229:
	c = l.next(c)
	switch {
	default:
		goto yyrule43
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate230:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'h':
		goto yystate231
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate87
	}

yystate231:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'i':
		goto yystate232
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate87
	}

yystate232:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'l':
		goto yystate233
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate87
	}

yystate233:
	c = l.next(c)
	switch {
	default:
		goto yyrule45
	case c == 'e':
		goto yystate234
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate87
	}

yystate234:
	c = l.next(c)
	switch {
	default:
		goto yyrule44
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate87
	}

yystate235:
	c = l.next(c)
	goto yyrule77

yystate236:
	c = l.next(c)
	switch {
	default:
		goto yyrule98
	case c == '=':
		goto yystate237
	case c == '|':
		goto yystate238
	}

yystate237:
	c = l.next(c)
	goto yyrule64

yystate238:
	c = l.next(c)
	goto yyrule71

yystate239:
	c = l.next(c)
	goto yyrule78

yystate240:
	c = l.next(c)
	goto yyrule89

	goto yystate241 // silence unused label error
yystate241:
	c = l.next(c)
yystart241:
	switch {
	default:
		goto yystate242 // c >= '\x01' && c <= '\b' || c == '\n' || c >= '\r' && c <= '\x1f' || c >= '!' && c <= '/' || c >= ':' && c <= 'ÿ'
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate243
	case c == '\x00':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate245
	}

yystate242:
	c = l.next(c)
	goto yyrule5

yystate243:
	c = l.next(c)
	switch {
	default:
		goto yyrule3
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate244
	}

yystate244:
	c = l.next(c)
	switch {
	default:
		goto yyrule3
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate244
	}

yystate245:
	c = l.next(c)
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= '9':
		goto yystate246
	}

yystate246:
	c = l.next(c)
	switch {
	default:
		goto yyrule4
	case c >= '0' && c <= '9':
		goto yystate246
	}

	goto yystate247 // silence unused label error
yystate247:
	c = l.next(c)
yystart247:
	switch {
	default:
		goto yystate248 // c >= '\x01' && c <= '\b' || c == '\n' || c >= '\r' && c <= '\x1f' || c == '!' || c >= '#' && c <= 'ÿ'
	case c == '"':
		goto yystate251
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate249
	case c == '\x00':
		goto yystate2
	}

yystate248:
	c = l.next(c)
	goto yyrule8

yystate249:
	c = l.next(c)
	switch {
	default:
		goto yyrule6
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate250
	}

yystate250:
	c = l.next(c)
	switch {
	default:
		goto yyrule6
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate250
	}

yystate251:
	c = l.next(c)
	goto yyrule7

	goto yystate252 // silence unused label error
yystate252:
	c = l.next(c)
yystart252:
	switch {
	default:
		goto yystate254 // c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ'
	case c == '"':
		goto yystate258
	case c == '\n':
		goto yystate256
	case c == '\x00':
		goto yystate253
	}

yystate253:
	c = l.next(c)
	goto yyrule1

yystate254:
	c = l.next(c)
	switch {
	default:
		goto yystate254 // c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ'
	case c == '"':
		goto yystate258
	case c == '\n':
		goto yystate256
	case c == '\x00':
		goto yystate255
	}

yystate255:
	c = l.next(c)
	goto yyrule10

yystate256:
	c = l.next(c)
	switch {
	default:
		goto yyrule10
	case c == '\n':
		goto yystate256
	case c == '\x00':
		goto yystate255
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate257
	}

yystate257:
	c = l.next(c)
	switch {
	default:
		goto yyabort
	case c == '\n':
		goto yystate256
	case c == '\x00':
		goto yystate255
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ':
		goto yystate257
	}

yystate258:
	c = l.next(c)
	goto yyrule9

	goto yystate259 // silence unused label error
yystate259:
	c = l.next(c)
yystart259:
	switch {
	default:
		goto yyrule11
	case c == '\x00':
		goto yystate2
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate260
	}

yystate260:
	c = l.next(c)
	switch {
	default:
		goto yyrule11
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate260
	}

	goto yystate261 // silence unused label error
yystate261:
	c = l.next(c)
yystart261:
	switch {
	default:
		goto yystate3 // c >= '\x01' && c <= '\b' || c >= '\r' && c <= '\x1f' || c == '$' || c == '@' || c == '\\' || c == '`' || c >= '\u007f' && c <= 'ÿ'
	case c == '!':
		goto yystate6
	case c == '"':
		goto yystate8
	case c == '#':
		goto yystate262
	case c == '%':
		goto yystate12
	case c == '&':
		goto yystate15
	case c == '(':
		goto yystate22
	case c == ')':
		goto yystate23
	case c == '*':
		goto yystate24
	case c == '+':
		goto yystate26
	case c == ',':
		goto yystate29
	case c == '-':
		goto yystate30
	case c == '.':
		goto yystate34
	case c == '/':
		goto yystate42
	case c == '0':
		goto yystate47
	case c == ':':
		goto yystate70
	case c == ';':
		goto yystate72
	case c == '<':
		goto yystate73
	case c == '=':
		goto yystate79
	case c == '>':
		goto yystate81
	case c == '?':
		goto yystate85
	case c == 'L':
		goto yystate88
	case c == '[':
		goto yystate90
	case c == '\'':
		goto yystate18
	case c == '\n':
		goto yystate5
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate4
	case c == '\x00':
		goto yystate2
	case c == ']':
		goto yystate91
	case c == '^':
		goto yystate92
	case c == 'a':
		goto yystate94
	case c == 'b':
		goto yystate98
	case c == 'c':
		goto yystate103
	case c == 'd':
		goto yystate119
	case c == 'e':
		goto yystate131
	case c == 'f':
		goto yystate143
	case c == 'g':
		goto yystate150
	case c == 'i':
		goto yystate154
	case c == 'l':
		goto yystate158
	case c == 'r':
		goto yystate162
	case c == 's':
		goto yystate174
	case c == 't':
		goto yystate202
	case c == 'u':
		goto yystate209
	case c == 'v':
		goto yystate220
	case c == 'w':
		goto yystate230
	case c == '{':
		goto yystate235
	case c == '|':
		goto yystate236
	case c == '}':
		goto yystate239
	case c == '~':
		goto yystate240
	case c >= '1' && c <= '9':
		goto yystate68
	case c >= 'A' && c <= 'K' || c >= 'M' && c <= 'Z' || c == '_' || c == 'h' || c == 'j' || c == 'k' || c >= 'm' && c <= 'q' || c >= 'x' && c <= 'z':
		goto yystate86
	}

yystate262:
	c = l.next(c)
	goto yyrule2

	goto yystate263 // silence unused label error
yystate263:
	c = l.next(c)
yystart263:
	switch {
	default:
		goto yystate242 // c >= '\x01' && c <= '\b' || c == '\n' || c >= '\r' && c <= '\x1f' || c >= '!' && c <= '/' || c >= ':' && c <= 'ÿ'
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate243
	case c == '\x00':
		goto yystate2
	case c >= '0' && c <= '9':
		goto yystate245
	}

	goto yystate264 // silence unused label error
yystate264:
	c = l.next(c)
yystart264:
	switch {
	default:
		goto yystate248 // c >= '\x01' && c <= '\b' || c == '\n' || c >= '\r' && c <= '\x1f' || c == '!' || c >= '#' && c <= 'ÿ'
	case c == '"':
		goto yystate251
	case c == '\t' || c == '\v' || c == '\f' || c == ' ':
		goto yystate249
	case c == '\x00':
		goto yystate2
	}

	goto yystate265 // silence unused label error
yystate265:
	c = l.next(c)
yystart265:
	switch {
	default:
		goto yystate254 // c >= '\x01' && c <= '\t' || c >= '\v' && c <= '!' || c >= '#' && c <= 'ÿ'
	case c == '"':
		goto yystate258
	case c == '\n':
		goto yystate256
	case c == '\x00':
		goto yystate253
	}

	goto yystate266 // silence unused label error
yystate266:
	c = l.next(c)
yystart266:
	switch {
	default:
		goto yyrule11
	case c == '\x00':
		goto yystate2
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate260
	}

yyrule1: // \0
	{
		return 0
	}
yyrule2: // ^#
	{
		sc = LINE
		goto yystate0
	}
yyrule3: // [ \t\v\f]+

	goto yystate0
yyrule4: // {D}+
	{

		var err error
		if line, err = strconv.Atoi(string(l.token)); err != nil {
			panic("internal error")
		}
		sc = FILE
		goto yystate0
	}
yyrule5: // .|\n
	{
		sc = EOL
		goto yystate0
	}
yyrule6: // [ \t\v\f]+

	goto yystate0
yyrule7: // \"
	{
		sc = FN
		goto yystate0
	}
yyrule8: // .|\n
	{
		sc = EOL
		goto yystate0
	}
yyrule9: // [^\n\"]*\"
	{

		l.file = l.dict.Id(string(l.token[:len(l.token)-1]))
		l.line = line - 1
		sc = EOL
		goto yystate0
	}
yyrule10: // [^\"]*(\n|\0)
	{
		sc = INITIAL
		goto yystate0
	}
yyrule11: // .*
	{
		sc = INITIAL
		goto yystate0
	}
yyrule12: // \/\*([^*]|\*+[^*/])*\*+\/

	goto yystate0
yyrule13: // "auto"
	{
		ret = AUTO
		goto yystate0
	}
yyrule14: // "break"
	{
		ret = BREAK
		goto yystate0
	}
yyrule15: // "case"
	{
		ret = CASE
		goto yystate0
	}
yyrule16: // "char"
	{
		ret = CHAR
		goto yystate0
	}
yyrule17: // "const"
	{
		ret = CONST
		goto yystate0
	}
yyrule18: // "continue"
	{
		ret = CONTINUE
		goto yystate0
	}
yyrule19: // "default"
	{
		ret = DEFAULT
		goto yystate0
	}
yyrule20: // "do"
	{
		ret = DO
		goto yystate0
	}
yyrule21: // "double"
	{
		ret = DOUBLE
		goto yystate0
	}
yyrule22: // "else"
	{
		ret = ELSE
		goto yystate0
	}
yyrule23: // "enum"
	{
		ret = ENUM
		goto yystate0
	}
yyrule24: // "extern"
	{
		ret = EXTERN
		goto yystate0
	}
yyrule25: // "float"
	{
		ret = FLOAT
		goto yystate0
	}
yyrule26: // "for"
	{
		ret = FOR
		goto yystate0
	}
yyrule27: // "goto"
	{
		ret = GOTO
		goto yystate0
	}
yyrule28: // "if"
	{
		ret = IF
		goto yystate0
	}
yyrule29: // "int"
	{
		ret = INT
		goto yystate0
	}
yyrule30: // "long"
	{
		ret = LONG
		goto yystate0
	}
yyrule31: // "register"
	{
		ret = REGISTER
		goto yystate0
	}
yyrule32: // "return"
	{
		ret = RETURN
		goto yystate0
	}
yyrule33: // "short"
	{
		ret = SHORT
		goto yystate0
	}
yyrule34: // "signed"
	{
		ret = SIGNED
		goto yystate0
	}
yyrule35: // "sizeof"
	{
		ret = SIZEOF
		goto yystate0
	}
yyrule36: // "static"
	{
		ret = STATIC
		goto yystate0
	}
yyrule37: // "struct"
	{
		ret = STRUCT
		goto yystate0
	}
yyrule38: // "switch"
	{
		ret = SWITCH
		goto yystate0
	}
yyrule39: // "typedef"
	{
		ret = TYPEDEF
		goto yystate0
	}
yyrule40: // "union"
	{
		ret = UNION
		goto yystate0
	}
yyrule41: // "unsigned"
	{
		ret = UNSIGNED
		goto yystate0
	}
yyrule42: // "void"
	{
		ret = VOID
		goto yystate0
	}
yyrule43: // "volatile"
	{
		ret = VOLATILE
		goto yystate0
	}
yyrule44: // "while"
	{
		ret = WHILE
		goto yystate0
	}
yyrule45: // {L}({L}|{D})*
	{
		// { count(); return(check_type()); }
		goto yystate0
	}
yyrule46: // 0[xX]{H}+{IS}?
	{
		// { count(); return(CONSTANT); }
		goto yystate0
	}
yyrule47: // 0{D}+{IS}?
	{
		// { count(); return(CONSTANT); }
		goto yystate0
	}
yyrule48: // {D}+{IS}?
	{
		// { count(); return(CONSTANT); }
		goto yystate0
	}
yyrule49: // L?'(\\.|[^\\'])+'
	{
		// { count(); return(CONSTANT); }
		goto yystate0
	}
yyrule50: // {D}+{E}{FS}?
	{
		// { count(); return(CONSTANT); }
		goto yystate0
	}
yyrule51: // {D}*"."{D}+({E})?{FS}?
	{
		// { count(); return(CONSTANT); }
		goto yystate0
	}
yyrule52: // {D}+"."{D}*({E})?{FS}?
	{
		// { count(); return(CONSTANT); }
		goto yystate0
	}
yyrule53: // L?\"(\\.|[^\\"])*\"
	{
		// { count(); return(STRING_LITERAL); }
		goto yystate0
	}
yyrule54: // "..."
	{
		ret = ELLIPSIS
		goto yystate0
	}
yyrule55: // ">>="
	{
		ret = RIGHT_ASSIGN
		goto yystate0
	}
yyrule56: // "<<="
	{
		ret = LEFT_ASSIGN
		goto yystate0
	}
yyrule57: // "+="
	{
		ret = ADD_ASSIGN
		goto yystate0
	}
yyrule58: // "-="
	{
		ret = SUB_ASSIGN
		goto yystate0
	}
yyrule59: // "*="
	{
		ret = MUL_ASSIGN
		goto yystate0
	}
yyrule60: // "/="
	{
		ret = DIV_ASSIGN
		goto yystate0
	}
yyrule61: // "%="
	{
		ret = MOD_ASSIGN
		goto yystate0
	}
yyrule62: // "&="
	{
		ret = AND_ASSIGN
		goto yystate0
	}
yyrule63: // "^="
	{
		ret = XOR_ASSIGN
		goto yystate0
	}
yyrule64: // "|="
	{
		ret = OR_ASSIGN
		goto yystate0
	}
yyrule65: // ">>"
	{
		ret = RIGHT_OP
		goto yystate0
	}
yyrule66: // "<<"
	{
		ret = LEFT_OP
		goto yystate0
	}
yyrule67: // "++"
	{
		ret = INC_OP
		goto yystate0
	}
yyrule68: // "--"
	{
		ret = DEC_OP
		goto yystate0
	}
yyrule69: // "->"
	{
		ret = PTR_OP
		goto yystate0
	}
yyrule70: // "&&"
	{
		ret = AND_OP
		goto yystate0
	}
yyrule71: // "||"
	{
		ret = OR_OP
		goto yystate0
	}
yyrule72: // "<="
	{
		ret = LE_OP
		goto yystate0
	}
yyrule73: // ">="
	{
		ret = GE_OP
		goto yystate0
	}
yyrule74: // "=="
	{
		ret = EQ_OP
		goto yystate0
	}
yyrule75: // "!="
	{
		ret = NE_OP
		goto yystate0
	}
yyrule76: // ";"
	{
		// { count(); return(';'); }
		goto yystate0
	}
yyrule77: // ("{"|"<%")
	{
		// { count(); return('{'); }
		goto yystate0
	}
yyrule78: // ("}"|"%>")
	{
		// { count(); return('}'); }
		goto yystate0
	}
yyrule79: // ","
	{
		// { count(); return(','); }
		goto yystate0
	}
yyrule80: // ":"
	{
		// { count(); return(':'); }
		goto yystate0
	}
yyrule81: // "="
	{
		// { count(); return('='); }
		goto yystate0
	}
yyrule82: // "("
	{
		// { count(); return('('); }
		goto yystate0
	}
yyrule83: // ")"
	{
		// { count(); return(')'); }
		goto yystate0
	}
yyrule84: // ("["|"<:")
	{
		// { count(); return('['); }
		goto yystate0
	}
yyrule85: // ("]"|":>")
	{
		// { count(); return(']'); }
		goto yystate0
	}
yyrule86: // "."
	{
		// { count(); return('.'); }
		goto yystate0
	}
yyrule87: // "&"
	{
		// { count(); return('&'); }
		goto yystate0
	}
yyrule88: // "!"
	{
		// { count(); return('!'); }
		goto yystate0
	}
yyrule89: // "~"
	{
		// { count(); return('~'); }
		goto yystate0
	}
yyrule90: // "-"
	{
		// { count(); return('-'); }
		goto yystate0
	}
yyrule91: // "+"
	{
		// { count(); return('+'); }
		goto yystate0
	}
yyrule92: // "*"
	{
		// { count(); return('*'); }
		goto yystate0
	}
yyrule93: // "/"
	{
		// { count(); return('/'); }
		goto yystate0
	}
yyrule94: // "%"
	{
		// { count(); return('%'); }
		goto yystate0
	}
yyrule95: // "<"
	{
		// { count(); return('<'); }
		goto yystate0
	}
yyrule96: // ">"
	{
		// { count(); return('>'); }
		goto yystate0
	}
yyrule97: // "^"
	{
		// { count(); return('^'); }
		goto yystate0
	}
yyrule98: // "|"
	{
		// { count(); return('|'); }
		goto yystate0
	}
yyrule99: // "?"
	{
		// { count(); return('?'); }
		goto yystate0
	}
yyrule100: // ({LW}|\n)+
	{
		// { count(); }
		goto yystate0
	}
yyrule101: // .
	{
		ret = c
		goto yystate0
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
	panic(fmt.Errorf(
		"%s.%d:%d: unreachable, sc %d, l.peek() %d, l.token %q",
		l.file, l.line, l.col, sc, l.peek(), l.token,
	))
}
