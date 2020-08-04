// Copyright 2019 The Go Language Server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"errors"
	"fmt"
	"net/url"
	"path/filepath"
	"runtime"
	"strings"
	"unicode"

	uripb "go.lsp.dev/api/uri"
)

type Scheme string

var (
	// FileScheme schema of filesystem path.
	FileScheme = Scheme(strings.ToLower(uripb.Scheme_SCHEME_FILE.Enum().String()))

	// HTTPScheme schema of http.
	HTTPScheme = Scheme(strings.ToLower(uripb.Scheme_SCHEME_HTTP.Enum().String()))

	// HTTPSScheme schema of https.
	HTTPSScheme = Scheme(strings.ToLower(uripb.Scheme_SCHEME_HTTPS.Enum().String()))
)

// URI Uniform Resource Identifier (URI) https://tools.ietf.org/html/rfc3986.
//
// This class is a simple parser which creates the basic component parts
// (http://tools.ietf.org/html/rfc3986#section-3) with minimal validation
// and encoding.
//
//        foo://example.com:8042/over/there?name=ferret#nose
//        \_/\/  \______________/\_________/ \_________/ \__/
//         |  \         |            |            |        |
//      scheme hier  authority       path        query   fragment
//         |   _____________________|__
//        / \ /                        \
//        urn:example:animal:ferret:nose
type URI = uripb.URI

const (
	hierPart = "://"
)

// Filename returns the file path for the given URI.
// It is an error to call this on a URI that is not a valid filename.
// func (u *URI) Filename() string {
// 	filename, err := filename(u)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	return filepath.FromSlash(filename)
// }

func filename(uri *URI) (string, error) {
	u, err := url.ParseRequestURI(uri.Uri)
	if err != nil {
		return "", fmt.Errorf("failed to parse request URI: %w", err)
	}

	if u.Scheme != string(FileScheme) {
		return "", fmt.Errorf("only file URIs are supported, got %v", u.Scheme)
	}

	if isWindowsDriveURI(u.Path) {
		u.Path = u.Path[1:]
	}

	return u.Path, nil
}

// New parses and creates a new URI from s.
func New(s string) *URI {
	if u, err := url.PathUnescape(s); err == nil {
		s = u
	}

	if strings.HasPrefix(s, string(FileScheme+hierPart)) {
		return &URI{
			Uri: s,
		}
	}

	return File(s)
}

// File parses and creates a new filesystem URI from path.
func File(path string) *URI {
	const goRootPragma = "$GOROOT"
	if len(path) >= len(goRootPragma) && strings.EqualFold(goRootPragma, path[:len(goRootPragma)]) {
		path = runtime.GOROOT() + path[len(goRootPragma):]
	}

	if !isWindowsDrivePath(path) {
		if abs, err := filepath.Abs(path); err == nil {
			path = abs
		}
	}

	if isWindowsDrivePath(path) {
		path = "/" + path
	}

	path = filepath.ToSlash(path)
	u := url.URL{
		Scheme: string(FileScheme),
		Path:   path,
	}

	return &URI{
		Uri: u.String(),
	}
}

// Parse parses and creates a new URI from s.
func Parse(s string) (u *URI, err error) {
	us, err := url.Parse(s)
	if err != nil {
		return u, fmt.Errorf("url.Parse: %w", err)
	}

	switch Scheme(us.Scheme) {
	case FileScheme:
		ut := url.URL{
			Scheme:  us.Scheme,
			Path:    us.Path,
			RawPath: filepath.FromSlash(us.Path),
		}
		u = &URI{
			Uri: ut.String(),
		}

	case HTTPScheme, HTTPSScheme:
		ut := url.URL{
			Scheme:   us.Scheme,
			Host:     us.Host,
			Path:     us.Path,
			RawQuery: us.Query().Encode(),
			Fragment: us.Fragment,
		}
		u = &URI{
			Uri: ut.String(),
		}

	default:
		return u, errors.New("unknown scheme")
	}

	return
}

// From returns the new URI from args.
func From(scheme, authority, path, query, fragment string) *URI {
	switch Scheme(scheme) {
	case FileScheme:
		u := url.URL{
			Scheme:  string(scheme),
			Path:    path,
			RawPath: filepath.FromSlash(path),
		}
		return &URI{
			Uri: u.String(),
		}

	case HTTPScheme, HTTPSScheme:
		u := url.URL{
			Scheme:   string(scheme),
			Host:     authority,
			Path:     path,
			RawQuery: url.QueryEscape(query),
			Fragment: fragment,
		}
		return &URI{
			Uri: u.String(),
		}

	default:
		panic(fmt.Sprintf("unknown scheme: %s", scheme))
	}
}

// isWindowsDrivePath returns true if the file path is of the form used by Windows.
//
// We check if the path begins with a drive letter, followed by a ":".
func isWindowsDrivePath(path string) bool {
	if len(path) < 4 {
		return false
	}
	return unicode.IsLetter(rune(path[0])) && path[1] == ':'
}

// isWindowsDriveURI returns true if the file URI is of the format used by
// Windows URIs. The url.Parse package does not specially handle Windows paths
// (see https://golang.org/issue/6027). We check if the URI path has
// a drive prefix (e.g. "/C:"). If so, we trim the leading "/".
func isWindowsDriveURI(uri string) bool {
	if len(uri) < 4 {
		return false
	}
	return uri[0] == '/' && unicode.IsLetter(rune(uri[1])) && uri[2] == ':'
}
