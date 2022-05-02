package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"os"
	"time"

	caculate "gokit/pkg/endpoint"
	"gokit/pkg/transport/rpc"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http/jsonrpc"
)

var KIT_SERVER_URL string

func main() {
	// if !strings.HasPrefix(instance, "http") {
	// 	instance = "http://" + instance
	// }

	// Mongo_host
	KIT_SERVER_URL = os.Getenv("kit-server-url")
	if KIT_SERVER_URL == "" {
		log.Println("env: KIT_SERVER_URL is empty")
		KIT_SERVER_URL = "127.0.0.1"
	}

	u, err := url.Parse(fmt.Sprintf("http://%v/rpc", KIT_SERVER_URL))
	if err != nil {
		os.Exit(1)
	}

	// We construct a single ratelimiter middleware, to limit the total outgoing
	// QPS from this client to all methods on the remote instance. We also
	// construct per-endpoint circuitbreaker middlewares to demonstrate how
	// that's done, although they could easily be combined into a single breaker
	// for the entire remote instance, too.
	// limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))

	var sumEndpoint endpoint.Endpoint
	{
		sumEndpoint = jsonrpc.NewClient(
			u,
			"sum",
			jsonrpc.ClientRequestEncoder(rpc.EncodeSumRequest),
			jsonrpc.ClientResponseDecoder(rpc.DecodeSumResponse),
		).Endpoint()
	}

	svc := caculate.Caculate{
		SumEndpoint: sumEndpoint,
	}

	rand.Seed(time.Now().UnixNano())

	for {
		a, b := rand.Intn(100), rand.Intn(100)
		// a, b := 0, 0
		v, err := svc.Sum(context.Background(), a, b)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stdout, " %v %v + %v = %v\n", err, a, b, v)
		time.Sleep(time.Millisecond * 100)
	}

	// var concatEndpoint endpoint.Endpoint
	// {
	// 	concatEndpoint = jsonrpc.NewClient(
	// 		u,
	// 		"concat",
	// 		jsonrpc.ClientRequestEncoder(encodeConcatRequest),
	// 		jsonrpc.ClientResponseDecoder(decodeConcatResponse),
	// 	).Endpoint()
	// 	concatEndpoint = opentracing.TraceClient(tracer, "Concat")(concatEndpoint)
	// 	concatEndpoint = limiter(concatEndpoint)
	// 	concatEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
	// 		Name:    "Concat",
	// 		Timeout: 30 * time.Second,
	// 	}))(concatEndpoint)
	// }

	// Returning the endpoint.Set as a service.Service relies on the
	// endpoint.Set implementing the Service methods. That's just a simple bit
	// of glue code.
	// return addendpoint.Set{
	// 	SumEndpoint: sumEndpoint,
	// 	// ConcatEndpoint: concatEndpoint,
	// }, nil

}

// makeEndpointCodecMap returns a codec map configured for the addsvc.
// func makeEndpointCodecMap(endpoints addendpoint.Set) jsonrpc.EndpointCodecMap {
// 	return jsonrpc.EndpointCodecMap{
// 		"sum": jsonrpc.EndpointCodec{
// 			Endpoint: endpoints.SumEndpoint,
// 			Decode:   decodeSumRequest,
// 			Encode:   encodeSumResponse,
// 		},
// 		"concat": jsonrpc.EndpointCodec{
// 			Endpoint: endpoints.ConcatEndpoint,
// 			Decode:   decodeConcatRequest,
// 			Encode:   encodeConcatResponse,
// 		},
// 	}
// }
