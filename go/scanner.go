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
	c      int
	col    int
	Errors []error
	i0     int
	i      int
	lcol   int
	line   int
	ncol   int
	nline  int
	sc     int
	src    []byte
	val    []byte
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

yystate0:

	l.val = l.val[:0]
	l.i0, l.line, l.col = l.i, l.nline, l.ncol

	switch yyt := l.sc; yyt {
	default:
		panic(fmt.Errorf(`invalid start condition %d`, yyt))
	case 0: // start condition: INITIAL
		goto yystart1
	case 1: // start condition: S1
		goto yystart191
	case 2: // start condition: S2
		goto yystart196
	}

	goto yystate1 // silence unused label error
yystate1:
	c = l.next()
yystart1:
	switch {
	default:
		goto yyabort
	case c == '!':
		goto yystate4
	case c == '"':
		goto yystate6
	case c == '%':
		goto yystate7
	case c == '&':
		goto yystate9
	case c == '(':
		goto yystate18
	case c == ')':
		goto yystate19
	case c == '*':
		goto yystate20
	case c == '+':
		goto yystate22
	case c == ',':
		goto yystate25
	case c == '-':
		goto yystate26
	case c == '.':
		goto yystate29
	case c == '/':
		goto yystate37
	case c == '0':
		goto yystate43
	case c == ':':
		goto yystate50
	case c == ';':
		goto yystate52
	case c == '<':
		goto yystate53
	case c == '=':
		goto yystate58
	case c == '>':
		goto yystate60
	case c == '[':
		goto yystate65
	case c == '\'':
		goto yystate14
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate3
	case c == '\x00':
		goto yystate2
	case c == ']':
		goto yystate66
	case c == '^':
		goto yystate67
	case c == '`':
		goto yystate69
	case c == 'b':
		goto yystate70
	case c == 'c':
		goto yystate75
	case c == 'd':
		goto yystate91
	case c == 'e':
		goto yystate100
	case c == 'f':
		goto yystate104
	case c == 'g':
		goto yystate120
	case c == 'i':
		goto yystate124
	case c == 'm':
		goto yystate139
	case c == 'p':
		goto yystate142
	case c == 'r':
		goto yystate149
	case c == 's':
		goto yystate159
	case c == 't':
		goto yystate175
	case c == 'v':
		goto yystate179
	case c == '{':
		goto yystate182
	case c == '|':
		goto yystate183
	case c == '}':
		goto yystate186
	case c >= '1' && c <= '9':
		goto yystate49
	case c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'h' || c >= 'j' && c <= 'l' || c == 'n' || c == 'o' || c == 'q' || c == 'u' || c >= 'w' && c <= 'z':
		goto yystate64
	case c >= 'Â' && c <= 'ß':
		goto yystate187
	case c >= 'à' && c <= 'ï':
		goto yystate189
	case c >= 'ð' && c <= 'ô':
		goto yystate190
	}

yystate2:
	c = l.next()
	goto yyrule1

yystate3:
	c = l.next()
	switch {
	default:
		goto yyrule2
	case c == '\t' || c == '\n' || c == '\r' || c == ' ':
		goto yystate3
	}

yystate4:
	c = l.next()
	switch {
	default:
		goto yyrule5
	case c == '=':
		goto yystate5
	}

yystate5:
	c = l.next()
	goto yyrule6

yystate6:
	c = l.next()
	goto yyrule80

yystate7:
	c = l.next()
	switch {
	default:
		goto yyrule7
	case c == '=':
		goto yystate8
	}

yystate8:
	c = l.next()
	goto yyrule8

yystate9:
	c = l.next()
	switch {
	default:
		goto yyrule9
	case c == '&':
		goto yystate10
	case c == '=':
		goto yystate11
	case c == '^':
		goto yystate12
	}

yystate10:
	c = l.next()
	goto yyrule10

yystate11:
	c = l.next()
	goto yyrule11

yystate12:
	c = l.next()
	switch {
	default:
		goto yyrule12
	case c == '=':
		goto yystate13
	}

