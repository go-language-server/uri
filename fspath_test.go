// Copyright 2026 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"go/parser"
	"go/token"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestFileFor(t *testing.T) {
	tests := map[string]struct {
		platform Platform
		path     string
		want     string
		wantPath string
		wantAuth string
	}{
		"success: posix absolute clean": {
			platform: PlatformPOSIX,
			path:     "/abs/clean/path.go",
			want:     "file:///abs/clean/path.go",
			wantPath: "/abs/clean/path.go",
		},
		"success: posix relative is not absolutized": {
			platform: PlatformPOSIX,
			path:     "./foo/bar",
			want:     "file:///./foo/bar",
			wantPath: "/./foo/bar",
		},
		"success: posix backslash is filename byte": {
			platform: PlatformPOSIX,
			path:     `c:\win\path`,
			want:     "file:///c%3A%5Cwin%5Cpath",
			wantPath: `/c:\win\path`,
		},
		"success: windows backslash converts to slash": {
			platform: PlatformWindows,
			path:     `c:\win\path`,
			want:     "file:///c%3A/win/path",
			wantPath: "/c:/win/path",
		},
		"success: windows mixed separators": {
			platform: PlatformWindows,
			path:     `c:\win/path`,
			want:     "file:///c%3A/win/path",
			wantPath: "/c:/win/path",
		},
		"success: windows drive lowercases for canonical string and component view": {
			platform: PlatformWindows,
			path:     "C:/Win/Path",
			want:     "file:///c%3A/Win/Path",
			wantPath: "/c:/Win/Path",
		},
		"success: unc authority extracted": {
			platform: PlatformWindows,
			path:     `\\server\share\file.txt`,
			want:     "file://server/share/file.txt",
			wantAuth: "server",
			wantPath: "/share/file.txt",
		},
		"success: posix double slash also extracts authority": {
			platform: PlatformPOSIX,
			path:     "//server/share/file.txt",
			want:     "file://server/share/file.txt",
			wantAuth: "server",
			wantPath: "/share/file.txt",
		},
		"success: file-like input is treated as path": {
			platform: PlatformPOSIX,
			path:     "file://path/to/file",
			want:     "file:///file%3A//path/to/file",
			wantPath: "/file://path/to/file",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := FileFor(tt.platform, tt.path)
			if got.String() != tt.want {
				t.Fatalf("FileFor().String() = %q, want %q", got.String(), tt.want)
			}
			if got.Authority() != tt.wantAuth {
				t.Fatalf("Authority() = %q, want %q", got.Authority(), tt.wantAuth)
			}
			if got.Path() != tt.wantPath {
				t.Fatalf("Path() = %q, want %q", got.Path(), tt.wantPath)
			}
		})
	}
}

func TestFsPathFor(t *testing.T) {
	tests := map[string]struct {
		uri                   string
		platform              Platform
		keepDriveLetterCasing bool
		want                  string
	}{
		"success: posix clean file": {
			uri:      "file:///home/user/x.go",
			platform: PlatformPOSIX,
			want:     "/home/user/x.go",
		},
		"success: posix drive strips leading slash and lowercases": {
			uri:      "file:///C:/test/me",
			platform: PlatformPOSIX,
			want:     "c:/test/me",
		},
		"success: canonical posix drive remains lowercase when casing requested": {
			uri:                   "file:///C:/test/me",
			platform:              PlatformPOSIX,
			keepDriveLetterCasing: true,
			want:                  "c:/test/me",
		},
		"success: windows drive lowercases and converts slash": {
			uri:      "file:///C:/test/me",
			platform: PlatformWindows,
			want:     `c:\test\me`,
		},
		"success: posix unc includes authority": {
			uri:      "file://shares/files/c%23/p.cs",
			platform: PlatformPOSIX,
			want:     "//shares/files/c#/p.cs",
		},
		"success: windows unc includes authority": {
			uri:      "file://shares/files/c%23/p.cs",
			platform: PlatformWindows,
			want:     `\\shares\files\c#\p.cs`,
		},
		"success: unc authority uses canonical component casing": {
			uri:      "file://SERVER/Share/X.go",
			platform: PlatformPOSIX,
			want:     "//server/Share/X.go",
		},
		"success: file authority with root path does not become unc": {
			uri:      "file://server",
			platform: PlatformPOSIX,
			want:     "/",
		},
		"success: underscore is not drive": {
			uri:      "file:///_:/path",
			platform: PlatformPOSIX,
			want:     "/_:/path",
		},
		"success: non-file still returns path": {
			uri:      "foo://server/share/x.go",
			platform: PlatformPOSIX,
			want:     "/share/x.go",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			u := MustParse(tt.uri)
			if got := FsPathFor(u, tt.platform, tt.keepDriveLetterCasing); got != tt.want {
				t.Fatalf("FsPathFor() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestFileAndFsPathAllocationGates(t *testing.T) {
	tests := map[string]struct {
		alloc     func() float64
		maxAllocs float64
	}{
		"FileFor clean absolute uses at most one allocation": {
			maxAllocs: 1,
			alloc: func() float64 {
				return testing.AllocsPerRun(1000, func() {
					u := FileFor(PlatformPOSIX, "/abs/clean/path.go")
					if u.String() != "file:///abs/clean/path.go" {
						t.Fatalf("unexpected URI %q", u.String())
					}
				})
			},
		},
		"FsPathFor clean posix file is zero allocation": {
			maxAllocs: 0,
			alloc: func() float64 {
				u := MustParse("file:///home/user/x.go")
				return testing.AllocsPerRun(1000, func() {
					if got := FsPathFor(u, PlatformPOSIX, false); got != "/home/user/x.go" {
						t.Fatalf("FsPathFor() = %q", got)
					}
				})
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			allocs := tt.alloc()
			if allocs > tt.maxAllocs {
				t.Fatalf("allocs = %v, want <= %v", allocs, tt.maxAllocs)
			}
		})
	}
}

func TestNoFilesystemImportsInNonTestFiles(t *testing.T) {
	tests := map[string]struct {
		forbidden map[string]bool
	}{
		"success: no os or filepath imports": {
			forbidden: map[string]bool{"os": true, "path/filepath": true},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			files, err := filepath.Glob("*.go")
			if err != nil {
				t.Fatal(err)
			}
			for _, file := range files {
				if strings.HasSuffix(file, "_test.go") {
					continue
				}
				fset := token.NewFileSet()
				parsed, err := parser.ParseFile(fset, file, nil, parser.ImportsOnly)
				if err != nil {
					t.Fatalf("ParseFile(%s): %v", file, err)
				}
				for _, imp := range parsed.Imports {
					path, err := strconv.Unquote(imp.Path.Value)
					if err != nil {
						t.Fatalf("Unquote(%s): %v", imp.Path.Value, err)
					}
					if tt.forbidden[path] {
						t.Fatalf("%s imports forbidden package %q", file, path)
					}
				}
			}
		})
	}
}
