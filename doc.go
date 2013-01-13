// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz

// Golex is a lex/flex like (not fully POSIX lex compatible) utility.
// It renders .l formated data (http://flex.sourceforge.net/manual/Format.html#Format) to Go source code.
// The .l data can come from a file named in a command line argument.
// If no non-opt args are given, golex reads stdin.
//
// Options:
//	-DFA            print the DFA to stdout and quit
//	-nodfaopt       disable DFA optimization - don't use this for production code
//	-o fname        write to file `fname`, default is `lex.yy.go`
//	-t              write to stdout
//	-v              write some scanner statistics to stderr
//	-32bit          assume unicode rune lexer (partially implemented, disabled)
//
// Golex output is not pretty formated, use the standard tool:
//
//	    $ golex -t file.l | gofmt > file.go
//	or
//	    $ golex -o file.go file.l && gofmt -w file.go
//
// Missing/differing functionality of the current renderer (compared to flex):
//	- No runtime tokenizer package/environment
//	  (but the freedom to have/write any fitting one's specific task(s)).
//	- The generated FSM picks the rules in the order of their appearance in the .l source,
//	  but "flex picks the rule that matches the most text".
//	- No back step. For patterns `a` and `abc` the text `abxy` doesn't match anything.
//	  Flex would match pattern `a` after seeing `x` and its next scan will start at `b`.
//	  The lex.L FSM has all the data to properly render this functionality,
//	  it's only this tool that doesn't implement it.
//	- And probably more.
// Further limitations on the .l source are listed in the cznic/lex package godocs.
//
// A simple golex program example (make example1 && ./example1):
//
//	%{
//	package main
//	
//	import (
//	    "bufio"
//	    "fmt"
//	    "log"
//	    "os"
//	)
//	
//	var (
//	    src      = bufio.NewReader(os.Stdin)
//	    buf      []byte
//	    current  byte
//	)
//	
//	func getc() byte {
//	    if current != 0 {
//	        buf = append(buf, current)
//	    }
//	    current = 0
//	    if b, err := src.ReadByte(); err == nil {
//	        current = b
//	    }
//	    return current
//	}
//	
//	//    %yyc is a "macro" to access the "current" character.
//	//
//	//    %yyn is a "macro" to move to the "next" character.
//	//
//	//    %yyb is a "macro" to return the begining-of-line status (a bool typed value).
//	//        It is used for patterns like `^re`.
//	//        Example: %yyb prev == 0 || prev == '\n'
//	//
//	//    %yyt is a "macro" to return the top/current start condition (an int typed value).
//	//        It is used when there are patterns with conditions like `<cond>re`.
//	//        Example: %yyt startCond
//
//	func main() { // This left brace is closed by *1
//	    c := getc() // init
//	%}
//	
//	%yyc c
//	%yyn c = getc()
//	
//	D   [0-9]+
//	
//	%%
//	    buf = buf[:0]   // Code before the first rule is executed before every scan cycle (state 0 action)
//	
//	[ \t\n\r]+          // Ignore whitespace
//	
//	{D}                 fmt.Printf("int %q\n", buf)
//	
//	{D}\.{D}?|\.{D}     fmt.Printf("float %q\n", buf)
//	
//	\0                  return // Exit on EOF or any other error
//	
//	.                   fmt.Printf("%q\n", buf) // Printout any other unrecognized stuff
//	
//	%%
//	    // The rendered scanner enters top of the user code section when 
//	    // lexem recongition fails. In this example it should never happen.
//	    log.Fatal("scanner internal error")
//	
//	} // *1 this right brace
package documentation
