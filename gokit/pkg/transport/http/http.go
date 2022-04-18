package http

import (
	"context"
	"encoding/json"
	"gokit/pkg/endpoint"
	"gokit/pkg/service"
	"gokit/pkg/service/middleware"
	"time"

	"net/http"

	kep "github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"

	httptransport "github.com/go-kit/kit/transport/http"
)

const rateBucketNum = 20

func InitHttp() {
	svc := service.NewCaculate()
	var sum kep.Endpoint
	sum = endpoint.MakeSumEndpoint(svc)
	sum = middleware.TokenBucketLimitter(rate.NewLimiter(rate.Every(time.Second*1), rateBucketNum))(sum)

	uppercaseHandler := httptransport.NewServer(
		sum,
		decodeUppercaseRequest,
		encodeResponse,
	)

	// countHandler := httptransport.NewServer(
	// 	makeCountEndpoint(svc),
	// 	decodeCountRequest,
	// 	encodeResponse,
	// )

	http.Handle("/uppercase", uppercaseHandler)
	// http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request endpoint.SumRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

// func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
// 	var request countRequest
// 	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
// 		return nil, err
// 	}
// 	return request, nil
// }

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