yystate13:
	c = l.next()
	goto yyrule13

yystate14:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate15
	case c == '\\':
		goto yystate16
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate14
	}

yystate15:
	c = l.next()
	goto yyrule82

yystate16:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '\'':
		goto yystate17
	case c == '\\':
		goto yystate16
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate14
	}

yystate17:
	c = l.next()
	switch {
	default:
		goto yyrule82
	case c == '\'':
		goto yystate15
	case c == '\\':
		goto yystate16
	case c >= '\x01' && c <= '&' || c >= '(' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate14
	}

yystate18:
	c = l.next()
	goto yyrule14

yystate19:
	c = l.next()
	goto yyrule15

yystate20:
	c = l.next()
	switch {
	default:
		goto yyrule16
	case c == '=':
		goto yystate21
	}

yystate21:
	c = l.next()
	goto yyrule17

yystate22:
	c = l.next()
	switch {
	default:
		goto yyrule18
	case c == '+':
		goto yystate23
	case c == '=':
		goto yystate24
	}

yystate23:
	c = l.next()
	goto yyrule19

yystate24:
	c = l.next()
	goto yyrule20

yystate25:
	c = l.next()
	goto yyrule21

yystate26:
	c = l.next()
	switch {
	default:
		goto yyrule22
	case c == '-':
		goto yystate27
	case c == '=':
		goto yystate28
	}

yystate27:
	c = l.next()
	goto yyrule23

yystate28:
	c = l.next()
	goto yyrule24

yystate29:
	c = l.next()
	switch {
	default:
		goto yyrule25
	case c == '.':
		goto yystate30
	case c >= '0' && c <= '9':
		goto yystate32
	}

yystate30:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate31
	}

yystate31:
	c = l.next()
	goto yyrule26

yystate32:
	c = l.next()
	switch {
	default:
		goto yyrule79
	case c == 'E' || c == 'e':
		goto yystate33
	case c == 'i':
		goto yystate36
	case c >= '0' && c <= '9':
		goto yystate32
	}

yystate33:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '+' || c == '-':
		goto yystate34
	case c >= '0' && c <= '9':
		goto yystate35
	}

yystate34:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9':
		goto yystate35
	}

yystate35:
	c = l.next()
	switch {
	default:
		goto yyrule79
	case c == 'i':
		goto yystate36
	case c >= '0' && c <= '9':
		goto yystate35
	}

yystate36:
	c = l.next()
	goto yyrule77

yystate37:
	c = l.next()
	switch {
	default:
		goto yyrule27
	case c == '*':
		goto yystate38
	case c == '/':
		goto yystate41
	case c == '=':
		goto yystate42
	}

yystate38:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate39
	case c >= '\x01' && c <= ')' || c >= '+' && c <= 'ÿ':
		goto yystate38
	}

yystate39:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '*':
		goto yystate39
	case c == '/':
		goto yystate40
	case c >= '\x01' && c <= ')' || c >= '+' && c <= '.' || c >= '0' && c <= 'ÿ':
		goto yystate38
	}

yystate40:
	c = l.next()
	goto yyrule3

yystate41:
	c = l.next()
	switch {
	default:
		goto yyrule4
	case c >= '\x01' && c <= '\t' || c >= '\v' && c <= 'ÿ':
		goto yystate41
	}

yystate42:
	c = l.next()
	goto yyrule28

yystate43:
	c = l.next()
	switch {
	default:
		goto yyrule78
	case c == '.':
		goto yystate32
	case c == '8' || c == '9':
		goto yystate45
	case c == 'E' || c == 'e':
		goto yystate33
	case c == 'X' || c == 'x':
		goto yystate47
	case c == 'i':
		goto yystate46
	case c >= '0' && c <= '7':
		goto yystate44
	}

yystate44:
	c = l.next()
	switch {
	default:
		goto yyrule78
	case c == '.':
		goto yystate32
	case c == '8' || c == '9':
		goto yystate45
	case c == 'E' || c == 'e':
		goto yystate33
	case c == 'i':
		goto yystate46
	case c >= '0' && c <= '7':
		goto yystate44
	}

