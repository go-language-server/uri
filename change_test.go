// Copyright 2026 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"encoding/json"
	"errors"
	"testing"
)

func TestWith(t *testing.T) {
	str := func(s string) *string { return &s }
	tests := map[string]struct {
		base      string
		change    Change
		want      string
		wantError error
	}{
		"success: zero change returns same value": {
			base: "foo:bar/path",
			want: "foo:bar/path",
		},
		"success: same components returns same value": {
			base:   "foo:bar/path",
			change: Change{Scheme: str("foo"), Path: str("bar/path")},
			want:   "foo:bar/path",
		},
		"success: replace scheme": {
			base:   "before:some/file/path",
			change: Change{Scheme: str("after")},
			want:   "after:some/file/path",
		},
		"success: empty scheme follows vscode file fallback": {
			base:   "foo:bar/path",
			change: Change{Scheme: str("")},
			want:   "file:///bar/path",
		},
		"success: http query encodes": {
			base:   "s:",
			change: Change{Scheme: str("http"), Path: str("/api/files/test.me"), Query: str("t=1234")},
			want:   "http:/api/files/test.me?t%3D1234",
		},
		"success: remove authority with empty string": {
			base:   "scheme://authority/path",
			change: Change{Authority: str("")},
			want:   "scheme:/path",
		},
		"success: remove authority with clear pointer": {
			base:   "scheme://authority/path",
			change: Change{Authority: str("")},
			want:   "scheme:/path",
		},
		"success: clear path leaves authority": {
			base:   "scheme:/path",
			change: Change{Authority: str("authority"), Path: str("")},
			want:   "scheme://authority",
		},
		"error: invalid scheme": {
			base:      "foo:bar/path",
			change:    Change{Scheme: str("fai:l")},
			wantError: ErrInvalidScheme,
		},
		"error: authority requires slash path": {
			base:      "foo:bar/path",
			change:    Change{Authority: str("fail")},
			wantError: ErrAuthorityPath,
		},
		"error: path without authority cannot start double slash": {
			base:      "foo:bar/path",
			change:    Change{Path: str("//fail")},
			wantError: ErrPathAuthority,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			base := MustParse(tt.base)
			got, err := base.With(tt.change)
			if tt.wantError != nil {
				if !errors.Is(err, tt.wantError) {
					t.Fatalf("With() error = %v, want %v", err, tt.wantError)
				}
				return
			}
			if err != nil {
				t.Fatalf("With() error = %v", err)
			}
			if got.String() != tt.want {
				t.Fatalf("With().String() = %q, want %q", got.String(), tt.want)
			}
		})
	}
}

func TestTextMarshaling(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"success: canonical text": {input: "https://host/p?name=ferret#f", want: "https://host/p?name%3Dferret#f"},
		"success: file text":      {input: "file:///a b.go", want: "file:///a%20b.go"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			u := MustParse(tt.input)
			text, err := u.MarshalText()
			if err != nil {
				t.Fatalf("MarshalText() error = %v", err)
			}
			if string(text) != tt.want {
				t.Fatalf("MarshalText() = %q, want %q", text, tt.want)
			}
			var got URI
			if err := got.UnmarshalText(text); err != nil {
				t.Fatalf("UnmarshalText() error = %v", err)
			}
			if got != u {
				t.Fatalf("UnmarshalText() = %q, want %q", got.String(), u.String())
			}
		})
	}
}

func TestJSONTextMarshaling(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"success: canonical JSON string": {
			input: "https://host/p?name=ferret#f",
			want:  `"https://host/p?name%3Dferret#f"`,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			u := MustParse(tt.input)
			data, err := json.Marshal(u)
			if err != nil {
				t.Fatalf("json.Marshal() error = %v", err)
			}
			if string(data) != tt.want {
				t.Fatalf("json.Marshal() = %s, want %s", data, tt.want)
			}
			var got URI
			if err := json.Unmarshal(data, &got); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			if got != u {
				t.Fatalf("json.Unmarshal() = %q, want %q", got.String(), u.String())
			}
		})
	}
}

func TestMustParsePanic(t *testing.T) {
	tests := map[string]struct {
		input string
	}{
		"success: invalid URI panics": {input: "fäil:path"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			defer func() {
				if recover() == nil {
					t.Fatal("MustParse did not panic")
				}
			}()
			_ = MustParse(tt.input)
		})
	}
}
