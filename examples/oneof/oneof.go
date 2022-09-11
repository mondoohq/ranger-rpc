package oneof

// To regenerate the protocol buffer output for this package, run `go generate`
//go:generate protoc --go_out=. --go_opt=paths=source_relative --rangerrpc_out=. oneof.proto
//go:generate protoc --go_out=./client --go_opt=paths=source_relative --rangerrpc_out=./client --rangerrpc_opt=client-only=true oneof.proto
