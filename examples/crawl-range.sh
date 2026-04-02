#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT="${ROOT}/examples/out/range"
mkdir -p "${OUT}"
cd "${ROOT}"
go run ./cmd/leetcode-crawler crawl --problem 1-2 --lang python3 --path "${OUT}" --naming kebab-case
