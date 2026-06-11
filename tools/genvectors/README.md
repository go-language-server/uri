# URI conformance vector generator

This directory pins the Node reference dependency used to regenerate
`../../testdata/vectors.json` from upstream `vscode-uri`.

The Go package intentionally stores only the canonical string identity so that
`URI` stays comparable and safe as a map key. For parse vectors, the generator
therefore records:

1. `URI.parse(input).toString()` from pinned `vscode-uri`.
2. Components and filesystem paths from reparsing that canonical string.

That makes the corpus explicit about the Go contract: canonical component
accessors, not original parse-history casing from the first JavaScript object.

## Normal regeneration

```sh
npm ci --prefix tools/genvectors
node tools/genvectors/main.mjs
```

The default generator is intentionally fail-fast: if the pinned `vscode-uri`
dependency is unavailable, generation exits with an error instead of silently
reusing the checked-in fixture. `go test` also asserts that the committed corpus
has `"generator": "vscode-uri-canonical-reparse"`, a
`vscodeURIVersion`, and the `go-comparable-canonical-uri` contract metadata.

## Explicit static fixture mode

```sh
node tools/genvectors/main.mjs --use-static-fixture
```

Static mode is only for offline inspection or emergency fixture preservation. Do
not use it for conformance updates, CI, or final release verification.
