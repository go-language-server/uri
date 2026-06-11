// Copyright 2026 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"net/url"
	"strings"
	"testing"

	_ "embed"
)

//go:embed testdata/corpus/uri_bench.tsv
var benchmarkCorpusTSV string

var (
	benchmarkURISink    URI
	benchmarkStringSink string
	benchmarkURLSink    *url.URL
	benchmarkIntSink    int
)

type benchmarkCase struct {
	name string
	text string
}

func BenchmarkParse(b *testing.B) {
	tests := loadBenchmarkCorpus(b)
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			var got URI
			for b.Loop() {
				u, err := Parse(tt.text)
				if err != nil {
					b.Fatal(err)
				}
				got = u
			}
			benchmarkURISink = got
		})
	}
}

func BenchmarkParseNetURLBaseline(b *testing.B) {
	tests := loadBenchmarkCorpus(b)
	for _, tt := range tests {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			var got *url.URL
			for b.Loop() {
				u, err := url.Parse(tt.text)
				if err != nil {
					b.Fatal(err)
				}
				got = u
			}
			benchmarkURLSink = got
		})
	}
}

func BenchmarkFsPath(b *testing.B) {
	tests := map[string]struct {
		uri      string
		platform Platform
	}{
		"file-posix-clean": {
			uri:      "file:///home/user/project/main.go",
			platform: PlatformPOSIX,
		},
		"file-gomodcache-at": {
			uri:      "file:///Users/me/go/pkg/mod/example.com/mod@v1.2.3/file.go",
			platform: PlatformPOSIX,
		},
		"file-windows-drive": {
			uri:      "file:///C:/Users/me/project/main.go",
			platform: PlatformWindows,
		},
		"file-unc-windows": {
			uri:      "file://server/share/project/main.go",
			platform: PlatformWindows,
		},
	}
	for name, tt := range tests {
		b.Run(name, func(b *testing.B) {
			u := MustParse(tt.uri)
			b.ReportAllocs()
			var got string
			for b.Loop() {
				got = FsPathFor(u, tt.platform, false)
			}
			benchmarkStringSink = got
		})
	}
}

func BenchmarkString(b *testing.B) {
	u := MustParse("file:///home/user/project/main.go")
	b.ReportAllocs()
	var got string
	for b.Loop() {
		got = u.String()
	}
	benchmarkStringSink = got
}

func BenchmarkFile(b *testing.B) {
	tests := map[string]struct {
		path     string
		platform Platform
	}{
		"posix-clean-absolute": {
			path:     "/home/user/project/main.go",
			platform: PlatformPOSIX,
		},
		"posix-gomodcache-at": {
			path:     "/Users/me/go/pkg/mod/example.com/mod@v1.2.3/file.go",
			platform: PlatformPOSIX,
		},
		"windows-drive": {
			path:     `C:\Users\me\project\main.go`,
			platform: PlatformWindows,
		},
		"windows-unc": {
			path:     `\\server\share\project\main.go`,
			platform: PlatformWindows,
		},
	}
	for name, tt := range tests {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			var got URI
			for b.Loop() {
				got = FileFor(tt.platform, tt.path)
			}
			benchmarkURISink = got
		})
	}
}

func BenchmarkComponents(b *testing.B) {
	tests := map[string]string{
		"file-posix-clean":        "file:///home/user/project/main.go",
		"file-unicode-escaped":    "file:///home/user/Z%C3%BCrich/c%23/plugin.json",
		"https-query-fragment":    "https://host/p?name=ferret#frag",
		"authority-with-userinfo": "https://User:Pass@host:8443/path?q#f",
	}
	for name, text := range tests {
		b.Run(name, func(b *testing.B) {
			u := MustParse(text)
			b.ReportAllocs()
			var got int
			for b.Loop() {
				c := u.Components()
				got += len(c.Scheme) + len(c.Authority) + len(c.Path) + len(c.Query) + len(c.Fragment)
			}
			benchmarkIntSink = got
		})
	}
}

func BenchmarkPercentDecodeMalformedRun(b *testing.B) {
	tests := map[string]string{
		"invalid-100":         strings.Repeat("%FF", 100),
		"invalid-1000":        strings.Repeat("%FF", 1_000),
		"invalid-10000":       strings.Repeat("%FF", 10_000),
		"invalid-10000-ascii": strings.Repeat("%FF", 10_000) + "%41",
	}
	for name, input := range tests {
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			var got string
			for b.Loop() {
				got = percentDecode(input)
			}
			benchmarkStringSink = got
		})
	}
}

func BenchmarkMapKeyURI(b *testing.B) {
	uris := make([]URI, 10_000)
	for i := range uris {
		uris[i] = MustParse("file:///home/user/project/" + benchDecimal(i) + ".go")
	}
	b.ReportAllocs()
	var got int
	for b.Loop() {
		m := make(map[URI]int, len(uris))
		for i, u := range uris {
			m[u] = i
		}
		for _, u := range uris {
			got += m[u]
		}
	}
	benchmarkIntSink = got
}

func BenchmarkMapKeyNetURLStringBaseline(b *testing.B) {
	urls := make([]*url.URL, 10_000)
	for i := range urls {
		u, err := url.Parse("file:///home/user/project/" + benchDecimal(i) + ".go")
		if err != nil {
			b.Fatal(err)
		}
		urls[i] = u
	}
	b.ReportAllocs()
	var got int
	for b.Loop() {
		m := make(map[string]int, len(urls))
		for i, u := range urls {
			m[u.String()] = i
		}
		for _, u := range urls {
			got += m[u.String()]
		}
	}
	benchmarkIntSink = got
}

func loadBenchmarkCorpus(b *testing.B) []benchmarkCase {
	b.Helper()
	var tests []benchmarkCase
	for lineNumber, line := range strings.Split(benchmarkCorpusTSV, "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		name, text, ok := strings.Cut(line, "\t")
		if !ok || name == "" || text == "" {
			b.Fatalf("invalid benchmark corpus line %d: %q", lineNumber+1, line)
		}
		tests = append(tests, benchmarkCase{name: name, text: text})
	}
	if len(tests) == 0 {
		b.Fatal("empty benchmark corpus")
	}
	return tests
}

func benchDecimal(n int) string {
	var buf [20]byte
	i := len(buf)
	for {
		i--
		buf[i] = byte('0' + n%10)
		n /= 10
		if n == 0 {
			return string(buf[i:])
		}
	}
}
