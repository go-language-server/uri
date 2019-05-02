// Copyright 2019 The uri Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package uri implementation vscode-uri for Go.
package uri

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

// Components
type Components struct {
	*URI
}

// State
type State struct {
	*URI
	Mid      float64 `json:"$mid,omitempty"`
	External string  `json:"external,omitempty"`
}