yystate45:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '.':
		goto yystate32
	case c == 'E' || c == 'e':
		goto yystate33
	case c == 'i':
		goto yystate46
	case c >= '0' && c <= '9':
		goto yystate45
	}

yystate46:
	c = l.next()
	goto yyrule76

yystate47:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate48
	}

yystate48:
	c = l.next()
	switch {
	default:
		goto yyrule78
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'F' || c >= 'a' && c <= 'f':
		goto yystate48
	}

yystate49:
	c = l.next()
	switch {
	default:
		goto yyrule78
	case c == '.':
		goto yystate32
	case c == 'E' || c == 'e':
		goto yystate33
	case c == 'i':
		goto yystate46
	case c >= '0' && c <= '9':
		goto yystate49
	}

yystate50:
	c = l.next()
	switch {
	default:
		goto yyrule29
	case c == '=':
		goto yystate51
	}

yystate51:
	c = l.next()
	goto yyrule30

yystate52:
	c = l.next()
	goto yyrule31

yystate53:
	c = l.next()
	switch {
	default:
		goto yyrule32
	case c == '-':
		goto yystate54
	case c == '<':
		goto yystate55
	case c == '=':
		goto yystate57
	}

yystate54:
	c = l.next()
	goto yyrule33

yystate55:
	c = l.next()
	switch {
	default:
		goto yyrule34
	case c == '=':
		goto yystate56
	}

yystate56:
	c = l.next()
	goto yyrule35

yystate57:
	c = l.next()
	goto yyrule36

yystate58:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '=':
		goto yystate59
	}

yystate59:
	c = l.next()
	goto yyrule37

yystate60:
	c = l.next()
	switch {
	default:
		goto yyrule38
	case c == '=':
		goto yystate61
	case c == '>':
		goto yystate62
	}

yystate61:
	c = l.next()
	goto yyrule39

yystate62:
	c = l.next()
	switch {
	default:
		goto yyrule40
	case c == '=':
		goto yystate63
	}

yystate63:
	c = l.next()
	goto yyrule41

yystate64:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate65:
	c = l.next()
	goto yyrule42

yystate66:
	c = l.next()
	goto yyrule43

yystate67:
	c = l.next()
	switch {
	default:
		goto yyrule44
	case c == '=':
		goto yystate68
	}

yystate68:
	c = l.next()
	goto yyrule45

yystate69:
	c = l.next()
	goto yyrule81

yystate70:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'r':
		goto yystate71
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate64
	}

yystate71:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate72
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate72:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate73
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate64
	}

yystate73:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'k':
		goto yystate74
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate64
	}

yystate74:
	c = l.next()
	switch {
	default:
		goto yyrule51
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate75:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate76
	case c == 'h':
		goto yystate79
	case c == 'o':
		goto yystate82
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'g' || c >= 'i' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate64
	}

yystate76:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 's':
		goto yystate77
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate64
	}

yystate77:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate78
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate78:
	c = l.next()
	switch {
	default:
		goto yyrule52
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate79:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate80
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate64
	}

yystate80:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'n':
		goto yystate81
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate64
	}

yystate81:
	c = l.next()
	switch {
	default:
		goto yyrule53
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate82:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'n':
		goto yystate83
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate64
	}

yystate83:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 's':
		goto yystate84
	case c == 't':
		goto yystate86
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate84:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 't':
		goto yystate85
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate85:
	c = l.next()
	switch {
	default:
		goto yyrule54
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate86:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'i':
		goto yystate87
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate64
	}

yystate87:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'n':
		goto yystate88
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate64
	}

yystate88:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'u':
		goto yystate89
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate64
	}

yystate89:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate90
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate90:
	c = l.next()
	switch {
	default:
		goto yyrule55
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate91:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate92
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate92:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'f':
		goto yystate93
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate64
	}

yystate93:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate94
	case c == 'e':
		goto yystate98
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate94:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'u':
		goto yystate95
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate64
	}

