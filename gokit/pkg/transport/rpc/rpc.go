package rpc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"gokit/pkg/endpoint"
	"gokit/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"net/http"

	kep "github.com/go-kit/kit/endpoint"

	jsonrpc "github.com/go-kit/kit/transport/http/jsonrpc"
)

const rateBucketNum = 20

func InitRPC() {
	svc := service.NewCaculate()
	var sum kep.Endpoint
	sum = endpoint.MakeSumEndpoint(svc)
	// sum = middleware.TokenBucketLimitter(rate.NewLimiter(rate.Every(time.Second*1), rateBucketNum))(sum)

	handler := jsonrpc.NewServer(jsonrpc.EndpointCodecMap{
		"sum": jsonrpc.EndpointCodec{
			Endpoint: sum,
			Decode:   decodeSumRequest,
			Encode:   encodeSumResponse,
		},
	})

	srv := &http.Server{
		Addr:    ":80",
		Handler: handler,
	}

	log.Println("start Server ...")
	_, cancel := context.WithCancel(context.Background())
	srv.RegisterOnShutdown(cancel)
	ch := make(chan os.Signal, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("SERVER err:", err)
		}
	}()

	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	//阻塞channel
	<-ch

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

}

/// server ///
func decodeSumRequest(ctx context.Context, msg json.RawMessage) (interface{}, error) {
	var req endpoint.SumRequest
	err := json.Unmarshal(msg, &req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func encodeSumResponse(ctx context.Context, result interface{}) (json.RawMessage, error) {
	fmt.Printf("SumResponse %v", result)
	sum, ok := result.(endpoint.SumResponse)
	if !ok {
		return nil, errors.New("result is not an int")
	}

	b, err := json.Marshal(sum)
	if err != nil {
		return nil, err
	}
	return b, nil
}

/// client ///

func DecodeSumResponse(_ context.Context, res jsonrpc.Response) (interface{}, error) {
	if res.Error != nil {
		return nil, *res.Error
	}
	var sumres endpoint.SumResponse
	err := json.Unmarshal(res.Result, &sumres)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshal body to SumResponse: %s", err)
	}
	return sumres, nil
}

func EncodeSumRequest(_ context.Context, obj interface{}) (json.RawMessage, error) {
	req, ok := obj.(endpoint.SumRequest)
	if !ok {
		return nil, fmt.Errorf("couldn't assert request as SumRequest, got %T", obj)
	}
	b, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal request: %s", err)
	}
	return b, nil
}
