package transport

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"

	"counter-service/pkg/counter/endpoints"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
)

// NewService wires Go kit endpoints to the HTTP transport.
func NewHTTPHandler(ep endpoints.Set) http.Handler {
	// set-up router and initialize http endpoints
	router := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	// HTTP Post - /increment
	router.Methods("POST").Path("/counter/increment").Handler(kithttp.NewServer(
		ep.IncrementEndpoint,
		decodeHTTPPostRequest,
		encodeResponse,
		options...,
	))

	// HTTP Post - /decrement
	router.Methods("POST").Path("/counter/decrement").Handler(kithttp.NewServer(
		ep.DecrementEndpoint,
		decodeHTTPPostRequest,
		encodeResponse,
		options...,
	))

	// HTTP GET - get counter value
	router.Methods("GET").Path("/counter").Handler(kithttp.NewServer(
		ep.GetCounterEndpoint,
		decodeHTTPGetRequest,
		encodeResponse,
		options...,
	))
	// counter service health check
	router.Methods("GET").Path("/health").Handler(kithttp.NewServer(
		ep.ServiceStatusEndpoint,
		decodeHTTPServiceStatusRequest,
		encodeResponse,
	))
	// HTTP Post - /increment
	router.Methods("POST").Path("/counter/reset").Handler(kithttp.NewServer(
		ep.ResetCounterEndpoint,
		decodeHTTPPostRequest,
		encodeResponse,
		options...,
	))
	return router
}


func decodeHTTPGetRequest(_ context.Context, r *http.Request) (interface{}, error) {
	//var req endpoints.GetRequest
	if reqHeadersBytes, err := json.Marshal(r.Header); err != nil {
		fmt.Println("Could not Marshal Req Headers")
	} else {
		fmt.Println(string(reqHeadersBytes))
	}
	return nil, nil
}



func decodeHTTPPostRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if r.ContentLength == 0 {
		logger.Log("POST request with no body ;)")
		return nil, nil
	}
	return nil, nil
}

func decodeHTTPServiceStatusRequest(_ context.Context, _ *http.Request) (interface{}, error) {
	var req endpoints.ServiceStatusRequest
	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(error); ok && e != nil {
		encodeError(ctx, e, w)
		return nil
	}
	//w.WriteHeader()
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
