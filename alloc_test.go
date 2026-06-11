// Copyright 2026 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import "testing"

func TestAllocs(t *testing.T) {
	tests := map[string]struct {
		maxAllocs float64
		fn        func(t *testing.T)
	}{
		"Parse clean canonical file is zero alloc": {
			maxAllocs: 0,
			fn: func(t *testing.T) {
				u, err := Parse("file:///home/user/x.go")
				if err != nil {
					t.Fatal(err)
				}
				if u.String() != "file:///home/user/x.go" {
					t.Fatalf("unexpected string %q", u.String())
				}
			},
		},
		"String is zero alloc": {
			maxAllocs: 0,
			fn: func(t *testing.T) {
				u := MustParse("file:///home/user/x.go")
				if u.String() == "" {
					t.Fatal("empty string")
				}
			},
		},
		"FsPath clean POSIX file is zero alloc": {
			maxAllocs: 0,
			fn: func(t *testing.T) {
				u := MustParse("file:///home/user/x.go")
				if got := FsPathFor(u, PlatformPOSIX, false); got != "/home/user/x.go" {
					t.Fatalf("FsPathFor() = %q", got)
				}
			},
		},
		"File clean absolute is at most one alloc": {
			maxAllocs: 1,
			fn: func(t *testing.T) {
				u := FileFor(PlatformPOSIX, "/abs/clean/path.go")
				if u.String() != "file:///abs/clean/path.go" {
					t.Fatalf("FileFor() = %q", u.String())
				}
			},
		},
		"Parse https opaque query is at most two allocs": {
			maxAllocs: 2,
			fn: func(t *testing.T) {
				u, err := Parse("https://host/p?name=ferret#f")
				if err != nil {
					t.Fatal(err)
				}
				if u.String() != "https://host/p?name%3Dferret#f" {
					t.Fatalf("Parse() = %q", u.String())
				}
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			allocs := testing.AllocsPerRun(1000, func() { tt.fn(t) })
			if allocs > tt.maxAllocs {
				t.Fatalf("allocs = %v, want <= %v", allocs, tt.maxAllocs)
			}
		})
	}
}
