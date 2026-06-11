#!/usr/bin/env node
import { readFileSync, writeFileSync } from 'node:fs';
import { createRequire } from 'node:module';

const require = createRequire(import.meta.url);
const out = new URL('../../testdata/vectors.json', import.meta.url);
const packageFile = new URL('./package.json', import.meta.url);
const useStaticFixture = process.argv.includes('--use-static-fixture');

function readExistingVectors() {
  return JSON.parse(readFileSync(out, 'utf8'));
}

function pinnedVscodeURIVersion() {
  const manifest = JSON.parse(readFileSync(packageFile, 'utf8'));
  return manifest.dependencies?.['vscode-uri'] ?? 'unknown';
}

function staticFixture() {
  const existing = readExistingVectors();
  existing.generatedAt = new Date(0).toISOString();
  existing.generator = 'static-fixture';
  existing.note = 'Explicit --use-static-fixture mode preserved existing checked-in vectors without consulting vscode-uri.';
  return existing;
}

function referencePayload() {
  const { URI, Utils } = require('vscode-uri');
  const base = readExistingVectors();
  const parseInputs = [
    ...base.parse.map((v) => v.input),
    'foo:?',
    'foo:#',
    'foo:?x#',
    'https://host/p?#',
    'file:///C:/test',
    'file:///c%3A/test',
    'file://SERVER/Share/X.go',
  ];
  const uniqueParseInputs = [...new Set(parseInputs)];

  const payload = {
    ...base,
    generator: 'vscode-uri-canonical-reparse',
    vscodeURIVersion: pinnedVscodeURIVersion(),
    generatedAt: new Date(0).toISOString(),
    contract: 'go-comparable-canonical-uri',
    referenceGenerated: [
      'parse.string',
      'parse.stringNoEncoding',
      'parse.components.fromCanonicalReparse',
      'parse.fsPathPOSIX.fromCanonicalReparse',
      'parse.fsPathWindows.fromCanonicalReparse',
      'paths',
    ],
    curated: ['errors'],
    note:
      'Go URI values compare by canonical string identity. Parse vectors derive component and fsPath fields by reparsing vscode-uri toString() output, not by preserving original parse-history casing.',
  };
  payload.parse = uniqueParseInputs.map((input) => {
    const canonical = URI.parse(URI.parse(input).toString());
    const existing = base.parse.find((v) => v.input === input);
    return {
      name: existing?.name ?? input,
      input,
      components: {
        scheme: canonical.scheme,
        authority: canonical.authority,
        path: canonical.path,
        query: canonical.query,
        fragment: canonical.fragment,
      },
      string: canonical.toString(),
      stringNoEncoding: canonical.toString(true),
      fsPathPOSIX: fsPathFor(canonical, false),
      fsPathWindows: fsPathFor(canonical, true),
    };
  });
  payload.paths = base.paths.map((v) => {
    const u = URI.parse(v.uri);
    let got;
    switch (v.op) {
      case 'join':
        got = Utils.joinPath(u, ...(v.segments ?? [])).toString();
        break;
      case 'resolve':
        got = Utils.resolvePath(u, ...(v.segments ?? [])).toString();
        break;
      case 'dirname':
        got = Utils.dirname(u).toString();
        break;
      case 'basename':
        got = Utils.basename(u);
        break;
      case 'extname':
        got = Utils.extname(u);
        break;
      default:
        throw new Error(`unknown path op ${v.op}`);
    }
    return { ...v, want: got };
  });
  return payload;
}

function fsPathFor(u, windows) {
  let value;
  if (u.scheme === 'file' && u.authority && u.path.length > 1) {
    value = `//${u.authority}${u.path}`;
  } else if (/^\/[A-Za-z]:/.test(u.path)) {
    value = `${u.path[1].toLowerCase()}${u.path.slice(2)}`;
  } else {
    value = u.path;
  }
  return windows ? value.replaceAll('/', '\\') : value;
}

let payload;
if (useStaticFixture) {
  payload = staticFixture();
} else {
  try {
    payload = referencePayload();
  } catch (err) {
    console.error('tools/genvectors requires the pinned vscode-uri dependency. Run `npm ci --prefix tools/genvectors` before generating vectors.');
    throw err;
  }
}

writeFileSync(out, `${JSON.stringify(payload, null, 2)}\n`);
