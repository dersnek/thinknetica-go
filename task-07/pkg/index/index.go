package index

import (
	"dersnek/task-7/pkg/store"
)

// Data contains a hashmap with words as keys and as values -
// an array of IDs of documents (webpages) containing given word.
type Data map[string][]uint

// Interface определяет контракт службы индексирования документов.
type Interface interface {
	Build([]store.Doc, store.Interface) []store.Doc
	Restore([]store.Doc, Data, store.Interface)
	Find(string) ([]uint, bool)
	Index() Data
}
