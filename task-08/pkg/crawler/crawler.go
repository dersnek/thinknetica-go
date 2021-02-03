package crawler

import "dersnek/task-8/pkg/store"

// Interface определяет контракт поискового робота.
type Interface interface {
	Scan(url string, depth int) ([]store.Doc, error)
}
