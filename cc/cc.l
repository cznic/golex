/*

Copyright (c) 2013 Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Substantial parts of this file are an adaption of

	http://www.lysator.liu.se/c/ANSI-C-grammar-l.html

There is no copyright declared there as of 2013-01-28.

CAUTION: If this file is 'lex.yy.go', it was generated
automatically from 'cc.l' - DO NOT EDIT in that case!

TODO http://en.wikipedia.org/wiki/The_lexer_hack
TODO remove single char tokens from the lexical grammar

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

%{
package cc

import (
	"fmt"
	"io"
	"strconv"

	"github.com/cznic/strutil"
	"github.com/cznic/fileutil"
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
	const bufSize = 1<<16

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
		case !fileutil.IsEOF(l.err):
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

%}

D			[0-9]
L			[a-zA-Z_]
H			[a-fA-F0-9]
E			[Ee][+-]?{D}+
FS			(f|F|l|L)
IS			(u|U|l|L)*
LW			[ \t\v\f]

%yyc c
%yyn c = l.next(c)
%yyb l.prev == '\n' || l.prev == 0
%yyt sc

%x LINE FILE FN EOL

%%
			if ret != 0 {
				return
			}
		
			l.token = l.token[:0]

<*>\0			return 0

<INITIAL>^#		sc = LINE

<LINE>[ \t\v\f]+

<LINE>{D}+
			var err error
			if line, err = strconv.Atoi(string(l.token)); err != nil {
				panic("internal error")
			}
			sc = FILE

<LINE>.|\n		sc = EOL

<FILE>[ \t\v\f]+

<FILE>\"		sc = FN

<FILE>.|\n		sc = EOL

<FN>[^\n\"]*\"
			l.file = l.dict.Id(string(l.token[:len(l.token)-1]))
			l.line = line - 1
			sc = EOL

<FN>[^\"]*(\n|\0)	sc = INITIAL

<EOL>.*			sc = INITIAL

\/\*([^*]|\*+[^*/])*\*+\/

"auto"			ret = AUTO
"break"			ret = BREAK
"case"			ret = CASE
"char"			ret = CHAR
"const"			ret = CONST
"continue"		ret = CONTINUE
"default"		ret = DEFAULT
"do"			ret = DO
"double"		ret = DOUBLE
"else"			ret = ELSE
"enum"			ret = ENUM
"extern"		ret = EXTERN
"float"			ret = FLOAT
"for"			ret = FOR
"goto"			ret = GOTO
"if"			ret = IF
"int"			ret = INT
"long"			ret = LONG
"register"		ret = REGISTER
"return"		ret = RETURN
"short"			ret = SHORT
"signed"		ret = SIGNED
"sizeof"		ret = SIZEOF
"static"		ret = STATIC
"struct"		ret = STRUCT
"switch"		ret = SWITCH
"typedef"		ret = TYPEDEF
"union"			ret = UNION
"unsigned"		ret = UNSIGNED
"void"			ret = VOID
"volatile"		ret = VOLATILE
"while"			ret = WHILE

{L}({L}|{D})*		// { count(); return(check_type()); }

0[xX]{H}+{IS}?		// { count(); return(CONSTANT); }
0{D}+{IS}?		// { count(); return(CONSTANT); }
{D}+{IS}?		// { count(); return(CONSTANT); }
L?'(\\.|[^\\'])+'	// { count(); return(CONSTANT); }

{D}+{E}{FS}?		// { count(); return(CONSTANT); }
{D}*"."{D}+({E})?{FS}?	// { count(); return(CONSTANT); }
{D}+"."{D}*({E})?{FS}?	// { count(); return(CONSTANT); }

L?\"(\\.|[^\\"])*\"	// { count(); return(STRING_LITERAL); }

"..."			ret = ELLIPSIS
">>="			ret = RIGHT_ASSIGN
"<<="			ret = LEFT_ASSIGN
"+="			ret = ADD_ASSIGN
"-="			ret = SUB_ASSIGN
"*="			ret = MUL_ASSIGN
"/="			ret = DIV_ASSIGN
"%="			ret = MOD_ASSIGN
"&="			ret = AND_ASSIGN
"^="			ret = XOR_ASSIGN
"|="			ret = OR_ASSIGN
">>"			ret = RIGHT_OP
"<<"			ret = LEFT_OP
"++"			ret = INC_OP
"--"			ret = DEC_OP
"->"			ret = PTR_OP
"&&"			ret = AND_OP
"||"			ret = OR_OP
"<="			ret = LE_OP
">="			ret = GE_OP
"=="			ret = EQ_OP
"!="			ret = NE_OP

("{"|"<%")		ret = '{'
("}"|"%>")		ret = '}'
("["|"<:")		ret = '['
("]"|":>")		ret = ']'

({LW}|\n)+

.			ret = c

%%
			panic(fmt.Errorf(
				"%s.%d:%d: unreachable, sc %d, l.peek() %d, l.token %q",
				l.file, l.line, l.col, sc, l.peek(), l.token,
			))
}
