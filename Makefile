
test:
	go test -race ./...
.PHONY: test

static-analysis:
	bash -c 'diff -u <(echo -n) <(gofmt -s -d .)'
	go vet ./...
	gosec ./...
	nancy go.sum
.PHONY: static-analysis

t: static-analysis test
.PHONY: test-all

generate:
	go generate ./...
.PHONY: generate
