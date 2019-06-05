// Copyright 2019 The go-language-server Authors. All rights reserved.
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
		name    string
		s       string
		want    *uri.URI
		wantErr bool
	}{
		{
			name: "ValidFileScheme",
			s:    "file://code.visualstudio.com/docs/extensions/overview.md",
			want: &uri.URI{
				Scheme:    uri.FileScheme,
				Authority: "code.visualstudio.com",
				Path:      "/docs/extensions/overview.md",
				FsPath:    filepath.FromSlash("/docs/extensions/overview.md"),
			},
			wantErr: false,
		},
		{
			name: "ValidHTTPScheme",
			s:    "http://code.visualstudio.com/docs/extensions/overview#frag",
			want: &uri.URI{
				Scheme:    uri.HTTPScheme,
				Authority: "code.visualstudio.com",
				Path:      "/docs/extensions/overview",
				Fragment:  "frag",
			},
			wantErr: false,
		},
		{
			name: "ValidHTTPSScheme",
			s:    "https://code.visualstudio.com/docs/extensions/overview#frag",
			want: &uri.URI{
				Scheme:    uri.HTTPSScheme,
				Authority: "code.visualstudio.com",
				Path:      "/docs/extensions/overview",
				Fragment:  "frag",
			},
			wantErr: false,
		},
		{
			name:    "Invalid",
			s:       "foo://user@example.com:8042/over/there?name=ferret#nose",
			want:    new(uri.URI),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(uri.Parse(tt.s), tt.want, cmp.AllowUnexported(*tt.want)); (diff != "") != tt.wantErr {
				t.Errorf("%s: (-got, +want)\n%s", tt.name, diff)
			}
		})
	}
}

func TestFile(t *testing.T) {
	tests := []struct {
		name    string
		path    string
		want    *uri.URI
		wantErr bool
	}{
		{
			name: "ValidFileScheme",
			path: "/users/me/c#-projects/",
			want: &uri.URI{
				Scheme: uri.FileScheme,
				Path:   "/users/me/c#-projects/",
				FsPath: filepath.FromSlash("/users/me/c#-projects/"),
			},
			wantErr: false,
		},
		{
			name:    "Invalid",
			path:    "users-me-c#-projects",
			want:    new(uri.URI),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(uri.File(tt.path), tt.want, cmp.AllowUnexported(*tt.want)); (diff != "") != tt.wantErr {
				t.Errorf("%s: (-got, +want)\n%s", tt.name, diff)
			}
		})
	}
}

func TestFrom(t *testing.T) {
	type args struct {
		scheme    string
		authority string
		path      string
		query     string
		fragment  string
	}
	tests := []struct {
		name string
		args args
		want *uri.URI
	}{
		{
			name: "ValidFileScheme",
			args: args{
				scheme:    "file",
				authority: "example.com",
				path:      "/over/there",
				query:     "name=ferret",
				fragment:  "nose",
			},
			want: &uri.URI{
				Scheme:    uri.FileScheme,
				Authority: "example.com",
				Path:      "/over/there",
				FsPath:    filepath.FromSlash("/over/there"),
				Query:     "name=ferret",
				Fragment:  "nose",
			},
		},
		{
			name: "ValidHTTPScheme",
			args: args{
				scheme:    "http",
				authority: "example.com:8042",
				path:      "/over/there",
				query:     "name=ferret",
				fragment:  "nose",
			},
			want: &uri.URI{
				Scheme:    uri.HTTPScheme,
				Authority: "example.com:8042",
				Path:      "/over/there",
				Query:     "name=ferret",
				Fragment:  "nose",
			},
		},
		{
			name: "ValidHTTPSScheme",
			args: args{
				scheme:    "https",
				authority: "example.com:8042",
				path:      "/over/there",
				query:     "name=ferret",
				fragment:  "nose",
			},
			want: &uri.URI{
				Scheme:    uri.HTTPSScheme,
				Authority: "example.com:8042",
				Path:      "/over/there",
				Query:     "name=ferret",
				Fragment:  "nose",
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(uri.From(tt.args.scheme, tt.args.authority, tt.args.path, tt.args.query, tt.args.fragment), tt.want, cmp.AllowUnexported(*tt.want)); diff != "" {
				t.Errorf("%s: (-got, +want)\n%s", tt.name, diff)
			}
		})
	}
}

func TestURI_String(t *testing.T) {
	type fields struct {
		Authority string
		Fragment  string
		FsPath    string
		Path      string
		Query     string
		Scheme    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "ValidFileScheme",
			fields: fields{
				Authority: "code.visualstudio.com",
				Path:      "/docs/extensions/overview.md",
				FsPath:    filepath.FromSlash("/docs/extensions/overview.md"),
				Scheme:    string(uri.FileScheme),
			},
			want: "file://code.visualstudio.com/docs/extensions/overview.md",
		},
		{
			name: "ValidHTTPScheme",
			fields: fields{
				Authority: "code.visualstudio.com",
				Fragment:  "frag",
				Path:      "/docs/extensions/overview",
				FsPath:    filepath.FromSlash("/docs/extensions/overview"),
				Query:     "test",
				Scheme:    string(uri.HTTPScheme),
			},
			want: "http://code.visualstudio.com/docs/extensions/overview?test#frag",
		},
		{
			name: "ValidHTTPSScheme",
			fields: fields{
				Authority: "code.visualstudio.com",
				Fragment:  "frag",
				Path:      "/docs/extensions/overview",
				FsPath:    filepath.FromSlash("/docs/extensions/overview"),
				Query:     "test",
				Scheme:    string(uri.HTTPSScheme),
			},
			want: "https://code.visualstudio.com/docs/extensions/overview?test#frag",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			u := &uri.URI{
				Authority: tt.fields.Authority,
				Fragment:  tt.fields.Fragment,
				FsPath:    tt.fields.FsPath,
				Path:      tt.fields.Path,
				Query:     tt.fields.Query,
				Scheme:    tt.fields.Scheme,
			}

			if got := u.String(); !cmp.Equal(got, tt.want) {
				t.Errorf("URI.String() = %v, want %v", got, tt.want)
			}
			if got2 := u.String(); !cmp.Equal(got2, tt.want) { // cache with u.formatted
				t.Errorf("URI.String() = %v, want %v", got2, tt.want)
			}
		})
	}
}
