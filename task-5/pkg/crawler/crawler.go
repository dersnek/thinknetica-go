package crawler

import "dersnek/task-5/pkg/w"

// Interface определяет контракт поискового робота.
type Interface interface {
	Scan(url string, depth int) ([]w.Doc, error)
}
