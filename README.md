## CounterService

CounterService is created by using go-kit library which is used to create microservices in Go.

1. This simple service has a in-memory stateful counter. There are five endpoints exposed by this service.
```
GetCounter:         GET /counter  

IncrementCounter:   POST /counter/increment

DecrementCounter:   POST /counter/decrement

ResetCounter:       POST /counter/reset

Health check:       GET /health

The response returned by endpoints is
```
The example response returned by counter endpoints.
```json
{"Counter":{"value":2,"desc":"current value of the counter"}}
```
To run the application, 
```bash
cd counter-service/cmd/counter
go run main.go
```

To test the application, either 
1. run main_test.go
The main_test.go uses resty library to create the REST client tests to call the API. 
```bash
cd counter-service/cmd/counter
go test -v main_test.go
```
Also you can import the json collection which is on the root path of the project to execute the REST API
endpoints in Postman.

Application endpoints are exposed on port 8081. Press Ctrl C to send the signal to stop the application.


