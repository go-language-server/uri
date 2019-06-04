// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"fmt"
	"net/url"
	"path/filepath"
)

// FileScheme schema of filesystem path.
const FileScheme = "file"

// URI Uniform Resource Identifier (URI) http://tools.ietf.org/html/rfc3986.
//
// This class is a simple parser which creates the basic component parts
// (http://tools.ietf.org/html/rfc3986#section-3) with minimal validation
// and encoding.
//
//        foo://example.com:8042/over/there?name=ferret#nose
//        \_/   \______________/\_________/ \_________/ \__/
//         |           |            |            |        |
//      scheme     authority       path        query   fragment
//         |   _____________________|__
//        / \ /                        \
//        urn:example:animal:ferret:nose
type URI struct {
	// Authority is the 'www.msft.com' part of 'http://www.msft.com/some/path?query#fragment'.
	// The part between the first double slashes and the next slash.
	Authority string `json:"authority,omitempty"`

	// Fragment is the 'fragment' part of 'http://www.msft.com/some/path?query#fragment'.
	Fragment string `json:"fragment,omitempty"`

	// FsPath returns a string representing the corresponding file system path of this URI.
	//
	// Will handle UNC paths, normalizes windows drive letters to lower-case, and uses the
	// platform specific path separator.
	//
	// * Will *not* validate the path for invalid characters and semantics.
	// * Will *not* look at the scheme of this URI.
	// * The result shall *not* be used for display purposes but for accessing a file on disk.
	//
	//
	// The *difference* to `URI#path` is the use of the platform specific separator and the handling
	// of UNC paths. See the below sample of a file-uri with an authority (UNC path).
	//
	//  const u = URI.parse('file://server/c$/folder/file.txt')
	//  u.authority === 'server'
	//  u.path === '/shares/c$/file.txt'
	//  u.fsPath === '\\server\c$\folder\file.txt'
	//
	// Using `URI#path` to read a file (using fs-apis) would not be enough because parts of the path,
	// namely the server name, would be missing.
	//
	// Therefore `URI#fsPath` exists - it's sugar to ease working with URIs that represent files on disk (`file` scheme).
	FsPath string `json:"fsPath,omitempty"`

	// Path is the '/some/path' part of 'http://www.msft.com/some/path?query#fragment'.
	Path string `json:"path,omitempty"`

	// Query is the 'query' part of 'http://www.msft.com/some/path?query#fragment'.
	Query string `json:"query,omitempty"`

	// Scheme is the 'http' part of 'http://www.msft.com/some/path?query#fragment'.
	//
	// The part before the first colon.
	Scheme string `json:"scheme,omitempty"`
}

// String implements fmt.Stringer.
func (u *URI) String() string {
	switch u.Scheme {
	case FileScheme:
		uri := &url.URL{
			Scheme: u.Scheme,
			Path:   u.Path,
		}
		return uri.String()

	case "http", "https":
		uri := &url.URL{
			Scheme:   u.Scheme,
			Host:     u.Authority,
			Path:     u.Path,
			RawQuery: url.QueryEscape(u.Query),
			Fragment: u.Fragment,
		}
		return uri.String()

	default:
		return "unknown schema"
	}
}

// State represents a URI State.
type State struct {
	*URI
	Mid      float64 `json:"$mid,omitempty"`
	External string  `json:"external,omitempty"`
}

// Parse parses and creates a new URI from uri.
func Parse(s string) *URI {
	u, err := url.Parse(s)
	if err != nil {
		panic(fmt.Sprintf("url.Parse: %#v\n", err))
	}

	return &URI{
		Scheme:    u.Scheme,
		Authority: u.Host,
		Path:      u.Path,
		Query:     u.Query().Encode(),
		Fragment:  u.Fragment,
	}
}

// File parses and creates a new URI filesystem path from path.
func File(path string) *URI {
	return &URI{
		Scheme: FileScheme,
		Path:   path,
		FsPath: filepath.FromSlash(path),
	}
}

// From returns the new URI from args.
func From(scheme, authority, path, query, fragment string) *URI {
	return &URI{
		Scheme:    scheme,
		Authority: authority,
		Path:      path,
		Query:     query,
		Fragment:  fragment,
	}
}
