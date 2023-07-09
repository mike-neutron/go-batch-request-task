package src

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SetStockService struct {
	m sync.Mutex
	Limit Limit
}

type Limit struct {
	Count int
	until time.Time
}

func (s *SetStockService) GetLimits() {
	s.Limit = Limit{Count: 100, until: time.Now().Add(time.Hour)}
}

func (s *SetStockService) Process(ctx context.Context, batch Batch) error {
	s.GetLimits()
	wg := &sync.WaitGroup{}
	for _, item := range batch {
		wg.Add(1)
		go s.processItem(wg, item)
	}
	wg.Wait()
	fmt.Println(s.Limit.Count)
	return nil
}

func (s *SetStockService) processItem(wg *sync.WaitGroup, item Item) error {
	fmt.Println(item)
	s.m.Lock()
	s.Limit.Count -= 1
	s.m.Unlock()
	wg.Done()
	return nil
}
