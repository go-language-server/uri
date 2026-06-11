# Migration notes for the canonical vscode-uri rewrite

This branch is an intentional behavior rewrite of `go.lsp.dev/uri`. It replaces
the previous `net/url`-backed behavior with canonical `vscode-uri` semantics
while keeping the public representation as immutable comparable `type URI
string`.

## Removed or replaced APIs

| Old API | Replacement | Notes |
| --- | --- | --- |
| `New(string)` | `Parse(string)` or `File(string)` | `New` mixed URI parsing with filesystem fallback and could emit invalid URIs. Use `Parse` for URI text and `File`/`FileFor` for filesystem paths. |
| `Filename()` | `FsPath()` or `FsPathFor(URI, Platform, bool)` | `FsPathFor` exposes deterministic POSIX/Windows behavior for tests and cross-platform LSP code. |
| `From(scheme, authority, path, query, fragment)` | `From(Components)` | Components are decoded fields. Query and fragment are opaque components, matching `vscode-uri`. |
| `FileScheme`, `HTTPScheme`, `HTTPSScheme` | `Scheme()`, `IsFile()`, local constants if needed | The rewrite keeps scheme strings internal to avoid preserving the old narrow scheme list as API surface. |

## Behavioral differences to expect

- `Parse` follows `vscode-uri` non-strict parsing, including empty-scheme
  fallback to `file`.
- `ParseStrict` requires a scheme and reports typed sentinel errors.
- `File` and `FileFor` never call `os.Getwd`, `filepath.Abs`, or
  `runtime.GOROOT`; relative-looking input is encoded as provided under
  `vscode-uri` file semantics.
- Empty query and fragment delimiters are canonicalized away, matching
  `URI.parse(input).toString()` from `vscode-uri`.
- `URI` remains `type URI string`, so existing direct string conversions keep
  compiling. Prefer constructors for new values: direct `URI("...")`
  conversions preserve the supplied bytes and do not validate or canonicalize.
- Constructor-produced `URI` equality is canonical-string equality. To keep
  native Go `==` and map keys safe, component accessors are derived from the
  canonical string rather than from original parse-history casing. For example,
  `file://SERVER/x` and `file://server/x` compare equal and expose authority
  `server`; `file:///C:/x` and `file:///c%3A/x` compare equal and expose path
  `/c:/x`.
- `Change` mirrors `vscode-uri` `with` behavior: a nil field keeps the existing
  component; an empty string clears replaceable components; an empty `Scheme`
  follows the non-strict scheme fix and becomes `file`.

## Typical migrations

```go
// Before.
u := uri.New(text)
filename := u.Filename()

// After: URI text.
u, err := uri.Parse(text)
if err != nil {
	return err
}
filename := u.FsPath()

// After: filesystem path.
u = uri.File(path)
```

```go
// Before.
u := uri.From("https", "example.com", "/p", "q=1", "frag")

// After.
u, err := uri.From(uri.Components{
	Scheme:    "https",
	Authority: "example.com",
	Path:      "/p",
	Query:     "q=1",
	Fragment:  "frag",
})
if err != nil {
	return err
}
```
