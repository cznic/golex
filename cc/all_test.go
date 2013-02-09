// Copyright (c) 2013 jnml. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc

import (
	"bytes"
	"testing"
)

func TestScanPeekNext(t *testing.T) {
	const N = 1
	b := make([]byte, N)
	for i := range b {
		b[i] = byte(i + 1)
	}
	for n := 0; n < len(b); n++ {
		b = b[:n]
		for sz := 0; sz <= len(b); sz++ {
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
			for i := 1; i <= len(b)+1; i++ {
				_ = l.peek()
				switch {
				case failed:
					t.Fatal("TODO36")
				default:
					t.Fatal("TODO39")
				}
			}
		}
	}

}
