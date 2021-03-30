package counter

import (
"context"
"counter-service/internal"
)

type Service interface {
	GetCounter(ctx context.Context) (internal.Counter)
	IncrementCounter(ctx context.Context) internal.Counter
	DecrementCounter(ctx context.Context) internal.Counter
	ResetCounter(ctx context.Context) internal.Counter
	ServiceStatus(ctx context.Context) int
}