// Copyright 2026 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import "testing"

func TestPosixNormalize(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"success: relative single":        {input: "a", want: "a"},
		"success: absolute single":        {input: "/a", want: "/a"},
		"success: trailing slash":         {input: "a/", want: "a/"},
		"success: duplicate slashes":      {input: "/a/foo/bar///x", want: "/a/foo/bar/x"},
		"success: absolute trailing":      {input: "/a/foo/bar/x/", want: "/a/foo/bar/x/"},
		"success: relative trailing many": {input: "a/foo/bar/x//", want: "a/foo/bar/x/"},
		"success: leading double slash":   {input: "//a/foo/bar/x//", want: "/a/foo/bar/x/"},
		"success: current dir removed":    {input: "a/./b", want: "a/b"},
		"success: parent removed":         {input: "a/n/../b", want: "a/b"},
		"success: parent keeps trailing":  {input: "a/n/../", want: "a/"},
		"success: relative parent root":   {input: "a/..", want: "."},
		"success: absolute above root":    {input: "/a/n/../../..", want: "/"},
		"success: relative above root":    {input: "..", want: ".."},
		"success: empty":                  {input: "", want: "."},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := posixNormalize(tt.input); got != tt.want {
				t.Fatalf("posixNormalize(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestJoinPath(t *testing.T) {
	tests := map[string]struct {
		uri      string
		segments []string
		want     string
	}{
		"success: append segment":        {uri: "foo://a/foo/bar", segments: []string{"x"}, want: "foo://a/foo/bar/x"},
		"success: trim base trailing":    {uri: "foo://a/foo/bar/", segments: []string{"x"}, want: "foo://a/foo/bar/x"},
		"success: absolute segment":      {uri: "foo://a/foo/bar/", segments: []string{"/x"}, want: "foo://a/foo/bar/x"},
		"success: preserve result slash": {uri: "foo://a/foo/bar/", segments: []string{"x/"}, want: "foo://a/foo/bar/x/"},
		"success: multiple segments":     {uri: "foo://a/foo/bar/", segments: []string{"x/", "/y"}, want: "foo://a/foo/bar/x/y"},
		"success: current dir segment":   {uri: "foo://a/foo/bar/", segments: []string{".", "/y"}, want: "foo://a/foo/bar/y"},
		"success: parent segment":        {uri: "foo://a/foo/bar/", segments: []string{"x/y/z", ".."}, want: "foo://a/foo/bar/x/y"},
		"success: untitled relative":     {uri: "untitled:untitled-1", segments: []string{"..", "untitled-2"}, want: "untitled:untitled-2"},
		"success: normalize only":        {uri: "foo://a/a/foo/bar//x", want: "foo://a/a/foo/bar/x"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			u := MustParse(tt.uri)
			got, err := JoinPath(u, tt.segments...)
			if err != nil {
				t.Fatalf("JoinPath() error = %v", err)
			}
			if got.String() != tt.want {
				t.Fatalf("JoinPath().String() = %q, want %q", got.String(), tt.want)
			}
		})
	}
}

func TestResolvePath(t *testing.T) {
	tests := map[string]struct {
		uri     string
		segment string
		want    string
	}{
		"success: append relative":         {uri: "foo://a/foo/bar", segment: "x", want: "foo://a/foo/bar/x"},
		"success: base trailing":           {uri: "foo://a/foo/bar/", segment: "x", want: "foo://a/foo/bar/x"},
		"success: absolute overrides":      {uri: "foo://a/foo/bar/", segment: "/x", want: "foo://a/x"},
		"success: strip result trailing":   {uri: "foo://a/foo/bar/", segment: "x/", want: "foo://a/foo/bar/x"},
		"success: empty authority path":    {uri: "foo://a", segment: "x/", want: "foo://a/x"},
		"success: collapse complex abs":    {uri: "foo://a/b", segment: "/x/..//y/.", want: "foo://a/y"},
		"success: collapse complex rel":    {uri: "foo://a/b", segment: "x/..//y/.", want: "foo://a/b/y"},
		"success: untitled relative":       {uri: "untitled:untitled-1", segment: "../foo", want: "untitled:foo"},
		"success: untitled empty to child": {uri: "untitled:", segment: "foo", want: "untitled:foo"},
		"success: untitled empty parent":   {uri: "untitled:", segment: "..", want: "untitled:"},
		"success: untitled abs strips":     {uri: "untitled:", segment: "/foo", want: "untitled:foo"},
		"success: untitled slash keeps":    {uri: "untitled:/", segment: "/foo", want: "untitled:/foo"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			u := MustParse(tt.uri)
			got, err := ResolvePath(u, tt.segment)
			if err != nil {
				t.Fatalf("ResolvePath() error = %v", err)
			}
			if got.String() != tt.want {
				t.Fatalf("ResolvePath().String() = %q, want %q", got.String(), tt.want)
			}
		})
	}
}

func TestPathNameUtilities(t *testing.T) {
	tests := map[string]struct {
		uri      string
		dirname  string
		basename string
		extname  string
	}{
		"success: file leaf":       {uri: "foo://a/some/file/test.txt", dirname: "foo://a/some/file", basename: "test.txt", extname: ".txt"},
		"success: trailing slash":  {uri: "foo://a/some/file/", dirname: "foo://a/some", basename: "file", extname: ""},
		"success: trailing many":   {uri: "foo://a/some/file///", dirname: "foo://a/some", basename: "file", extname: ""},
		"success: no extension":    {uri: "foo://a/foo/bar", dirname: "foo://a/foo", basename: "bar", extname: ""},
		"success: extension":       {uri: "foo://a/foo/bar.foo", dirname: "foo://a/foo", basename: "bar.foo", extname: ".foo"},
		"success: dotfile no ext":  {uri: "foo://a/foo/.foo", dirname: "foo://a/foo", basename: ".foo", extname: ""},
		"success: authority root":  {uri: "foo://a/", dirname: "foo://a/", basename: "", extname: ""},
		"success: authority empty": {uri: "foo://a", dirname: "foo://a", basename: "", extname: ""},
		"success: scheme empty":    {uri: "foo://", dirname: "foo:", basename: "", extname: ""},
		"success: untitled":        {uri: "untitled:untitled-1", dirname: "untitled:", basename: "untitled-1", extname: ""},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			u := MustParse(tt.uri)
			if got := Dirname(u).String(); got != tt.dirname {
				t.Fatalf("Dirname().String() = %q, want %q", got, tt.dirname)
			}
			if got := Basename(u); got != tt.basename {
				t.Fatalf("Basename() = %q, want %q", got, tt.basename)
			}
			if got := Extname(u); got != tt.extname {
				t.Fatalf("Extname() = %q, want %q", got, tt.extname)
			}
		})
	}
}
