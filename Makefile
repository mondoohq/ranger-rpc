.PHONY: install
install:
	go install ./protoc-gen-rangerrpc
	go install ./protoc-gen-rangerrpc-swagger

.PHONY: generate/examples
generate/examples:
	go generate ./examples/pingpong
	go generate ./examples/oneof
	go generate ./examples/swagger

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