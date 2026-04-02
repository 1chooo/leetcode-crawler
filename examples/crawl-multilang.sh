#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT="${ROOT}/examples/out/multilang"
mkdir -p "${OUT}"
cd "${ROOT}"
go run ./cmd/leetcode-crawler crawl --problem 1 --lang python3,golang --path "${OUT}" --naming camelCase