yystate95:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'l':
		goto yystate96
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate64
	}

yystate96:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 't':
		goto yystate97
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate97:
	c = l.next()
	switch {
	default:
		goto yyrule56
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate98:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'r':
		goto yystate99
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate64
	}

yystate99:
	c = l.next()
	switch {
	default:
		goto yyrule57
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate100:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'l':
		goto yystate101
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate64
	}

yystate101:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 's':
		goto yystate102
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'r' || c >= 't' && c <= 'z':
		goto yystate64
	}

yystate102:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate103
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate103:
	c = l.next()
	switch {
	default:
		goto yyrule58
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate104:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate105
	case c == 'o':
		goto yystate115
	case c == 'u':
		goto yystate117
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'n' || c >= 'p' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate64
	}

yystate105:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'l':
		goto yystate106
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate64
	}

yystate106:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'l':
		goto yystate107
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate64
	}

yystate107:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 't':
		goto yystate108
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate108:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'h':
		goto yystate109
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate64
	}

yystate109:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'r':
		goto yystate110
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate64
	}

yystate110:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'o':
		goto yystate111
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate64
	}

yystate111:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'u':
		goto yystate112
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate64
	}

yystate112:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'g':
		goto yystate113
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate64
	}

yystate113:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'h':
		goto yystate114
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate64
	}

yystate114:
	c = l.next()
	switch {
	default:
		goto yyrule59
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate115:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'r':
		goto yystate116
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate64
	}

yystate116:
	c = l.next()
	switch {
	default:
		goto yyrule60
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate117:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'n':
		goto yystate118
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate64
	}

yystate118:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'c':
		goto yystate119
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate64
	}

yystate119:
	c = l.next()
	switch {
	default:
		goto yyrule61
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate120:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'o':
		goto yystate121
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate64
	}

yystate121:
	c = l.next()
	switch {
	default:
		goto yyrule62
	case c == 't':
		goto yystate122
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate122:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'o':
		goto yystate123
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate64
	}

yystate123:
	c = l.next()
	switch {
	default:
		goto yyrule63
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate124:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'f':
		goto yystate125
	case c == 'm':
		goto yystate126
	case c == 'n':
		goto yystate131
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'l' || c >= 'o' && c <= 'z':
		goto yystate64
	}

yystate125:
	c = l.next()
	switch {
	default:
		goto yyrule64
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate126:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'p':
		goto yystate127
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate64
	}

yystate127:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'o':
		goto yystate128
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'n' || c >= 'p' && c <= 'z':
		goto yystate64
	}

yystate128:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'r':
		goto yystate129
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate64
	}

yystate129:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 't':
		goto yystate130
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate130:
	c = l.next()
	switch {
	default:
		goto yyrule65
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate131:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 't':
		goto yystate132
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate132:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate133
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate133:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'r':
		goto yystate134
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate64
	}

yystate134:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'f':
		goto yystate135
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'e' || c >= 'g' && c <= 'z':
		goto yystate64
	}

yystate135:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate136
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate64
	}

yystate136:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'c':
		goto yystate137
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate64
	}

yystate137:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate138
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate138:
	c = l.next()
	switch {
	default:
		goto yyrule66
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate139:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate140
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate64
	}

yystate140:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'p':
		goto yystate141
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate64
	}

yystate141:
	c = l.next()
	switch {
	default:
		goto yyrule67
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate142:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate143
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate64
	}

yystate143:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'c':
		goto yystate144
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate64
	}

yystate144:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'k':
		goto yystate145
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'j' || c >= 'l' && c <= 'z':
		goto yystate64
	}

yystate145:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate146
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate64
	}

yystate146:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'g':
		goto yystate147
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate64
	}

yystate147:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate148
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate148:
	c = l.next()
	switch {
	default:
		goto yyrule68
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate149:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate150
	case c == 'e':
		goto yystate154
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate150:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'n':
		goto yystate151
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate64
	}

