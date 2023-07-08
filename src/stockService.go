package src

import (
	"sync"
	"time"
	"context"
)

type SetStockService struct {
	m sync.Mutex
	limits int64
}

func (s *SetStockService) GetLimits() (n int64, p time.Duration) {
	return int64(10000), time.Hour
}

func (s *SetStockService) Process(ctx context.Context, batch Batch) error {
	return nil
}
