// Copyright 2019 The go-language-server Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uri

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"path/filepath"

	"golang.org/x/xerrors"
)

const (
	// FileScheme schema of filesystem path.
	FileScheme = "file"

	// HTTPScheme schema of http.
	HTTPScheme = "http"

	// HTTPSScheme schema of https.
	HTTPSScheme = "https"
)

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
	Authority string `json:"authority"`

	// Fragment is the 'fragment' part of 'http://www.msft.com/some/path?query#fragment'.
	Fragment string `json:"fragment"`

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
	FsPath string `json:"fsPath"`

	// Path is the '/some/path' part of 'http://www.msft.com/some/path?query#fragment'.
	Path string `json:"path"`

	// Query is the 'query' part of 'http://www.msft.com/some/path?query#fragment'.
	Query string `json:"query"`

	// Scheme is the 'http' part of 'http://www.msft.com/some/path?query#fragment'.
	//
	// The part before the first colon.
	Scheme string `json:"scheme"`
}

// MarshalJSON implements json.Marshaler.
func (u *URI) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBuffer(make([]byte, 0))
	buf.WriteString("{")

	comma := false
	// "Authority" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "authority" field
	if comma {
		buf.WriteString(",")
	}

	buf.WriteString("\"authority\": ")
	if tmp, err := json.Marshal(u.Authority); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}

	comma = true
	// "Fragment" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "fragment" field
	if comma {
		buf.WriteString(",")
	}

	buf.WriteString("\"fragment\": ")
	if tmp, err := json.Marshal(u.Fragment); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}

	comma = true
	// "FsPath" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "fsPath" field
	if comma {
		buf.WriteString(",")
	}

	buf.WriteString("\"fsPath\": ")
	if tmp, err := json.Marshal(u.FsPath); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}

	comma = true
	// "Path" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "path" field
	if comma {
		buf.WriteString(",")
	}

	buf.WriteString("\"path\": ")
	if tmp, err := json.Marshal(u.Path); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}

	comma = true
	// "Query" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "query" field
	if comma {
		buf.WriteString(",")
	}

	buf.WriteString("\"query\": ")
	if tmp, err := json.Marshal(u.Query); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}

	comma = true
	// "Scheme" field is required
	// only required object types supported for marshal checking (for now)
	// Marshal the "scheme" field
	if comma {
		buf.WriteString(",")
	}

	buf.WriteString("\"scheme\": ")
	if tmp, err := json.Marshal(u.Scheme); err != nil {
		return nil, err
	} else {
		buf.Write(tmp)
	}

	comma = true

	buf.WriteString("}")
	rv := buf.Bytes()

	return rv, nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *URI) UnmarshalJSON(b []byte) error {
	var authorityReceived bool
	var fragmentReceived bool
	var fsPathReceived bool
	var pathReceived bool
	var queryReceived bool
	var schemeReceived bool
	var jm map[string]json.RawMessage

	if err := json.Unmarshal(b, &jm); err != nil {
		return err
	}

	// parse all the defined properties
	for k, v := range jm {
		switch k {
		case "authority":
			if err := json.Unmarshal([]byte(v), &u.Authority); err != nil {
				return err
			}
			authorityReceived = true

		case "fragment":
			if err := json.Unmarshal([]byte(v), &u.Fragment); err != nil {
				return err
			}
			fragmentReceived = true

		case "fsPath":
			if err := json.Unmarshal([]byte(v), &u.FsPath); err != nil {
				return err
			}
			fsPathReceived = true

		case "path":
			if err := json.Unmarshal([]byte(v), &u.Path); err != nil {
				return err
			}
			pathReceived = true

		case "query":
			if err := json.Unmarshal([]byte(v), &u.Query); err != nil {
				return err
			}
			queryReceived = true

		case "scheme":
			if err := json.Unmarshal([]byte(v), &u.Scheme); err != nil {
				return err
			}
			schemeReceived = true

		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}

	// check if authority (a required property) was received
	if !authorityReceived {
		return xerrors.New("\"authority\" is required but was not present")
	}

	// check if fragment (a required property) was received
	if !fragmentReceived {
		return xerrors.New("\"fragment\" is required but was not present")
	}

	// check if fsPath (a required property) was received
	if !fsPathReceived {
		return xerrors.New("\"fsPath\" is required but was not present")
	}

	// check if path (a required property) was received
	if !pathReceived {
		return xerrors.New("\"path\" is required but was not present")
	}

	// check if query (a required property) was received
	if !queryReceived {
		return xerrors.New("\"query\" is required but was not present")
	}

	// check if scheme (a required property) was received
	if !schemeReceived {
		return xerrors.New("\"scheme\" is required but was not present")
	}

	return nil
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

	case HTTPScheme, HTTPSScheme:
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
