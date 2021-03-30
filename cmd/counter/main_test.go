package main

import (
	"counter-service/internal"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"
	"testing"
)

/*
Testing created endpoints on counter service
go test -v main_test.go
*/

// end to end flow to test the counter
func Test_API_All_Endpoints_Together(t *testing.T) {
	// at this point counter value should be zero as this is very first call
	responseGET := GetCounterAPICall()
	getCounterResponse := CounterDtoFromResponse(responseGET)
	assertOnResponse(t, responseGET, getCounterResponse, 0)

	// incrementing counter two times
	IncrementCounterAPICall(t)
	responseIncPOST := IncrementCounterAPICall(t)
	incrementResponse := CounterDtoFromResponse(responseIncPOST)
	// we expect counter value=2
	assertOnResponse(t, responseIncPOST, incrementResponse, 2)

	responseDecPOST := DecrementCounterAPICall(t)
	decrementResponse := CounterDtoFromResponse(responseDecPOST)
	// we expect counter value=1 after decrement
	assertOnResponse(t, responseDecPOST, decrementResponse, 1)

	// reset the counter so that next test case doesn't use the same counter
	ResetCounterAPICall(t)

}


func Test_GetCounter(t *testing.T) {
	response := GetCounterAPICall()
	getCounterResponse := CounterDtoFromResponse(response)
	assertOnResponse(t, response, getCounterResponse, 0)
}

func Test_Increment_Counter(t *testing.T) {
	response := IncrementCounterAPICall(t)
	incrementResponse := CounterDtoFromResponse(response)
	assertOnResponse(t, response, incrementResponse, 1)

	// reset the counter so that next test case doesn't use the same counter
	ResetCounterAPICall(t)
}

func Test_Decrement_Counter(t *testing.T) {
	response := DecrementCounterAPICall(t)
	incrementResponse := CounterDtoFromResponse(response)
	assertOnResponse(t, response, incrementResponse, -1)

	// reset the counter so that next test case doesn't use the same counter
	ResetCounterAPICall(t)
}

func GetCounterAPICall() *resty.Response{
	client := resty.New()
	resp, _ := client.R().Get("http://localhost:8081/counter")
	return resp
}

func IncrementCounterAPICall(t *testing.T) *resty.Response {
	client := resty.New()
	resp, _ := client.R().Post("http://localhost:8081/counter/increment")
	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}
	return resp
}

func DecrementCounterAPICall(t *testing.T) *resty.Response {
	client := resty.New()
	resp, _ := client.R().Post("http://localhost:8081/counter/decrement")
	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}
	return resp
}

func ResetCounterAPICall(t *testing.T) *resty.Response {
	client := resty.New()
	resp, _ := client.R().Post("http://localhost:8081/counter/reset")
	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}
	return resp
}

func assertOnResponse(t *testing.T, resp *resty.Response, counterServiceResponse CounterServiceResponse, counterValue int64) {
	assert.Equal(t, 200, resp.StatusCode())
	assert.Equal(t, counterValue, counterServiceResponse.Counter.Value)
}


func CounterDtoFromResponse(resp *resty.Response) CounterServiceResponse {
	var counterServiceResponse CounterServiceResponse
	json.Unmarshal([]byte(resp.String()), &counterServiceResponse)
	return counterServiceResponse
}

type CounterServiceResponse struct {
	Counter internal.Counter `json:"counter"`
}

