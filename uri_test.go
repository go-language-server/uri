// Copyright 2026 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseCoreComponents(t *testing.T) {
	tests := map[string]struct {
		input     string
		want      Components
		wantText  string
		wantRaw   string
		wantError error
	}{
		"success: http path gets leading slash": {
			input: "http:/api/files/test.me?t=1234",
			want: Components{
				Scheme: "http",
				Path:   "/api/files/test.me",
				Query:  "t=1234",
			},
			wantText: "http:/api/files/test.me?t%3D1234",
			wantRaw:  "http:/api/files/test.me?t=1234",
		},
		"success: http authority lowercases host for canonical components": {
			input: "http://API/files/test.me?t=1234",
			want: Components{
				Scheme:    "http",
				Authority: "api",
				Path:      "/files/test.me",
				Query:     "t=1234",
			},
			wantText: "http://api/files/test.me?t%3D1234",
			wantRaw:  "http://api/files/test.me?t=1234",
		},
		"success: file unicode and escaped path": {
			input: "file:///c:/Source/Z%C3%BCrich%20or%20Zurich%20(%CB%88zj%CA%8A%C9%99r%C9%AAk,/Code/resources/app/plugins/c%23/plugin.json",
			want: Components{
				Scheme: "file",
				Path:   "/c:/Source/Zürich or Zurich (ˈzjʊərɪk,/Code/resources/app/plugins/c#/plugin.json",
			},
			wantText: "file:///c%3A/Source/Z%C3%BCrich%20or%20Zurich%20%28%CB%88zj%CA%8A%C9%99r%C9%AAk%2C/Code/resources/app/plugins/c%23/plugin.json",
			wantRaw:  "file:///c:/Source/Zürich or Zurich (ˈzjʊərɪk,/Code/resources/app/plugins/c%23/plugin.json",
		},
		"success: invalid percent escape is preserved then percent encoded": {
			input: "file://some/%A0.txt",
			want: Components{
				Scheme:    "file",
				Authority: "some",
				Path:      "/%A0.txt",
			},
			wantText: "file://some/%25A0.txt",
			wantRaw:  "file://some/%A0.txt",
		},
		"success: invalid utf8 run recurses by triplet": {
			input: "file:///x/%C3%28.go",
			want: Components{
				Scheme: "file",
				Path:   "/x/%C3(.go",
			},
			wantText: "file:///x/%25C3%28.go",
			wantRaw:  "file:///x/%C3(.go",
		},
		"success: non strict missing scheme falls back to file": {
			input: "foo/bar",
			want: Components{
				Scheme: "file",
				Path:   "/foo/bar",
			},
			wantText: "file:///foo/bar",
			wantRaw:  "file:///foo/bar",
		},
		"success: non-rfc scheme accepted like vscode": {
			input: "f3ile:#d",
			want: Components{
				Scheme:   "f3ile",
				Fragment: "d",
			},
			wantText: "f3ile:#d",
			wantRaw:  "f3ile:#d",
		},
		"success: generic scheme keeps relative path": {
			input: "foo:api/files/test",
			want: Components{
				Scheme: "foo",
				Path:   "api/files/test",
			},
			wantText: "foo:api/files/test",
			wantRaw:  "foo:api/files/test",
		},
		"success: query and fragment are opaque components": {
			input: "https://host/p?b=2&a=1&a=0#frag ment",
			want: Components{
				Scheme:    "https",
				Authority: "host",
				Path:      "/p",
				Query:     "b=2&a=1&a=0",
				Fragment:  "frag ment",
			},
			wantText: "https://host/p?b%3D2%26a%3D1%26a%3D0#frag%20ment",
			wantRaw:  "https://host/p?b=2&a=1&a=0#frag ment",
		},
		"success: empty query delimiter is dropped": {
			input:    "foo:?",
			want:     Components{Scheme: "foo"},
			wantText: "foo:",
			wantRaw:  "foo:",
		},
		"success: empty fragment delimiter is dropped": {
			input:    "foo:#",
			want:     Components{Scheme: "foo"},
			wantText: "foo:",
			wantRaw:  "foo:",
		},
		"success: trailing empty fragment delimiter after query is dropped": {
			input:    "foo:?x#",
			want:     Components{Scheme: "foo", Query: "x"},
			wantText: "foo:?x",
			wantRaw:  "foo:?x",
		},
		"success: empty query and fragment delimiters on https are dropped": {
			input: "https://host/p?#",
			want: Components{
				Scheme:    "https",
				Authority: "host",
				Path:      "/p",
			},
			wantText: "https://host/p",
			wantRaw:  "https://host/p",
		},
		"success: uppercase file authority canonicalizes component view": {
			input: "file://SERVER/Share/X.go",
			want: Components{
				Scheme:    "file",
				Authority: "server",
				Path:      "/Share/X.go",
			},
			wantText: "file://server/Share/X.go",
			wantRaw:  "file://server/Share/X.go",
		},
		"error: file path without authority cannot begin double slash": {
			input:     "file:////shares/files/p.cs",
			wantError: ErrPathAuthority,
		},
		"error: authority requires slash path": {
			input:     "foo://hostpath",
			want:      Components{Scheme: "foo", Authority: "hostpath"},
			wantText:  "foo://hostpath",
			wantRaw:   "foo://hostpath",
			wantError: nil,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := Parse(tt.input)
			if tt.wantError != nil {
				if !errors.Is(err, tt.wantError) {
					t.Fatalf("Parse() error = %v, want %v", err, tt.wantError)
				}
				return
			}
			if err != nil {
				t.Fatalf("Parse() error = %v", err)
			}
			if diff := cmp.Diff(tt.want, got.Components()); diff != "" {
				t.Fatalf("Components() mismatch (-want +got):\n%s", diff)
			}
			if got.String() != tt.wantText {
				t.Fatalf("String() = %q, want %q", got.String(), tt.wantText)
			}
			if got.StringNoEncoding() != tt.wantRaw {
				t.Fatalf("StringNoEncoding() = %q, want %q", got.StringNoEncoding(), tt.wantRaw)
			}
		})
	}
}