yystate151:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'g':
		goto yystate152
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'f' || c >= 'h' && c <= 'z':
		goto yystate64
	}

yystate152:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate153
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate153:
	c = l.next()
	switch {
	default:
		goto yyrule69
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate154:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 't':
		goto yystate155
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate155:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'u':
		goto yystate156
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate64
	}

yystate156:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'r':
		goto yystate157
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate64
	}

yystate157:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'n':
		goto yystate158
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'm' || c >= 'o' && c <= 'z':
		goto yystate64
	}

yystate158:
	c = l.next()
	switch {
	default:
		goto yyrule70
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate159:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate160
	case c == 't':
		goto yystate165
	case c == 'w':
		goto yystate170
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 's' || c == 'u' || c == 'v' || c >= 'x' && c <= 'z':
		goto yystate64
	}

yystate160:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'l':
		goto yystate161
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'k' || c >= 'm' && c <= 'z':
		goto yystate64
	}

yystate161:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate162
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate162:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'c':
		goto yystate163
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate64
	}

yystate163:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 't':
		goto yystate164
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate164:
	c = l.next()
	switch {
	default:
		goto yyrule71
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate165:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'r':
		goto yystate166
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate64
	}

yystate166:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'u':
		goto yystate167
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 't' || c >= 'v' && c <= 'z':
		goto yystate64
	}

yystate167:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'c':
		goto yystate168
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate64
	}

yystate168:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 't':
		goto yystate169
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate169:
	c = l.next()
	switch {
	default:
		goto yyrule72
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate170:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'i':
		goto yystate171
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'h' || c >= 'j' && c <= 'z':
		goto yystate64
	}

yystate171:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 't':
		goto yystate172
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 's' || c >= 'u' && c <= 'z':
		goto yystate64
	}

yystate172:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'c':
		goto yystate173
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c == 'a' || c == 'b' || c >= 'd' && c <= 'z':
		goto yystate64
	}

yystate173:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'h':
		goto yystate174
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'g' || c >= 'i' && c <= 'z':
		goto yystate64
	}

yystate174:
	c = l.next()
	switch {
	default:
		goto yyrule73
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate175:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'y':
		goto yystate176
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'x' || c == 'z':
		goto yystate64
	}

yystate176:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'p':
		goto yystate177
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'o' || c >= 'q' && c <= 'z':
		goto yystate64
	}

yystate177:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'e':
		goto yystate178
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'd' || c >= 'f' && c <= 'z':
		goto yystate64
	}

yystate178:
	c = l.next()
	switch {
	default:
		goto yyrule74
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate179:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'a':
		goto yystate180
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'b' && c <= 'z':
		goto yystate64
	}

yystate180:
	c = l.next()
	switch {
	default:
		goto yyrule85
	case c == 'r':
		goto yystate181
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'q' || c >= 's' && c <= 'z':
		goto yystate64
	}

yystate181:
	c = l.next()
	switch {
	default:
		goto yyrule75
	case c >= '0' && c <= '9' || c >= 'A' && c <= 'Z' || c == '_' || c >= 'a' && c <= 'z':
		goto yystate64
	}

yystate182:
	c = l.next()
	goto yyrule46

yystate183:
	c = l.next()
	switch {
	default:
		goto yyrule47
	case c == '=':
		goto yystate184
	case c == '|':
		goto yystate185
	}

yystate184:
	c = l.next()
	goto yyrule48

yystate185:
	c = l.next()
	goto yyrule49

yystate186:
	c = l.next()
	goto yyrule50

yystate187:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\u0080' && c <= '¿':
		goto yystate188
	}

yystate188:
	c = l.next()
	goto yyrule86

yystate189:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\u0080' && c <= '¿':
		goto yystate187
	}

yystate190:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c >= '\u0080' && c <= '¿':
		goto yystate189
	}

	goto yystate191 // silence unused label error
yystate191:
	c = l.next()
