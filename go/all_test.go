// Copyright (c) 2013 Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scanner

import (
	"fmt"
	goscanner "go/scanner"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
)

func dbg(s string, va ...interface{}) {
	_, fn, fl, _ := runtime.Caller(1)
	fmt.Printf("%s:%d: ", path.Base(fn), fl)
	fmt.Printf(s, va...)
	fmt.Println()
}

type visitor struct {
	count    int
	tokCount int
	size     int64
	t        *testing.T
}

func (v *visitor) visitFile(path string, f os.FileInfo) {
	t := v.t
	ok, err := filepath.Match("*.go", filepath.Base(path))
	if err != nil {
		t.Fatal(err)
	}

	if !ok {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()
	src, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("----\n\n%s", path)
	var sc goscanner.Scanner
	fs := token.NewFileSet()
	base := fs.Base()
	fl := fs.AddFile(path, base, int(f.Size()))
	sc.Init(fl, src, nil, goscanner.ScanComments)

	l := newLexer(src)

	i := 1
loop:
	for {
		pos, tok, lit := sc.Scan()
		if tok == token.SEMICOLON && lit == "\n" { // injected
			continue
		}

		tok2, lit2 := l.Scan()
		pos2 := token.Pos(base + l.i0 - 1)
		p, p2 := fs.Position(pos), fs.Position(pos2)

		if g, e := tok2, tok; g != e {
			t.Fatalf("%d.%d %s(%d) %s(%d) %s %s", v.count, i, g, int(g), e, int(e), p2, p)
		}

		if g, e := p2, p; g != e {
			t.Fatal(v.count, i, tok, g, e)
		}

		if lit2 == nil {
			lit2 = tok2.String()
		}
		switch tok2 {
		case token.EOF:
			break loop
		case token.LPAREN:
			// nop
		default:
			if g, e := lit2.(string), lit; g != e {
				t.Fatalf("%d.%d %s %q %q %s %s", v.count, i, tok, g, e, p2, p)
			}
		}

		v.tokCount++
		i++
		t.Logf("%v %v %q %q", tok, p, lit2, lit)
	}

	v.count++
	v.size += f.Size()
}

func Test(t *testing.T) {
	v := &visitor{t: t}
	if err := filepath.Walk(runtime.GOROOT()+"/src", func(pth string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			v.visitFile(pth, info)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}

	t.Logf("%d .go files, %d bytes, %d tokens\n", v.count, v.size, v.tokCount)
}
