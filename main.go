// Copyright (c) 2011 CZ.NIC z.s.p.o. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// blame: jnml, labs.nic.cz

package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/cznic/lex"
	"io"
	"os"
)

const (
	OFILE = "lex.yy.go"
)

var (
	stdin  = bufio.NewReader(os.Stdin)
	stdout = bufio.NewWriter(os.Stdout)
	stderr = bufio.NewWriter(os.Stderr)
)

type renderer interface {
	render(srcname string, l *lex.L)
}

type writer interface {
	io.Writer
	wprintf(s string, args ...interface{}) (n int, err error)
}

type noRender struct {
	w io.Writer
}

func (r *noRender) Write(p []byte) (n int, err error) {
	return r.w.Write(p)
}

func (r *noRender) wprintf(s string, args ...interface{}) (n int, err error) {
	n, err = io.WriteString(r.w, fmt.Sprintf(s, args...))
	if err != nil {
		panic(err)
	}

	return
}

func q(c uint32) string {
	switch c {
	default:
		s := fmt.Sprintf("%q", string(c))
		return "'" + s[1:len(s)-1] + "'"
	case '\'':
		return "'\\''"
	case '"':
		return "'\"'"
	}

	panic("unreachable")
}

func main() {

	defer func() {
		if e := recover(); e != nil {
			fmt.Fprintf(stderr, "%s: %s\n", os.Args[0], e.(error))
			stderr.Flush()
			os.Exit(1)
		}
	}()

	oflag := ""
	var dfaflag, hflag, tflag, vflag, nodfaopt, bits32 bool

	flag.BoolVar(&dfaflag, "DFA", false, "write DFA on stdout and quit")
	flag.BoolVar(&hflag, "h", false, "show help and exit")
	flag.StringVar(&oflag, "o", OFILE, "lexer output")
	flag.BoolVar(&tflag, "t", false, "write scanner on stdout instead of "+OFILE)
	flag.BoolVar(&vflag, "v", false, "write summary of scanner statistics to stderr")
	flag.BoolVar(&nodfaopt, "nodfaopt", false, "disable DFA optimization - don't use this for production code")
	//flag.BoolVar(&bits32, "32bit", false, "assume unicode rune lexer (partially implemented)")
	flag.Parse()
	if hflag || flag.NArg() > 1 {
		flag.Usage()
		fmt.Fprintf(stderr, "\n%s [-o out_name] [other_options] [in_name]\n", os.Args[0])
		fmt.Fprintln(stderr, "  If no in_name is given then read from stdin.")
		stderr.Flush()
		os.Exit(1)
	}

	var (
		lfile  *bufio.Reader // source .l
		gofile *bufio.Writer // dest .go
	)

	lname := flag.Arg(0)
	if lname == "" {
		lfile = stdin
	} else {
		l, err := os.Open(lname)
		if err != nil {
			panic(err)
		}

		defer l.Close()
		lfile = bufio.NewReader(l)
	}

	l, err := lex.NewL(lname, lfile, nodfaopt, bits32)
	if err != nil {
		panic(err)
	}

	if dfaflag {
		fmt.Println(l.DfaString())
		panic(nil)
	}

	if tflag {
		gofile = stdout
	} else {
		if oflag == "" {
			oflag = OFILE
		}
		g, err := os.Create(oflag)
		if err != nil {
			panic(err)
		}

		defer g.Close()
		gofile = bufio.NewWriter(g)
	}
	defer gofile.Flush()
	renderGo{noRender{gofile}, map[int]bool{}}.render(lname, l)

	if vflag {
		fmt.Fprintln(os.Stderr, l.String())
	}
}
