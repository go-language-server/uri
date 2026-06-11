// Copyright 2026 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"strings"
	"testing"
)

func TestGeneratedTablesMatchVscodeURIEncodeContract(t *testing.T) {
	tests := map[string]struct {
		b             byte
		pathPass      bool
		authorityPass bool
		componentPass bool
	}{
		"unreserved alpha": {b: 'a', pathPass: true, authorityPass: true, componentPass: true},
		"unreserved tilde": {b: '~', pathPass: true, authorityPass: true, componentPass: true},
		"path slash":       {b: '/', pathPass: true, authorityPass: false, componentPass: false},
		"authority colon":  {b: ':', pathPass: false, authorityPass: true, componentPass: false},
		"authority open bracket": {
			b:             '[',
			pathPass:      false,
			authorityPass: true,
			componentPass: false,
		},
		"space escapes": {b: ' ', pathPass: false, authorityPass: false, componentPass: false},
		"percent escapes": {
			b:             '%',
			pathPass:      false,
			authorityPass: false,
			componentPass: false,
		},
		"sub delimiter bang escapes": {
			b:             '!',
			pathPass:      false,
			authorityPass: false,
			componentPass: false,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := canPassFast(tt.b, true, false); got != tt.pathPass {
				t.Fatalf("path pass = %t, want %t", got, tt.pathPass)
			}
			if got := canPassFast(tt.b, false, true); got != tt.authorityPass {
				t.Fatalf("authority pass = %t, want %t", got, tt.authorityPass)
			}
			if got := canPassFast(tt.b, false, false); got != tt.componentPass {
				t.Fatalf("component pass = %t, want %t", got, tt.componentPass)
			}
		})
	}
}

func TestPercentDecodeVscodeGracefulRules(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"success: no escape returns same text": {input: "abc", want: "abc"},
		"success: utf8 run decodes together":   {input: "%C3%BC", want: "ü"},
		"success: invalid hex remains literal": {input: "%zz", want: "%zz"},
		"success: invalid utf8 recurses":       {input: "%A0%C3%BC", want: "%A0ü"},
		"success: partial escape ignored":      {input: "%2-", want: "%2-"},
		"success: invalid leading utf8 then ascii": {
			input: "%C3%28",
			want:  "%C3(",
		},
		"success: valid prefix remains literal when suffix is invalid": {
			input: "%41%FF",
			want:  "%41%FF",
		},
		"success: valid suffix decodes after invalid triplets": {
			input: "%FF%41",
			want:  "%FFA",
		},
		"success: invalid hex preserves triplet and decodes later suffix": {
			input: "%zz%41",
			want:  "%zzA",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := percentDecode(tt.input); got != tt.want {
				t.Fatalf("percentDecode(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestPercentDecodeMalformedRunLarge(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"success: long malformed utf8 run stays literal without recursion": {
			input: strings.Repeat("%FF", 20_000),
			want:  strings.Repeat("%FF", 20_000),
		},
		"success: long malformed prefix with valid suffix decodes suffix": {
			input: strings.Repeat("%FF", 20_000) + "%41",
			want:  strings.Repeat("%FF", 20_000) + "A",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := percentDecode(tt.input); got != tt.want {
				t.Fatalf("percentDecode(long malformed run) length = %d, want %d", len(got), len(tt.want))
			}
		})
	}
}
