package membot

import "dersnek/task-8/pkg/store"

// Service - имитация служба поискового робота.
type Service struct{}

// New - констрктор имитации службы поискового робота.
func New() *Service {
	s := Service{}
	return &s
}

// Scan возвращает заранее подготовленный набор данных
func (s *Service) Scan(url string, depth int) ([]store.Doc, error) {

	data := []store.Doc{
		{
			ID:    0,
			URL:   "https://yandex.ru",
			Title: "Яндекс",
		},
		{
			ID:    1,
			URL:   "https://google.ru",
			Title: "Google",
		},
	}

	return data, nil
}
