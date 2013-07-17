%{

//TODO-
// (13:47) jnml@fsc-r550:~/src/github.com/cznic/golex/example2$ time ./example2 
// 1318 .go files, 10426201 bytes, 1854635 tokens
// 
// real	0m0.243s
// user	0m0.212s
// sys	0m0.036s
// (13:47) jnml@fsc-r550:~/src/github.com/cznic/golex/example2$

/*

Copyright (c) 2013 Go Authors. All rights reserved.
Use of this source code is governed by a BSD-style
license that can be found in the LICENSE file.

CAUTION: If this file is a Go source file (*.go), it was generated
automatically by '$ golex' from a *.l file - DO NOT EDIT in that case!

*/

package scanner

import (
	"fmt"
	"go/token"
	"strconv"
	"unicode"
	"unicode/utf8"

	"github.com/cznic/mathutil"
)

type Lexer struct {
	c       int
	col     int
	Errors  []error
	i0      int
	i       int
	lcol    int
	line    int
	ncol    int
	nline   int
	sc      int
	src     []byte
	val     []byte
}

func newLexer(src []byte) (l *Lexer) {
	l = &Lexer{
		src:   src,
		nline: 1,
		ncol:  0,
	}
	l.next()
	return
}

func (l *Lexer) next() int {
	if l.c != 0 {
		l.val = append(l.val, byte(l.c))
	}
	l.c = 0
	if l.i < len(l.src) {
		l.c = int(l.src[l.i])
		l.i++
	}
	switch l.c {
	case '\n':
		l.lcol = l.ncol
		l.nline++
		l.ncol = 0
	default:
		l.ncol++
	}
	return l.c
}

func (l *Lexer) err(s string, arg ...interface{}) {
	err := fmt.Errorf(fmt.Sprintf("%d:%d ", l.line, l.col)+s, arg...)
	l.Errors = append(l.Errors, err)
}

func (l *Lexer) Error(s string) {
	l.err(s)
}

func (l *Lexer) Scan() (tok token.Token, lval interface{}) {
	defer func() { fmt.Printf("%s(%d) %v\n", tok, int(tok), lval) }()
	const (
		INITIAL = iota
		S1
		S2
	)

	c := l.c
%}

%yyt l.sc
%yyc c
%yyn c = l.next()

int_lit         {decimal_lit}|{octal_lit}|{hex_lit}
decimal_lit     [1-9][0-9]*
octal_lit       0[0-7]*
hex_lit         0[xX][0-9a-fA-F]+

float_lit       {D}"."{D}?{E}?|{D}{E}|"."{D}{E}?
D        	[0-9]+
E        	[eE][-+]?[0-9]+

imaginary_ilit  {D}i
imaginary_lit   {float_lit}i
utf8_h2         [\xC2-\xDF]
utf8_h3         [\xE0-\xEF]
utf8_h4         [\xF0-\xF4]
utf8_cont       [\x80-\xBF]
utf8_2          {utf8_h2}{utf8_cont}
utf8_3          {utf8_h3}{utf8_cont}{utf8_cont}
utf8_4          {utf8_h4}{utf8_cont}{utf8_cont}{utf8_cont}
non_ascii       {utf8_2}|{utf8_3}|{utf8_4}

%yyc c
%yyn c = l.next()
%yyt l.sc

%x S1 S2

%%
			l.val = l.val[:0]
			l.i0, l.line, l.col = l.i, l.nline, l.ncol
			
<*>\0			l.i0++
			return token.EOF, lval

[ \t\n\r]+

\/\*([^*]|\*+[^*/])*\*+\/	|
\/\/.*			return token.COMMENT, string(l.val)

"!"			return token.NOT, lval
"!="			return token.NEQ, lval
"%"			return token.REM, lval
"%="			return token.REM_ASSIGN, lval
"&"			return token.AND, lval
"&&"			return token.LAND, lval
"&="			return token.AND_ASSIGN, lval
"&^"			return token.AND_NOT, lval
"&^="			return token.AND_NOT_ASSIGN, lval
"("			return token.LPAREN, lval
")"			return token.RPAREN, lval
"*"			return token.MUL, lval
"*="			return token.MUL_ASSIGN, lval
"+"			return token.ADD, lval
"++"			return token.INC, lval
"+="			return token.ADD_ASSIGN, lval
","			return token.COMMA, lval
"-"			return token.SUB, lval
"--"			return token.DEC, lval
"-="			return token.SUB_ASSIGN, lval
"."			return token.PERIOD, lval
"..."			return token.ELLIPSIS, lval
"/"			return token.QUO, lval
"/="			return token.QUO_ASSIGN, lval
":"			return token.COLON, lval
":="			return token.DEFINE, lval
";"			return token.SEMICOLON, lval
"<"			return token.LSS, lval
"<-"			return token.ARROW, lval
"<<"			return token.SHL, lval
"<<="			return token.SHL_ASSIGN, lval
"<="			return token.LEQ, lval
"=="			return token.EQL, lval
">"			return token.GTR, lval
">="			return token.GEQ, lval
">>"			return token.SHR, lval
">>="			return token.SHR_ASSIGN, lval
"["			return token.LBRACK, lval
"]"			return token.RBRACK, lval
"^"			return token.XOR, lval
"^="			return token.XOR_ASSIGN, lval
"{"			return token.LBRACE, lval
"|"			return token.OR, lval
"|="			return token.OR_ASSIGN, lval
"||"			return token.LOR, lval
"}"			return token.RBRACE, lval

break			return token.BREAK, lval
case			return token.CASE, lval
chan			return token.CHAN, lval
const			return token.CONST, lval
continue		return token.CONTINUE, lval
default			return token.DEFAULT, lval
defer			return token.DEFER, lval
else			return token.ELSE, lval
fallthrough		return token.FALLTHROUGH, lval
for			return token.FOR, lval
func			return token.FUNC, lval
go			return token.GO, lval
goto			return token.GOTO, lval
if			return token.IF, lval
import			return token.IMPORT, lval
interface		return token.INTERFACE, lval
map			return token.MAP, lval
package			return token.PACKAGE, lval
range			return token.RANGE, lval
return			return token.RETURN, lval
select			return token.SELECT, lval
struct			return token.STRUCT, lval
switch			return token.SWITCH, lval
type			return token.TYPE, lval
var			return token.VAR, lval

{imaginary_ilit}	return l.int(true)
{imaginary_lit}		return l.float(true)
{int_lit}		return l.int(false)
{float_lit}		return l.float(false)

\"			l.sc = S1
`			l.sc = S2

'(\\.|[^'])*'		if tok, lval = l.str(""); tok != token.STRING {
				return
			}
			return token.INT, int32(lval.(string)[0])

<S1>(\\.|[^\"])*\"	return l.str("\"")
<S2>([^`]|\n)*`		return l.str("`")

