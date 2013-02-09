/*

Copyright (c) 2013 jnml. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

Adaption of http://www.lysator.liu.se/c/ANSI-C-grammar-l.html (no copyright[s]
declared there) as of 2013-01-28.

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
// Copyright (c) 2013 jnml. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// CAUTION: If this file is 'lex.yy.go', it was generated
// automatically from 'cc.l' - DO NOT EDIT in that case!

//TODO http://en.wikipedia.org/wiki/The_lexer_hack

package cc

import (
	"io"
)

// Errf is an error reporting function. If it returns false then further
// procesing (scanning, parsing, compiling) should be aborted.
type Errf func(file string, line, col int, msg string, args ...interface{}) bool

type lexer struct {
	file  string
	line  int
	col   int
	r     io.Reader
	buf   []byte
	token []byte
	errf  Errf
	err   error
}

func newLexer(file string, r io.Reader, errf Errf, buf []byte) (l *lexer) {
	const bufSize = 1<<16

	l = &lexer{
		file: file,
		line: 1,
		col:  1,
		r:    r,
		buf:  buf,
		errf: errf,
	}
	if cap(buf) == 0 {
		l.buf = make([]byte, 0, bufSize)
	}
	return
}

func (l *lexer) error(msg string, args ...interface{}) {
	if l.errf(l.file, l.line, l.col, msg, args...) {
		l.err = io.EOF
	}
}

func (l *lexer) peek() (c int) {
	if len(l.buf) == 0 {
		return l.read()
	}

	return int(l.buf[0])
}

func (l *lexer) read() (c int) {
	if l.err != nil {
		return -1
	}

	var n int
	if n, l.err = l.r.Read(l.buf[:cap(l.buf)]); n == 0 {
		switch {
		case l.err == nil:
			l.err = io.EOF
		case l.err != io.EOF:
			l.error(l.err.Error())
		}
		return -1
	}
	l.buf = l.buf[:n]
	return int(l.buf[0])	
}

func (l *lexer) next(curr int) (c int) {
	l.token = append(l.token, byte(curr))
	if len(l.buf) > 1 {
		l.buf = l.buf[1:]
		return int(l.buf[0])
	}
	return l.read()
}

func (l *lexer) scan() (ret int) {
	c := l.peek()

%}

D			[0-9]
L			[a-zA-Z_]
H			[a-fA-F0-9]
E			[Ee][+-]?{D}+
FS			(f|F|l|L)
IS			(u|U|l|L)*

%yyc c
%yyn c = l.next(c)

%%
	if ret != 0 {
		return
	}

	l.token = l.token[:0]

\/\*([^*]|\*+[^*/])*\*+\/

"auto"			// { count(); return(AUTO); }
"break"			// { count(); return(BREAK); }
"case"			// { count(); return(CASE); }
"char"			// { count(); return(CHAR); }
"const"			// { count(); return(CONST); }
"continue"		// { count(); return(CONTINUE); }
"default"		// { count(); return(DEFAULT); }
"do"			// { count(); return(DO); }
"double"		// { count(); return(DOUBLE); }
"else"			// { count(); return(ELSE); }
"enum"			// { count(); return(ENUM); }
"extern"		// { count(); return(EXTERN); }
"float"			// { count(); return(FLOAT); }
"for"			// { count(); return(FOR); }
"goto"			// { count(); return(GOTO); }
"if"			// { count(); return(IF); }
"int"			// { count(); return(INT); }
"long"			// { count(); return(LONG); }
"register"		// { count(); return(REGISTER); }
"return"		// { count(); return(RETURN); }
"short"			// { count(); return(SHORT); }
"signed"		// { count(); return(SIGNED); }
"sizeof"		// { count(); return(SIZEOF); }
"static"		// { count(); return(STATIC); }
"struct"		// { count(); return(STRUCT); }
"switch"		// { count(); return(SWITCH); }
"typedef"		// { count(); return(TYPEDEF); }
"union"			// { count(); return(UNION); }
"unsigned"		// { count(); return(UNSIGNED); }
"void"			// { count(); return(VOID); }
"volatile"		// { count(); return(VOLATILE); }
"while"			// { count(); return(WHILE); }

{L}({L}|{D})*		// { count(); return(check_type()); }

0[xX]{H}+{IS}?		// { count(); return(CONSTANT); }
0{D}+{IS}?		// { count(); return(CONSTANT); }
{D}+{IS}?		// { count(); return(CONSTANT); }
L?'(\\.|[^\\'])+'	// { count(); return(CONSTANT); }

{D}+{E}{FS}?		// { count(); return(CONSTANT); }
{D}*"."{D}+({E})?{FS}?	// { count(); return(CONSTANT); }
{D}+"."{D}*({E})?{FS}?	// { count(); return(CONSTANT); }

L?\"(\\.|[^\\"])*\"	// { count(); return(STRING_LITERAL); }

"..."			// { count(); return(ELLIPSIS); }
">>="			// { count(); return(RIGHT_ASSIGN); }
"<<="			// { count(); return(LEFT_ASSIGN); }
"+="			// { count(); return(ADD_ASSIGN); }
"-="			// { count(); return(SUB_ASSIGN); }
"*="			// { count(); return(MUL_ASSIGN); }
"/="			// { count(); return(DIV_ASSIGN); }
"%="			// { count(); return(MOD_ASSIGN); }
"&="			// { count(); return(AND_ASSIGN); }
"^="			// { count(); return(XOR_ASSIGN); }
"|="			// { count(); return(OR_ASSIGN); }
">>"			// { count(); return(RIGHT_OP); }
"<<"			// { count(); return(LEFT_OP); }
"++"			// { count(); return(INC_OP); }
"--"			// { count(); return(DEC_OP); }
"->"			// { count(); return(PTR_OP); }
"&&"			// { count(); return(AND_OP); }
"||"			// { count(); return(OR_OP); }
"<="			// { count(); return(LE_OP); }
">="			// { count(); return(GE_OP); }
"=="			// { count(); return(EQ_OP); }
"!="			// { count(); return(NE_OP); }
";"			// { count(); return(';'); }
("{"|"<%")		// { count(); return('{'); }
("}"|"%>")		// { count(); return('}'); }
","			// { count(); return(','); }
":"			// { count(); return(':'); }
"="			// { count(); return('='); }
"("			// { count(); return('('); }
")"			// { count(); return(')'); }
("["|"<:")		// { count(); return('['); }
("]"|":>")		// { count(); return(']'); }
"."			// { count(); return('.'); }
"&"			// { count(); return('&'); }
"!"			// { count(); return('!'); }
"~"			// { count(); return('~'); }
"-"			// { count(); return('-'); }
"+"			// { count(); return('+'); }
"*"			// { count(); return('*'); }
"/"			// { count(); return('/'); }
"%"			// { count(); return('%'); }
"<"			// { count(); return('<'); }
">"			// { count(); return('>'); }
"^"			// { count(); return('^'); }
"|"			// { count(); return('|'); }
"?"			// { count(); return('?'); }

[ \t\v\n\f]		// { count(); }
.			// { /* ignore bad characters */ }

%%
	panic("unreachable")
}

