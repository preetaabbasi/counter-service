package endpoints

import "counter-service/internal"


type CounterServiceResponse struct {
	Counter internal.Counter
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