func TestParseStrict(t *testing.T) {
	tests := map[string]struct {
		input     string
		wantError error
	}{
		"success: scheme present":       {input: "file:///x.go"},
		"error: missing scheme":         {input: "x.go", wantError: ErrMissingScheme},
		"error: invalid scheme":         {input: "fäil:x", wantError: ErrInvalidScheme},
		"error: percent encoded scheme": {input: "f%6Fo:path", wantError: ErrInvalidScheme},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			_, err := ParseStrict(tt.input)
			if tt.wantError == nil {
				if err != nil {
					t.Fatalf("ParseStrict() error = %v", err)
				}
				return
			}
			if !errors.Is(err, tt.wantError) {
				t.Fatalf("ParseStrict() error = %v, want %v", err, tt.wantError)
			}
		})
	}
}

func TestFromCore(t *testing.T) {
	tests := map[string]struct {
		input Components
		want  string
	}{
		"success: empty scheme becomes file quirk": {
			input: Components{Path: "x.go"},
			want:  "file:///x.go",
		},
		"success: query is encoded as opaque component": {
			input: Components{Scheme: "http", Authority: "a-test-site.com", Path: "/", Query: "test=true"},
			want:  "http://a-test-site.com/?test%3Dtrue",
		},
		"success: authority userinfo preserves user case and lowercases host": {
			input: Components{Scheme: "http", Authority: "Foo:bAr@WWW.MSFT.com:8080", Path: "/my/path"},
			want:  "http://Foo:bAr@www.msft.com:8080/my/path",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := From(tt.input)
			if err != nil {
				t.Fatalf("From() error = %v", err)
			}
			if got.String() != tt.want {
				t.Fatalf("From().String() = %q, want %q", got.String(), tt.want)
			}
		})
	}
}

func TestMustParseAndComparability(t *testing.T) {
	tests := map[string]struct {
		input string
	}{
		"success: comparable map key": {input: "file:///x.go"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			u := MustParse(tt.input)
			m := map[URI]string{u: "ok"}
			if got := m[MustParse(tt.input)]; got != "ok" {
				t.Fatalf("map lookup = %q, want ok", got)
			}
			if u.IsZero() {
				t.Fatal("parsed URI reports IsZero")
			}
			if !URI("").IsZero() {
				t.Fatal("zero URI does not report IsZero")
			}
		})
	}
}

