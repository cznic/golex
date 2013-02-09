// Copyright (c) 2013 jnml. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//TODO Investigate "testing" bug(?) @ 4aad2c4e8945f0e047d3c767497385b55b951aa1

package cc

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"testing"
)

func TestScanPeekNext(t *testing.T) {
	const N = 17
	b := make([]byte, N)
	for i := range b {
		b[i] = byte(i + 1)
	}
	f := func(rf func() io.Reader) {
		for n := 0; n <= N; n++ {
			b = b[:n]
			for sz := 0; sz <= len(b)+2; sz++ {
				//r := bytes.NewReader(b)
				r := rf()
				failed := false
				l := newLexer(
					"file",
					r,
					func(file string, line, col int, msg string, args ...interface{}) bool {
						failed = true
						return true
					},
					make([]byte, sz),
				)
				t.Log("----")
				for i := 1; i <= n+1; i++ {
					t.Logf("n %d sz %d i %d", n, sz, i)
					c := l.peek()
					t.Logf("peek == %d", c)
					switch {
					case failed:
						t.Fatal("unexpected")
					case c == 0:
						if g, e := i, n+1; g != e {
							t.Fatalf("got peek == 0 at %d, expceted at %d", g, e)
						}

						c = l.next(c)
						t.Logf("next == %d", c)
						if g, e := c, 0; g != e {
							t.Fatalf("got next == %d, expected %d", g, e)
						}
					case c > 0 && i > n:
						t.Fatalf("got peek > 0 (%d), expected 0", c)
					default:
						if g, e := c, i; g != i {
							t.Fatalf("got peek == %d, expected %d", g, e)
						}

						c = l.next(c)
						t.Logf("next == %d", c)
						switch {
						case failed:
							t.Fatal("unexpected")
						case c == 0:
							if g, e := i, n; g != e {
								t.Fatalf("got next == 0 at %d, expected at %d", g, e)
							}
						default:
							if i > n {
								t.Fatalf("got next > 0 (%d), expected 0", c)
							}
						}
					}
				}
			}
		}
	}
	f(func() io.Reader { return bytes.NewReader(b) })
	f(func() io.Reader { return strings.NewReader(string(b)) })
	f(func() io.Reader { return bufio.NewReader(bytes.NewReader(b)) })

}

func TestLineDirectives(t *testing.T) {
	const (
		m = "main.c"
		f = "foo.c"
		b = "bar.c"
	)
	tab := []struct {
		s, file   string
		line, col int
	}{
		{"",
			m, 1, 1},
		{" ",
			m, 1, 2},
		{"  ",
			m, 1, 3},
		{"\n",
			m, 2, 1},
		{"\n ",
			m, 2, 2},
		{"\n  ",
			m, 2, 3},
		{" \n",
			m, 2, 1},
		{" \n ",
			m, 2, 2},
		{" \n  ",
			m, 2, 3},
		{"\n\n",
			m, 3, 1},
		{"\n\n ",
			m, 3, 2},
		{"\n \n",
			m, 3, 1},
		{"\n \n ",
			m, 3, 2},
		{" \n\n",
			m, 3, 1},
		{" \n\n ",
			m, 3, 2},
		{" \n \n",
			m, 3, 1},
		{" \n \n ",
			m, 3, 2},
		{`# 123456 "foo.c"`,
			f, 123455, 17},
		{`# 123456 "foo.c"` + "\n",
			f, 123456, 1},
		{` # 123456 "foo.c"`,
			m, 1, 18},
		{` # 123456 "foo.c"` + "\n",
			m, 2, 1},
		{`
# 123456 "foo.c"
`,
			f, 123456, 1},
		{`
# 123456 "foo.c"

`,
			f, 123457, 1},
		{`

# 123456 "foo.c"

`,
			f, 123457, 1},
		{`

# 123456 "foo.c"


`,
			f, 123458, 1},
		{`

# 999 "bar.c"

# 123456 "foo.c"

`,
			f, 123457, 1},

		{`#123456"foo.c"` + "\n",
			f, 123456, 1},
		{`#123456"foo.c" ` + "\n",
			f, 123456, 1},
		{`#123456 "foo.c"` + "\n",
			f, 123456, 1},
		{`#123456 "foo.c" ` + "\n",
			f, 123456, 1},
		{`# 123456"foo.c"` + "\n",
			f, 123456, 1},
		{`# 123456"foo.c" ` + "\n",
			f, 123456, 1},
		{`# 123456 "foo.c"` + "\n",
			f, 123456, 1},
		{`# 123456 "foo.c" ` + "\n",
			f, 123456, 1},
		{`# 123456 "foo.c"what so ever` + "\n",
			f, 123456, 1},

		{`#x 1000 "foo.c"` + "\n",
			m, 2, 1},
		{`# x1001 "foo.c"` + "\n",
			m, 2, 1},
		{`# 1002x "foo.c"` + "\n",
			m, 2, 1},
		{`# 1003 x"foo.c"` + "\n",
			m, 2, 1},
		{`# 1004 foo.c"` + "\n",
			m, 2, 1},
		{`# 1005 "foo.c` + "\n",
			m, 2, 1},
		{`# 1006 "foo.c`,
			m, 1, 14},
		{`# 1007 "` + "\n",
			m, 2, 1},
		{`# 1008 "`,
			m, 1, 9},
		{`# 1009 ` + "\n",
			m, 2, 1},
		{`# 1010 `,
			m, 1, 8},
		{`# 1011` + "\n",
			m, 2, 1},
		{`# 1012`,
			m, 1, 7},
		{`# ` + "\n",
			m, 2, 1},
		{`# `,
			m, 1, 3},
		{`#` + "\n",
			m, 2, 1},
		{`#`,
			m, 1, 2},
	}
	for i, test := range tab {
		r := strings.NewReader(test.s)
		l := newLexer(
			m,
			r,
			func(file string, line, col int, msg string, args ...interface{}) bool {
				t.Fatal(i, "unexpected")
				return false
			},
			nil,
		)

		for j := 0; ; j++ {
			c := l.scan()
			if c == 0 {
				break
			}
			if j > 100 {
				t.Fatal(i, j, "missed EOF")
			}
		}

		if l.file != test.file || l.line != test.line || l.col != test.col {
			t.Fatalf(
				`%d: (%q) got "%s.%d:%d", exp "%s.%d:%d"`,
				i, test.s,
				l.file, l.line, l.col,
				test.file, test.line, test.col,
			)
		}
	}
}
