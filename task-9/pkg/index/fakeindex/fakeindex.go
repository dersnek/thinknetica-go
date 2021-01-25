// Package fakeindex is a fake index
package fakeindex

import (
	"go-search/pkg/index"
)

// Service - search index service
type Service struct {
	index index.Data
	index.Interface
}

// New - creates and returns a new index service
func New() *Service {
	return &Service{}
}

// Find is a fake method which returns predetermined values
func (s *Service) Find(word string) ([]uint, bool) {
	switch word {
	case "a":
		return []uint{0}, true
	case "b":
		return []uint{0, 1}, true
	default:
		return []uint{}, false
	}
}
