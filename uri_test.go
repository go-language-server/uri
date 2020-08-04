// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/proto"
)

func TestFile(t *testing.T) {
	tests := map[string]struct {
		path    string
		want    *URI
		wantErr bool
	}{
		"ValidFileScheme": {
			path: "/users/me/c#-projects/",
			want: &URI{
				Uri: FileScheme.Enum().String() + hierPart + "/users/me/c%23-projects",
			},
			wantErr: false,
		},
		"Invalid": {
			path:    "users-me-c#-projects",
			want:    &URI{},
			wantErr: true,
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// if !proto.Equal(File(tt.path), tt.want) {
			if diff := cmp.Diff(File(tt.path), tt.want, cmp.Comparer(proto.Equal)); (diff != "") != tt.wantErr {
				t.Fatalf("%s: (-want, +got)", name)
			}
		})
	}
}

func TestParse(t *testing.T) {
	tests := map[string]struct {
		s    string
		want *URI
	}{
		"ValidFileScheme": {

			s: "file://code.visualstudio.com/docs/extensions/overview.md",
			want: &URI{
				Uri: FileScheme.Enum().String() + hierPart + "/docs/extensions/overview.md",
			},
		},
		"ValidHTTPScheme": {
			s: "http://code.visualstudio.com/docs/extensions/overview#frag",
			want: &URI{
				Uri: FileScheme.Enum().String() + hierPart + "code.visualstudio.com/docs/extensions/overview#frag",
			},
		},
		"ValidHTTPSScheme": {
			s: "https://code.visualstudio.com/docs/extensions/overview#frag",
			want: &URI{
				Uri: HTTPSScheme.Enum().String() + hierPart + "code.visualstudio.com/docs/extensions/overview#frag",
			},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got, err := Parse(tt.s)
			if err != nil {
				t.Fatal(err)
				return
			}

			if diff := cmp.Diff(got, tt.want, cmp.Comparer(proto.Equal)); diff != "" {
				t.Fatalf("%s: (-got, +want)\n%s", name, diff)
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
	tests := map[string]struct {
		args args
		want *URI
	}{
		"ValidFileScheme": {
			args: args{
				scheme:    "file",
				authority: "example.com",
				path:      "/over/there",
				query:     "name=ferret",
				fragment:  "nose",
			},
			want: &URI{
				Uri: FileScheme.Enum().String() + hierPart + "/over/there",
			},
		},
		"ValidHTTPScheme": {
			args: args{
				scheme:    "http",
				authority: "example.com:8042",
				path:      "/over/there",
				query:     "name=ferret",
				fragment:  "nose",
			},
			want: &URI{
				Uri: HTTPScheme.Enum().String() + hierPart + "example.com:8042/over/there?name%3Dferret#nose",
			},
		},
		"ValidHTTPSScheme": {
			args: args{
				scheme:    "https",
				authority: "example.com:8042",
				path:      "/over/there",
				query:     "name=ferret",
				fragment:  "nose",
			},
			want: &URI{
				Uri: HTTPSScheme.Enum().String() + hierPart + "example.com:8042/over/there?name%3Dferret#nose",
			},
		},
	}
	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			if diff := cmp.Diff(From(tt.args.scheme, tt.args.authority, tt.args.path, tt.args.query, tt.args.fragment), tt.want, cmp.Comparer(proto.Equal)); diff != "" {
				t.Fatalf("%s: (-got, +want)\n%s", name, diff)
			}
		})
	}
}
