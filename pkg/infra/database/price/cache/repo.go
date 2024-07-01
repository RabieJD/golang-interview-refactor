package cache

import (
	"errors"
	"interview/pkg/infra/database/price"
	"sync"
)

type repo struct {
	mx     sync.RWMutex
	prices map[string]float64
}

func NewRepo() price.Handler {
	return &repo{prices: itemPriceMapping}
}

func (r *repo) GetPrice(identifier string) (float64, error) {
	r.mx.Lock()
	defer r.mx.Unlock()
	if val, ok := r.prices[identifier]; ok {
		return val, nil
	}
	return 0, errors.New("not found")
}
