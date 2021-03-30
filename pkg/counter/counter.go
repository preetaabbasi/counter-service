package counter

import (
	"context"
	"net/http"
	"os"
	"sync/atomic"

	"counter-service/internal"

	"github.com/go-kit/kit/log"
)

type CounterValue int64
var counterValue CounterValue

func (counterValue *CounterValue) increment() int64 {
	return atomic.AddInt64((*int64)(counterValue), 1)
}

func (counterValue *CounterValue) get() int64 {
	return atomic.LoadInt64((*int64)(counterValue))
}

func (counterValue *CounterValue) decrement() int64 {
	return atomic.AddInt64((*int64)(counterValue), -1)
}

func (counterValue *CounterValue) resetCounter() int64 {
	return atomic.SwapInt64((*int64)(counterValue),0)
}

type counterService struct{}


func (c *counterService) GetCounter(_ context.Context) (internal.Counter) {
	counter := internal.Counter{
		Value: counterValue.get(),
		Desc:   "current value of the counter",
	}
	return counter
}


func (c *counterService) IncrementCounter(ctx context.Context) internal.Counter {
	counter := internal.Counter{
		Value: counterValue.increment(),
		Desc:   "current value of the counter",
	}
	return counter
}

func (c *counterService) DecrementCounter(ctx context.Context) internal.Counter {
	counter := internal.Counter{
		Value: counterValue.decrement(),
		Desc: "decremented value of the counter",
	}
	return counter
}

func (c *counterService) ServiceStatus(ctx context.Context) int {
	logger.Log("Checking the Service health...")
	return http.StatusOK
}

func (c *counterService) ResetCounter(ctx context.Context) internal.Counter {
	counter := internal.Counter{
		Value: counterValue.resetCounter(),
		Desc:   "reset value of the counter to 0",
	}
	return counter
}

func NewService() Service { return &counterService{} }

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}