[a-zA-Z_][a-zA-Z_0-9]*
	if c >= '\xC2' && c <= '\xF4' {
		l.i--
		l.ncol--
		for {
			ln, cl, runepos := l.nline, l.ncol, l.i
			rune := l.getRune()
			if !(rune == '_' || unicode.IsLetter(rune) || unicode.IsDigit(rune)) {
				l.i = runepos
				c = l.next()
				l.nline, l.ncol = ln, cl
				break
			}
		}
	}
	return token.IDENT, string(l.src[l.i0-1:l.i-1])

{non_ascii}
	l.i = l.i0
	if rune := l.getRune(); !unicode.IsLetter(rune) {
		l.err("expected unicode letter, got %U", rune)
		return token.ILLEGAL, lval
	}

	for {
		ln, cl, runepos := l.nline, l.ncol, l.i
		rune := l.getRune()
		if !(rune == '_' || unicode.IsLetter(rune) || unicode.IsDigit(rune)) {
			l.i = runepos
			c = l.next()
			l.nline, l.ncol = ln, cl
			break
		}
	}
	return token.IDENT, string(l.src[l.i0:l.i])

%%
	return token.ILLEGAL, lval
}

func (l *Lexer) getRune() rune {
	if rune, size := utf8.DecodeRune(l.src[l.i:]); size != 0 {
		l.i += size
		return rune
	}

	return 0
}

func (l *Lexer) str(pref string) (tok token.Token, lval interface{}) {
	l.sc = 0
	s := pref + string(l.val)
	s, err := strconv.Unquote(s)
	if err != nil {
		l.err("string literal: %v", err)
		return token.ILLEGAL, lval
	}

	return token.STRING, s
}

func (l *Lexer) int(im bool) (tok token.Token, lval interface{}) {
	if im {
		l.val = l.val[:len(l.val)-1]
	}
	n, err := strconv.ParseUint(string(l.val), 0, 64)
	if err != nil {
		l.err("integer literal: %v", err)
		return token.ILLEGAL, lval
	}

	if im {
		return token.IMAG, complex(0, float64(n))
	}

	switch {
	case n < mathutil.MaxInt:
		lval = int(n)
	default:
		lval = n
	}
	return token.INT, lval
}

func (l *Lexer) float(im bool) (tok token.Token, lval interface{}) {
	if im {
		l.val = l.val[:len(l.val)-1]
	}
	n, err := strconv.ParseFloat(string(l.val), 64)
	if err != nil {
		l.err("float literal: %v", err)
		return token.ILLEGAL, lval
	}

	if im {
		return token.IMAG, complex(0, n)
	}

	return token.FLOAT, n
}
