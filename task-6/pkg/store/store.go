package store

// Doc - a webpage. Does not have a field for body.
type Doc struct {
	ID    uint
	URL   string
	Title string
}

// Store is an interface for stores: file store, memstore, db
type Store interface {
}