func TestURIStringConversionCompatibility(t *testing.T) {
	t.Parallel()

	const text = "file:///x.go"
	u := URI(text)
	if string(u) != text {
		t.Fatalf("string(URI(%q)) = %q, want %q", text, string(u), text)
	}
	if u.String() != text {
		t.Fatalf("URI.String() = %q, want %q", u.String(), text)
	}
	if !u.IsFile() {
		t.Fatal("URI converted from file string does not report IsFile")
	}
	if u.Scheme() != "file" {
		t.Fatalf("URI.Scheme() = %q, want file", u.Scheme())
	}
	if u.Path() != "/x.go" {
		t.Fatalf("URI.Path() = %q, want /x.go", u.Path())
	}
}

func TestCanonicalEquivalentURIsCompareEqual(t *testing.T) {
	tests := map[string][]string{
		"success: drive case and escaped colon share canonical identity": {
			"file:///C:/test",
			"file:///c:/test",
			"file:///c%3A/test",
		},
		"success: authority host case shares canonical identity": {
			"file://SERVER/Share/X.go",
			"file://server/Share/X.go",
		},
	}
	for name, inputs := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			base := MustParse(inputs[0])
			for _, input := range inputs[1:] {
				got := MustParse(input)
				if got.String() != base.String() {
					t.Fatalf("String(%q) = %q, want canonical %q", input, got.String(), base.String())
				}
				if got != base {
					t.Fatalf("URI(%q) = %#v, want equality with %#v from %q", input, got, base, inputs[0])
				}
			}
		})
	}
}

func TestEncodeComponentFastCore(t *testing.T) {
	tests := map[string]struct {
		input       string
		isPath      bool
		isAuthority bool
		want        string
	}{
		"success: unreserved pass": {
			input: "azAZ09-._~",
			want:  "azAZ09-._~",
		},
		"success: slash passes only in path": {
			input:  "a/b/c",
			isPath: true,
			want:   "a/b/c",
		},
		"success: slash escapes outside path": {
			input: "a/b/c",
			want:  "a%2Fb%2Fc",
		},
		"success: drive colon escapes in path": {
			input:  "/c:/win/path",
			isPath: true,
			want:   "/c%3A/win/path",
		},
		"success: authority keeps colon and brackets": {
			input:       "[::1]:8080",
			isAuthority: true,
			want:        "[::1]:8080",
		},
		"success: authority escapes slash": {
			input:       "server/share",
			isAuthority: true,
			want:        "server%2Fshare",
		},
		"success: unicode utf8 uppercase hex": {
			input:  "pröjects/東京",
			isPath: true,
			want:   "pr%C3%B6jects/%E6%9D%B1%E4%BA%AC",
		},
		"success: query delimiters escape": {
			input: "LinkId=518008&foö&ké¥=üü",
			want:  "LinkId%3D518008%26fo%C3%B6%26k%C3%A9%C2%A5%3D%C3%BC%C3%BC",
		},
		"success: invalid utf8 byte escapes raw": {
			input: string([]byte{0xff}),
			want:  "%FF",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := encodeComponentFast(tt.input, tt.isPath, tt.isAuthority); got != tt.want {
				t.Fatalf("encodeComponentFast(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestEncodeComponentMinimalCore(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"success: escapes hash":                {input: "c#", want: "c%23"},
		"success: escapes question":            {input: "q?x", want: "q%3Fx"},
		"success: leaves percent triplet raw":  {input: "c%23", want: "c%23"},
		"success: leaves unicode raw":          {input: "pröjects/東京", want: "pröjects/東京"},
		"success: leaves query delimiters raw": {input: "LinkId=518008&foö", want: "LinkId=518008&foö"},
		"success: leaves colon and slash raw":  {input: "/c:/win/path", want: "/c:/win/path"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := encodeComponentMinimal(tt.input); got != tt.want {
				t.Fatalf("encodeComponentMinimal(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestParseCanonicalFileFastPathAllocs(t *testing.T) {
	tests := map[string]struct {
		input string
	}{
		"success: clean file URI": {input: "file:///home/user/project/main.go"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			allocs := testing.AllocsPerRun(1000, func() {
				u, err := Parse(tt.input)
				if err != nil {
					t.Fatalf("Parse() error = %v", err)
				}
				if u.String() != tt.input {
					t.Fatalf("String() = %q, want %q", u.String(), tt.input)
				}
			})
			if allocs != 0 {
				t.Fatalf("Parse(%q) allocs = %v, want 0", tt.input, allocs)
			}
		})
	}
}
