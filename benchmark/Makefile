
prep:
	command -v protoc-gen-go || go install github.com/golang/protobuf/protoc-gen-go
	command -v protoc-gen-go-grpc || go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	command -v protoc-gen-twirp || go install github.com/twitchtv/twirp/protoc-gen-twirp
	command -v hey || go install github.com/rakyll/hey
	go generate ./twirpbench
	go generate ./grpcbench
	go generate ./rangerbench

.PHONY: bench
bench:
	# NOTE: if you see too many open files error, set `ulimit -n 10240` on macOS
	go test -v -benchmem -benchtime 5s -bench . # -cpuprofile=cpu.out -memprofile=mem.out -trace trace.out 