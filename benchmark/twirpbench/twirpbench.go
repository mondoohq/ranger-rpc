package twirpbench

import (
	context "context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"

	"go.mondoo.com/ranger-rpc/benchmark/sample"
)

//go:generate protoc --go_out=. --go_opt=paths=source_relative --twirp_out=paths=source_relative:. twirpbench.proto

type BenchmarkServer struct{}

func (b *BenchmarkServer) RpcSmall(ctx context.Context, in *SmallQuery) (*SmallResponse, error) {
	return &SmallResponse{Id: in.Id, Message: in.Message, Name: in.Name}, nil
}
func (b *BenchmarkServer) RpcEmpty(ctx context.Context, in *Empty) (*DefaultResponse, error) {
	return &DefaultResponse{Message: sample.Message}, nil
}

func Serve(port int, ssl bool) int {
	// init implementation
	s := NewBenchmarkServiceServer(&BenchmarkServer{}, nil)

	var lis net.Listener
	var err error
	if ssl == false {
		lis, err = net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
		if err != nil {
			panic(err)
		}
	} else {
		cert, err := tls.X509KeyPair(sample.LocalhostCert, sample.LocalhostKey)
		if err != nil {
			panic(fmt.Sprintf("httptest: NewTLSServer: %v", err))
		}
		config := &tls.Config{Certificates: []tls.Certificate{cert}}

		lis, err = tls.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port), config)
		if err != nil {
			panic(err)
		}
	}

	go func() {
		err = http.Serve(lis, s)
		if err != nil {
			panic(err)
		}
	}()

	return lis.Addr().(*net.TCPAddr).Port
}
