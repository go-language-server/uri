# URI performance notes

This document records the benchmark and compiler-audit evidence for the
`vscode-uri`-compatible rewrite. The headline lane intentionally runs with
`GOEXPERIMENT` unset so the numbers are reproducible for downstream library
consumers that do not opt into local experiments.

## Reproduction

```sh
make bench
# or, without make:
env GOEXPERIMENT= go test -run '^$' -bench . -benchmem ./...
```

For publication-quality comparisons, run at least 10 samples and compare with
`benchstat`:

```sh
env GOEXPERIMENT= go test -run '^$' -bench . -benchmem -count=10 ./... > /tmp/uri.bench
benchstat /tmp/uri.bench
```

The committed benchmark corpus lives in `testdata/corpus/uri_bench.tsv` and
covers the LSP-shaped cases that should stay fast: clean POSIX `file://` URIs,
GOMODCACHE `@` paths, Windows drive paths, UNC file URIs, escaped Unicode paths,
HTTP(S) URIs with opaque query/fragment text, untitled URIs, and escape-heavy
inputs. The suite also includes `net/url` comparison scaffolding for parse and
map-key workloads.

## Headline environment

| Field | Value |
| --- | --- |
| Date | 2026-06-11 |
| Host | Apple M3 Max |
| OS/arch | darwin/arm64 |
| Go | go1.26.4 |
| GOEXPERIMENT | unset for headline benchmark command |
| Command | `make bench` (`env GOEXPERIMENT= go test -run '^$' -bench . -benchmem ./...`) |

## Current results

| Requirement | Current result | Target | Status |
| --- | ---: | ---: | --- |
| `Parse("file:///home/user/project/main.go")` | 15.5 ns/op, 0 B/op, 0 allocs/op | <= 40 ns/op, 0 allocs | PASS |
| Clean-but-non-canonical file URI (`@` corpus) | 201.1 ns/op, 64 B/op, 1 alloc/op | <= 1 alloc | PASS for allocation; latency tracked |
| `FsPathFor` on clean POSIX file URI | 11.5 ns/op, 0 B/op, 0 allocs/op | <= 15 ns/op, 0 allocs | PASS |
| `String()` | 1.9 ns/op, 0 B/op, 0 allocs/op | <= 2 ns/op, 0 allocs | PASS |
| `Parse("https://host/p?name=ferret#f")` | 157.0 ns/op, 32 B/op, 1 alloc/op | <= 80 ns/op, <= 2 allocs | GAP: latency |
| `FileFor(PlatformPOSIX, clean absolute path)` | 31.3 ns/op, 48 B/op, 1 alloc/op | <= 60 ns/op, <= 1 alloc | PASS |
| Map-key insert+lookup, `URI` | 453.0 us/op, 1.0 MiB/op, 33 allocs/op | >= 2x faster than net/url-shaped baseline | PASS |
| Map-key insert+lookup, `net/url` string baseline | 1105.2 us/op, 1.3 MiB/op, 20033 allocs/op | comparison | URI is 2.44x faster |

Malformed percent-run decoder benchmarks from the same run:

| Case | Current result | Shape |
| --- | ---: | --- |
| `strings.Repeat("%FF", 100)` | 544.7 ns/op | linear malformed run |
| `strings.Repeat("%FF", 1_000)` | 5.0 us/op | linear malformed run |
| `strings.Repeat("%FF", 10_000)` | 49.1 us/op | linear malformed run |
| `strings.Repeat("%FF", 10_000) + "%41"` | 51.1 us/op | linear malformed prefix + decoded suffix |

Representative `net/url` parse baselines from the same run:

| Case | `uri.Parse` | `net/url.Parse` baseline |
| --- | ---: | ---: |
| `file-posix-clean` | 15.5 ns/op, 0 allocs | 106.7 ns/op, 1 alloc |
| `file-gomodcache-at` | 201.1 ns/op, 1 alloc | 150.5 ns/op, 1 alloc |
| `https-query-fragment` | 157.0 ns/op, 1 alloc | 86.7 ns/op, 1 alloc |
| `escape-heavy` | 154.2 ns/op, 1 alloc | 152.3 ns/op, 2 allocs |

