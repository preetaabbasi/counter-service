package endpoints

import (
	"context"
	"counter-service/pkg/counter"
	"errors"
	"os"

	"counter-service/internal"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type Set struct {
	GetCounterEndpoint    endpoint.Endpoint
	IncrementEndpoint     endpoint.Endpoint
	DecrementEndpoint     endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	ResetCounterEndpoint  endpoint.Endpoint
}

func NewEndpointSet(svc counter.Service) Set {
	return Set{
		GetCounterEndpoint:    MakeGetEndpoint(svc),
		IncrementEndpoint:     MakeIncrementEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
		ResetCounterEndpoint:  MakeResetEndpoint(svc),
		DecrementEndpoint: 	   MakeDecrementEndpoint(svc),

	}
}

func MakeDecrementEndpoint(svc counter.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		counter := svc.DecrementCounter(ctx)
		return CounterServiceResponse{Counter: counter}, nil
	}
}

func MakeGetEndpoint(svc counter.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		counter := svc.GetCounter(ctx)
		return CounterServiceResponse{counter}, nil
	}
}

func (s *Set) GetCounter(ctx context.Context) (internal.Counter) {
	resp, _ := s.GetCounterEndpoint(ctx, "get")
	getResp := resp.(CounterServiceResponse)
	return getResp.Counter
}


func MakeIncrementEndpoint(svc counter.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		counter := svc.IncrementCounter(ctx)
		return CounterServiceResponse{Counter: counter}, nil
	}
}

func (s *Set) IncrementCounter(ctx context.Context) (internal.Counter) {
	resp, _ := s.IncrementEndpoint(ctx,"increment")
	adResp := resp.(CounterServiceResponse)
	return adResp.Counter
}

func MakeResetEndpoint(svc counter.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		counter := svc.ResetCounter(ctx)
		return CounterServiceResponse{Counter: counter}, nil
	}
}

func (s *Set) ResetCounter(ctx context.Context) (internal.Counter) {
	resp, _ := s.ResetCounterEndpoint(ctx,"reset")
	adResp := resp.(CounterServiceResponse)
	return adResp.Counter
}

func MakeServiceStatusEndpoint(svc counter.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(ServiceStatusRequest)
		code := svc.ServiceStatus(ctx)
		return ServiceStatusResponse{Code: code, Err: ""}, nil
	}
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})
	svcStatusResp := resp.(ServiceStatusResponse)
	if err != nil {
		return svcStatusResp.Code, err
	}
	if svcStatusResp.Err != "" {
		return svcStatusResp.Code, errors.New(svcStatusResp.Err)
	}
	return svcStatusResp.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}