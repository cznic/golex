// Copyright (c) 2013 jnml. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//TODO Investigate "testing" bug(?) @ 4aad2c4e8945f0e047d3c767497385b55b951aa1

package cc

import (
	"bytes"
	"testing"
)

func TestScanPeekNext(t *testing.T) {
	const N = 37
	b := make([]byte, N)
	for i := range b {
		b[i] = byte(i + 1)
	}
	for n := 0; n <= N; n++ {
		b = b[:n]
		for sz := 0; sz <= len(b)+2; sz++ {
			r := bytes.NewReader(b)
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
			//t.Log("----")
			for i := 1; i <= n+1; i++ {
				//t.Logf("n %d sz %d i %d", n, sz, i)
				c := l.peek()
				//t.Logf("peek == %d", c)
				switch {
				case failed:
					t.Fatal()
				case c < 0:
					if g, e := i, n+1; g != e {
						t.Fatalf("got peek == -1 at %d, expceted at %d", g, e)
					}

					c = l.next(c)
					//t.Logf("next == %d", c)
					if g, e := c, -1; g != e {
						t.Fatalf("got next == %d, expected %d", g, e)
					}
				case c >= 0 && i > n:
					t.Fatalf("got peek >= 0 (%d), expected -1", c)
				default:
					if g, e := c, i; g != i {
						t.Fatalf("got peek == %d, expected %d", g, e)
					}

					c = l.next(c)
					//t.Logf("next == %d", c)
					switch {
					case failed:
						t.Fatal()
					case c < 0:
						if g, e := i, n; g != e {
							t.Fatalf("got next == -1 at %d, expected at %d", g, e)
						}
					default:
						if i >= n {
							t.Fatalf("got next >= 0 (%d), expected -1", c)
						}
					}
				}
			}
		}
	}

}
