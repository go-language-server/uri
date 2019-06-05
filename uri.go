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
	"strconv"
	"strings"

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

const (
	keyScheme    = "scheme"
	keyAuthority = "authority"
	keyPath      = "path"
	keyFsPath    = "fsPath"
	keyQuery     = "query"
	keyFragment  = "fragment"
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
	// Scheme is the 'http' part of 'http://www.msft.com/some/path?query#fragment'.
	//
	// The part before the first colon.
	Scheme string `json:"scheme"`

	// Authority is the 'www.msft.com' part of 'http://www.msft.com/some/path?query#fragment'.
	// The part between the first double slashes and the next slash.
	Authority string `json:"authority"`

	// Path is the '/some/path' part of 'http://www.msft.com/some/path?query#fragment'.
	Path string `json:"path"`

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

	// Query is the 'query' part of 'http://www.msft.com/some/path?query#fragment'.
	Query string `json:"query,omitempty"`

	// Fragment is the 'fragment' part of 'http://www.msft.com/some/path?query#fragment'.
	Fragment string `json:"fragment,omitempty"`

	formatted    string
	skipEncoding bool
}

// MarshalJSON implements json.Marshaler.
func (u *URI) MarshalJSON() ([]byte, error) {
	buf := new(bytes.Buffer)

	buf.WriteString("{")

	buf.WriteString(`"` + strconv.Quote(keyScheme) + `: "`)
	scheme, err := json.Marshal(u.Scheme)
	if err != nil {
		return nil, err
	}
	buf.Write(scheme)

	buf.WriteString(",")
	buf.WriteString(`"` + strconv.Quote(keyAuthority) + `: "`)
	authority, err := json.Marshal(u.Authority)
	if err != nil {
		return nil, err
	}
	buf.Write(authority)

	buf.WriteString(",")
	buf.WriteString(`"` + strconv.Quote(keyPath) + `: "`)
	path, err := json.Marshal(u.Path)
	if err != nil {
		return nil, err
	}
	buf.Write(path)

	if u.Query != "" {
		buf.WriteString(",")
		buf.WriteString(`"` + strconv.Quote(keyQuery) + `: "`)
		query, err := json.Marshal(u.Query)
		if err != nil {
			return nil, err
		}
		buf.Write(query)
	}

	if u.FsPath != "" {
		buf.WriteString(",")
		buf.WriteString(`"` + strconv.Quote(keyFsPath) + `: "`)
		fsPath, err := json.Marshal(u.FsPath)
		if err != nil {
			return nil, err
		}
		buf.Write(fsPath)
	}

	if u.Fragment != "" {
		buf.WriteString(",")
		buf.WriteString(`"` + strconv.Quote(keyFragment) + `: "`)
		fragment, err := json.Marshal(u.Fragment)
		if err != nil {
			return nil, err
		}
		buf.Write(fragment)
	}

	buf.WriteString("}")

	return buf.Bytes(), nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (u *URI) UnmarshalJSON(b []byte) error {
	var schemeReceived bool
	var authorityReceived bool
	var pathReceived bool
	var jm map[string]json.RawMessage

	if err := json.Unmarshal(b, &jm); err != nil {
		return err
	}

	// parse all the defined properties
	for k, v := range jm {
		switch k {
		case keyScheme:
			if err := json.Unmarshal([]byte(v), &u.Scheme); err != nil {
				return err
			}
			schemeReceived = true

		case keyAuthority:
			if err := json.Unmarshal([]byte(v), &u.Authority); err != nil {
				return err
			}
			authorityReceived = true

		case keyPath:
			if err := json.Unmarshal([]byte(v), &u.Path); err != nil {
				return err
			}
			pathReceived = true

		case keyFsPath:
			if err := json.Unmarshal([]byte(v), &u.FsPath); err != nil {
				return err
			}

		case keyQuery:
			if err := json.Unmarshal([]byte(v), &u.Query); err != nil {
				return err
			}

		case keyFragment:
			if err := json.Unmarshal([]byte(v), &u.Fragment); err != nil {
				return err
			}

		default:
			return fmt.Errorf("additional property not allowed: \"" + k + "\"")
		}
	}

	// check if scheme (a required property) was received
	if !schemeReceived {
		return xerrors.Errorf("%q is required but was not present", keyScheme)
	}

	// check if authority (a required property) was received
	if !authorityReceived {
		return xerrors.Errorf("%q is required but was not present", keyAuthority)
	}

	// check if path (a required property) was received
	if !pathReceived {
		return xerrors.Errorf("%q is required but was not present", keyPath)
	}

	return nil
}

// String implements fmt.Stringer.
func (u URI) String() string {
	switch u.Scheme {
	case FileScheme, HTTPScheme, HTTPSScheme:
		if u.skipEncoding {
			return u.formatted
		}

		u.formatted = format(u, true)
		u.skipEncoding = true

		return u.formatted

	default:
		return "unknown scheme"
	}
}

var encodeTable = map[byte]string{
	Colon:              "%3A", // gen-delims
	Slash:              "%2F",
	QuestionMark:       "%3F",
	Hash:               "%23",
	OpenSquareBracket:  "%5B",
	CloseSquareBracket: "%5D",
	AtSign:             "%40",
	ExclamationMark:    "%21", // sub-delims
	DollarSign:         "%24",
	Ampersand:          "%26",
	SingleQuote:        "%27",
	OpenParen:          "%28",
	CloseParen:         "%29",
	Asterisk:           "%2A",
	Plus:               "%2B",
	Comma:              "%2C",
	Semicolon:          "%3B",
	Equals:             "%3D",
	Space:              "%20",
}

func encodeFast(uriComponent string, allowSlash bool) string {
	b := new(strings.Builder)
	nativeEncodePos := -1

	for pos := 0; pos < len(uriComponent); pos++ {
		code := uriComponent[pos]

		switch {
		case code >= LowerA && code <= LowerZ,
			code >= UpperA && code <= UpperZ,
			code >= Digit0 && code <= Digit9,
			code == Dash,
			code == Period,
			code == Underline,
			code == Tilde,
			allowSlash && (code == Slash):

			if nativeEncodePos != -1 {
				b.WriteString(uriComponent[nativeEncodePos:pos])
				nativeEncodePos = -1
			}
			str := b.String()
			if str != "" {
				b.WriteString(str)
				b.WriteString(uriComponent[:pos])
			}

		default:
			str := b.String()
			if str != "" {
				b.WriteString(str)
				b.WriteString(uriComponent[0:pos])
			}

			escaped := encodeTable[code]
			if escaped != "" {
				if nativeEncodePos != -1 {
					b.WriteString(uriComponent[nativeEncodePos:pos])
				}

				b.WriteString(escaped)
			} else if nativeEncodePos == -1 {
				nativeEncodePos = pos
			}
		}
	}

	if nativeEncodePos != -1 {
		b.WriteString(uriComponent[:nativeEncodePos])
	}

	return b.String()
}

func encodeMinimal(path string, _ bool) string {
	b := new(strings.Builder)

	for i := 0; i < len(path); i++ {
		code := path[i]
		if code == Hash || code == QuestionMark {
			if b.String() == "" {
				b.WriteString(path[0:i])
			}
			b.WriteString(encodeTable[code])
		} else if b.String() != "" {
			b.WriteByte(path[i])
		}
	}

	res := b.String()
	if res == "" {
		res = path
	}

	return res
}

func format(uri URI, skipEncoding bool) string {
	var encoder func(string, bool) string
	switch skipEncoding {
	case true:
		encoder = encodeMinimal
	case false:
		encoder = encodeFast
	}

	b := new(strings.Builder)

	scheme := uri.Scheme
	if scheme != "" {
		b.WriteString(scheme)
		b.WriteByte(':')
	}

	authority := uri.Authority
	if authority != "" || scheme == FileScheme {
		b.WriteRune(filepath.Separator)
		b.WriteRune(filepath.Separator)
	}

	if authority != "" {
		idx := strings.LastIndex(authority, "@")
		if idx != -1 {
			// <user>@<auth>
			userinfo := authority[:idx]
			authority = authority[idx+1:]

			uiIdx := strings.Index(userinfo, ":")
			if uiIdx == -1 {
				b.WriteString(encoder(userinfo, false))
			} else {
				// <user>:<pass>@<auth>
				b.WriteString(encoder(userinfo[:idx], false))
				b.WriteRune(':')
				b.WriteString(encoder(userinfo[idx+1:], false))
			}

			b.WriteRune('@')
		}

		authority = strings.ToLower(authority)

		idx = strings.Index(authority, ":")
		if idx == -1 {
			b.WriteString(encoder(authority, false))
		} else {
			// <auth>:<port>
			b.WriteString(encoder(authority[:idx], false))
			b.WriteString(authority[idx:])
		}
	}

	if path := uri.Path; path != "" {
		// lower-case windows drive letters in /C:/fff or C:/fff
		if len(path) >= 3 && path[0] == Slash && path[2] == Colon {
			code := path[1]
			if code >= UpperA && code <= UpperZ {
				path = "/" + string(code+32) + ":" + string(path[3]) // "/c:".length == 3
			}
		} else if len(path) >= 2 && path[1] == Colon {
			code := path[0]
			if code >= UpperA && code <= UpperZ {
				path = string(code+32) + ":" + string(path[2]) // "/c:".length == 3
			}
		}

		// encode the rest of the path
		b.WriteString(encoder(path, true))
	}

	if query := uri.Query; query != "" {
		b.WriteRune('?')
		b.WriteString(encoder(query, false))
	}

	if fragment := uri.Fragment; fragment != "" {
		b.WriteRune('#')

		if skipEncoding {
			b.WriteString(fragment)
		} else {
			b.WriteString(encodeFast(fragment, false))
		}
	}

	return b.String()
}

// Parse parses and creates a new URI from uri.
func Parse(s string) (u *URI, err error) {
	us, err := url.Parse(s)
	if err != nil {
		return nil, xerrors.Errorf("url.Parse: %w\n", err)
	}

	switch us.Scheme {
	case FileScheme:
		u = &URI{
			Scheme: FileScheme,
			Path:   us.Path,
			FsPath: filepath.FromSlash(us.Path),
		}

	case HTTPScheme, HTTPSScheme:
		u = &URI{
			Scheme:    us.Scheme,
			Authority: us.Host,
			Path:      us.Path,
			Query:     us.Query().Encode(),
			Fragment:  us.Fragment,
		}
	default:
		return nil, xerrors.New("unknown scheme")
	}

	return u, nil
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
func From(scheme, authority, path, query, fragment string) (u *URI) {
	switch scheme {
	case FileScheme:
		u = &URI{
			Scheme: FileScheme,
			Path:   path,
			FsPath: filepath.FromSlash(path),
		}

	case HTTPScheme, HTTPSScheme:
		u = &URI{
			Scheme:    scheme,
			Authority: authority,
			Path:      path,
			Query:     url.QueryEscape(query),
			Fragment:  fragment,
		}
	}

	return u
}
