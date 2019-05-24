// Copyright 2019 The uri Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri_test

import (
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/go-language-server/uri"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want *uri.URI
	}{
		{
			name: "Valid",
			s:    "https://code.visualstudio.com/docs/extensions/overview#frag",
			want: &uri.URI{
				Scheme:    "https",
				Authority: "code.visualstudio.com",
				Path:      "/docs/extensions/overview",
				Fragment:  "frag",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(uri.Parse(tt.s), tt.want); diff != "" {
				t.Errorf("%s: (-got, +want)\n%s", tt.name, diff)
			}
		})
	}
}

func TestFile(t *testing.T) {
	tests := []struct {
		name string
		path string
		want *uri.URI
	}{
		{
			name: "Valid",
			path: "/users/me/c#-projects/",
			want: &uri.URI{
				Scheme: uri.FileScheme,
				Path:   "/users/me/c#-projects/",
				FsPath: filepath.FromSlash("/users/me/c#-projects/"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(uri.File(tt.path), tt.want); diff != "" {
				t.Errorf("%s: (-got, +want)\n%s", tt.name, diff)
			}
		})
	}
}