The rewrite is intentionally optimized first for the LSP hot path: already
canonical clean `file://` values, equality/map-key use, `String`, `File`, and
POSIX `FsPath`. Dirty non-file URIs still pay for vscode-compatible opaque
component encoding and canonical re-splitting; this is the main remaining
latency gap.

## Tuning decisions in this pass

- Added a dedicated canonical `file:///...` parser path. It validates clean path
  bytes, preserves the `file:////...` validation error path, constructs offsets
  directly, and avoids allocation.
- Added a clean absolute POSIX `FileFor` path. It returns a canonical file URI
  with one string allocation and no filesystem syscall or `filepath.Abs` work.
- Added a clean POSIX `FsPathFor` path. It returns the stored path substring for
  no-authority, no-escape, non-drive file URIs.
- Wrote encoded path/query/fragment/authority bytes directly into the final
  formatter builder, removing intermediate encoded strings on dirty slow paths.
- Removed duplicated raw splitting and `%` pre-scans from parse classification;
  canonical checks already reject `%` and the slow path now uses an
  `IndexByte`-guarded decoder.

## Inline audit

Command:

```sh
env GOEXPERIMENT= go test -run '^$' -gcflags='go.lsp.dev/uri=-m=2' .
```

Relevant compiler output from the current tree:

| Function | Inline status | Note |
| --- | --- | --- |
| `Parse` | can inline, cost 71 | Public wrapper stays cheap. |
| `ParseStrict` | can inline, cost 71 | Public wrapper stays cheap. |
| `URI.String` | can inline, cost 3 | Hot string path is effectively a field read. |
| `URI.FsPath` | can inline, cost 66 | Wrapper inlines; full `FsPathFor` remains out of line. |
| `File` | can inline, cost 65 | Wrapper inlines; full `FileFor` remains out of line. |
| `newEncodedURIWithRaw` | can inline, cost 80 | Fast offset construction is at the budget edge. |
| `parseCanonicalFileFast` | cannot inline, cost 146 | Kept as a focused fast-path helper; benchmark validates benefit. |
| `parseCanonicalFast` | cannot inline, cost 702 | Full generic predicate is intentionally larger and correctness-heavy. |
| `formatComponents` | cannot inline, cost 793 | Central serializer; dirty-path bottleneck. |
| `percentDecode` | cannot inline, cost 117 | Run-based vscode-compatible graceful decoder. |

## BCE audit

Command:

```sh
env GOEXPERIMENT= go test -run '^$' -gcflags='go.lsp.dev/uri=-d=ssa/check_bce/debug=1' .
```

The audit passes but still reports bounds checks in dense scanner/serializer
code. The highest-value remaining sites are:

| Area | Representative sites | Current stance |
| --- | --- | --- |
| Generic parse predicate | `uri.go:240`, `uri.go:303-321` | Accept for now; not on the clean file fast path. |
| Clean file and fsPath fast paths | `uri.go:260-268`, `fspath.go:115-119` | Benchmarks already meet targets; further BCE work is optional. |
| Percent decoder | `decode.go` | Iterative suffix-validity scan avoids recursive quadratic malformed-run behavior. |
| Formatter | `encode.go`, `format.go` slice/write sites | Dirty-path latency gap; next tuning target if HTTP latency matters. |
| Node path helpers | `posixpath.go`, `utils.go` | Cold relative to parse/fsPath workloads. |

## Remaining performance gaps

1. Dirty HTTP(S) parse latency is still above the aspirational target
   (157.0 ns/op measured vs <= 80 ns/op). Allocation is already below target.
   Closing the latency gap likely requires a formatter that computes final
   offsets while serializing, avoiding the post-format `splitRaw`, and possibly
   a specialized no-percent dirty-query path.
2. Percent-encoded file paths and Windows/UNC filesystem conversion still
   allocate because they must decode, lowercase drives, prepend UNC authority,
   or replace slashes. These paths are correct but not yet at the clean POSIX
   hot-path cost.
3. Malformed percent-run decoding is now linear, but still allocates temporary
   byte/validity buffers for correctness and simplicity on the rare invalid path.
4. There is no committed old-package or goada benchmark dependency. The suite
   intentionally keeps runtime dependencies at zero and uses `net/url` baselines
   plus documented commands as comparison scaffolding.