yystart191:
	switch {
	default:
		goto yystate192 // c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ'
	case c == '"':
		goto yystate193
	case c == '\\':
		goto yystate194
	case c == '\x00':
		goto yystate2
	}

yystate192:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate193
	case c == '\\':
		goto yystate194
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate192
	}

yystate193:
	c = l.next()
	goto yyrule83

yystate194:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '"':
		goto yystate195
	case c == '\\':
		goto yystate194
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate192
	}

yystate195:
	c = l.next()
	switch {
	default:
		goto yyrule83
	case c == '"':
		goto yystate193
	case c == '\\':
		goto yystate194
	case c >= '\x01' && c <= '!' || c >= '#' && c <= '[' || c >= ']' && c <= 'ÿ':
		goto yystate192
	}

	goto yystate196 // silence unused label error
yystate196:
	c = l.next()
yystart196:
	switch {
	default:
		goto yystate197 // c >= '\x01' && c <= '_' || c >= 'a' && c <= 'ÿ'
	case c == '\x00':
		goto yystate2
	case c == '`':
		goto yystate198
	}

yystate197:
	c = l.next()
	switch {
	default:
		goto yyabort
	case c == '`':
		goto yystate198
	case c >= '\x01' && c <= '_' || c >= 'a' && c <= 'ÿ':
		goto yystate197
	}

yystate198:
	c = l.next()
	goto yyrule84

yyrule1: // \0
	{
		l.i0++
		return token.EOF, lval
	}
yyrule2: // [ \t\n\r]+

	goto yystate0
yyrule3: // \/\*([^*]|\*+[^*/])*\*+\/
yyrule4: // \/\/.*
	{
		return token.COMMENT, string(l.val)
	}
yyrule5: // "!"
	{
		return token.NOT, lval
	}
yyrule6: // "!="
	{
		return token.NEQ, lval
	}
yyrule7: // "%"
	{
		return token.REM, lval
	}
yyrule8: // "%="
	{
		return token.REM_ASSIGN, lval
	}
yyrule9: // "&"
	{
		return token.AND, lval
	}
yyrule10: // "&&"
	{
		return token.LAND, lval
	}
yyrule11: // "&="
	{
		return token.AND_ASSIGN, lval
	}
yyrule12: // "&^"
	{
		return token.AND_NOT, lval
	}
yyrule13: // "&^="
	{
		return token.AND_NOT_ASSIGN, lval
	}
yyrule14: // "("
	{
		return token.LPAREN, lval
	}
yyrule15: // ")"
	{
		return token.RPAREN, lval
	}
yyrule16: // "*"
	{
		return token.MUL, lval
	}
yyrule17: // "*="
	{
		return token.MUL_ASSIGN, lval
	}
yyrule18: // "+"
	{
		return token.ADD, lval
	}
yyrule19: // "++"
	{
		return token.INC, lval
	}
yyrule20: // "+="
	{
		return token.ADD_ASSIGN, lval
	}
yyrule21: // ","
	{
		return token.COMMA, lval
	}
yyrule22: // "-"
	{
		return token.SUB, lval
	}
yyrule23: // "--"
	{
		return token.DEC, lval
	}
yyrule24: // "-="
	{
		return token.SUB_ASSIGN, lval
	}
yyrule25: // "."
	{
		return token.PERIOD, lval
	}
yyrule26: // "..."
	{
		return token.ELLIPSIS, lval
	}
yyrule27: // "/"
	{
		return token.QUO, lval
	}
yyrule28: // "/="
	{
		return token.QUO_ASSIGN, lval
	}
yyrule29: // ":"
	{
		return token.COLON, lval
	}
yyrule30: // ":="
	{
		return token.DEFINE, lval
	}
yyrule31: // ";"
	{
		return token.SEMICOLON, lval
	}
yyrule32: // "<"
	{
		return token.LSS, lval
	}
yyrule33: // "<-"
	{
		return token.ARROW, lval
	}
yyrule34: // "<<"
	{
		return token.SHL, lval
	}
yyrule35: // "<<="
	{
		return token.SHL_ASSIGN, lval
	}
