VERSION ?= dev
LDFLAGS := -ldflags "-X github.com/1chooo/leetcode-crawler/cmd.version=$(VERSION)"

.PHONY: build install test

build:
	go build -o leetcode-crawler $(LDFLAGS) ./cmd/leetcode-crawler

install:
	go install $(LDFLAGS) ./cmd/leetcode-crawler

test:
	go test ./...
