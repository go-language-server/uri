// Copyright 2026 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"testing"
	"unicode/utf8"
)

func FuzzParseNoPanic(f *testing.F) {
	for _, seed := range []string{
		"file:///home/user/x.go",
		"https://host/p?name=ferret#f",
		"file://shares/files/c%23/p.cs",
		"file://some/%A0.txt",
		"foo:api/files/test",
		"file:////shares/files/p.cs",
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, input string) {
		_, _ = Parse(input)
	})
}

func FuzzParseRoundTripCanonical(f *testing.F) {
	for _, seed := range []string{
		"file:///home/user/x.go",
		"https://host/p?name=ferret#f",
		"file://shares/files/c%23/p.cs",
		"foo:api/files/test",
		"untitled:untitled-1",
	} {
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, input string) {
		if !utf8.ValidString(input) {
			return
		}
		u, err := Parse(input)
		if err != nil {
			return
		}
		v, err := Parse(u.String())
		if err != nil {
			t.Fatalf("Parse(canonical %q) error = %v", u.String(), err)
		}
		if v.String() != u.String() {
			t.Fatalf("canonical round trip = %q, want %q", v.String(), u.String())
		}
	})
}
