.PHONY: install
install:
	go install ./protoc-gen-rangerrpc
	go install ./protoc-gen-rangerrpc-swagger

prep:
	go install honnef.co/go/tools/cmd/staticcheck@latest

build/snapshot:
	goreleaser release --snapshot --skip-publish --rm-dist

.PHONY: generate/examples
generate/examples:
	go generate ./examples/pingpong
	go generate ./examples/oneof
	go generate ./examples/swagger
	go generate ./examples/rangerguard

.PHONY: run/example/server
run/example/server: install generate/examples
	go run examples/pingpong/server/main.go

.PHONY: run/example/client
run/example/client: generate/examples
	go run examples/pingpong/client/main.go

.PHONY: test
test:
	go test ./...
	go vet ./...
	staticcheck ./...

# Copywrite Check Tool: https://github.com/hashicorp/copywrite
license: license/headers/check

license/headers/check:
	copywrite headers --plan

license/headers/apply:
	copywrite headers
