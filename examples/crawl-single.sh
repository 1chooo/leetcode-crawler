#!/usr/bin/env bash
set -euo pipefail
ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUT="${ROOT}/examples/out/single"
mkdir -p "${OUT}"
cd "${ROOT}"
go run ./cmd/leetcode-crawler crawl --problem 1 --lang python3 --path "${OUT}" --naming snake_case
