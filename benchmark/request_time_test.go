package benchmark

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"testing"

	"github.com/rakyll/hey/requester"
	"go.mondoo.com/ranger-rpc/benchmark/rangerbench"
	"go.mondoo.com/ranger-rpc/benchmark/sample"
	twirpbench "go.mondoo.com/ranger-rpc/benchmark/twirpbench"
	"google.golang.org/protobuf/proto"
)

var payload []byte

func init() {
	var err error
	sr := &rangerbench.SmallQuery{Id: sample.Id, Name: sample.Name, Message: sample.Message}
	payload, err = proto.Marshal(sr)
	if err != nil {
		panic(err)
	}
}

func doHeyReq(req *http.Request, body []byte) {
	w := &requester.Work{
		Request:            req,
		RequestBody:        body,
		N:                  200,
		C:                  50,
		QPS:                0,
		Timeout:            20,
		DisableCompression: false,
		DisableKeepAlives:  false,
		DisableRedirects:   false,
		H2:                 false,
		ProxyAddr:          nil,
		Output:             "",
	}
	w.Init()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		w.Stop()
	}()
	w.Run()
}

func TestLoad(t *testing.T) {

	benchmarks := []struct {
		name        string
		port        int
		clienturl   string
		contentType string
		payload     []byte
	}{
		{"twrip http protobuf", twirpbench.Serve(0, false), "http://localhost:%d/twirp/twirp.BenchmarkService/RpcEmpty", "application/protobuf", []byte{}},
		{"twirp http json", twirpbench.Serve(0, false), "http://localhost:%d/twirp/twirp.BenchmarkService/RpcEmpty", "application/json", []byte("{}")},
		{"ranger http protobuf", rangerbench.Serve(0, false), "http://localhost:%d/BenchmarkService/RpcEmpty", "application/protobuf", []byte{}},
	}

	for _, bm := range benchmarks {
		t.Run(bm.name+" roundtrip", func(t *testing.T) {
			serverAddr := fmt.Sprintf(bm.clienturl, bm.port)

			body := bm.payload
			req, err := http.NewRequest("POST", serverAddr, nil)
			if err != nil {
				t.Fatalf("could not create http method")
			}

			header := make(http.Header)
			header.Set("Content-Type", bm.contentType)
			req.Header = header
			req.ContentLength = int64(len(body))
			doHeyReq(req, body)
		})
	}
}