yyrule36: // "<="
	{
		return token.LEQ, lval
	}
yyrule37: // "=="
	{
		return token.EQL, lval
	}
yyrule38: // ">"
	{
		return token.GTR, lval
	}
yyrule39: // ">="
	{
		return token.GEQ, lval
	}
yyrule40: // ">>"
	{
		return token.SHR, lval
	}
yyrule41: // ">>="
	{
		return token.SHR_ASSIGN, lval
	}
yyrule42: // "["
	{
		return token.LBRACK, lval
	}
yyrule43: // "]"
	{
		return token.RBRACK, lval
	}
yyrule44: // "^"
	{
		return token.XOR, lval
	}
yyrule45: // "^="
	{
		return token.XOR_ASSIGN, lval
	}
yyrule46: // "{"
	{
		return token.LBRACE, lval
	}
yyrule47: // "|"
	{
		return token.OR, lval
	}
yyrule48: // "|="
	{
		return token.OR_ASSIGN, lval
	}
yyrule49: // "||"
	{
		return token.LOR, lval
	}
yyrule50: // "}"
	{
		return token.RBRACE, lval
	}
yyrule51: // break
	{
		return token.BREAK, lval
	}
yyrule52: // case
	{
		return token.CASE, lval
	}
yyrule53: // chan
	{
		return token.CHAN, lval
	}
yyrule54: // const
	{
		return token.CONST, lval
	}
yyrule55: // continue
	{
		return token.CONTINUE, lval
	}
yyrule56: // default
	{
		return token.DEFAULT, lval
	}
yyrule57: // defer
	{
		return token.DEFER, lval
	}
yyrule58: // else
	{
		return token.ELSE, lval
	}
yyrule59: // fallthrough
	{
		return token.FALLTHROUGH, lval
	}
yyrule60: // for
	{
		return token.FOR, lval
	}
yyrule61: // func
	{
		return token.FUNC, lval
	}
yyrule62: // go
	{
		return token.GO, lval
	}
yyrule63: // goto
	{
		return token.GOTO, lval
	}
yyrule64: // if
	{
		return token.IF, lval
	}
yyrule65: // import
	{
		return token.IMPORT, lval
	}
yyrule66: // interface
	{
		return token.INTERFACE, lval
	}
yyrule67: // map
	{
		return token.MAP, lval
	}
yyrule68: // package
	{
		return token.PACKAGE, lval
	}
yyrule69: // range
	{
		return token.RANGE, lval
	}
yyrule70: // return
	{
		return token.RETURN, lval
	}
yyrule71: // select
	{
		return token.SELECT, lval
	}
yyrule72: // struct
	{
		return token.STRUCT, lval
	}
yyrule73: // switch
	{
		return token.SWITCH, lval
	}
yyrule74: // type
	{
		return token.TYPE, lval
	}
yyrule75: // var
	{
		return token.VAR, lval
	}
yyrule76: // {imaginary_ilit}
	{
		return l.int(true)
	}
yyrule77: // {imaginary_lit}
	{
		return l.float(true)
	}
yyrule78: // {int_lit}
	{
		return l.int(false)
	}
yyrule79: // {float_lit}
	{
		return l.float(false)
	}
yyrule80: // \"
	{
		l.sc = S1
		goto yystate0
	}
yyrule81: // `
	{
		l.sc = S2
		goto yystate0
	}
yyrule82: // '(\\.|[^'])*'
	{
		if tok, lval = l.str(""); tok != token.STRING {
			return
		}
		return token.INT, int32(lval.(string)[0])
	}
yyrule83: // (\\.|[^\"])*\"
	{
		return l.str("\"")
	}
yyrule84: // ([^`]|\n)*`
	{
		return l.str("`")
	}
yyrule85: // [a-zA-Z_][a-zA-Z_0-9]*
	{

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
		return token.IDENT, string(l.src[l.i0-1 : l.i-1])
	}
yyrule86: // {non_ascii}
	{

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
	}
	panic("unreachable")

	goto yyabort // silence unused label error

yyabort: // no lexem recognized
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
