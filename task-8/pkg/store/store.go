package store

// Doc - a webpage. Does not have a field for body.
type Doc struct {
	ID    uint
	URL   string
	Title string
}

// Interface is an interface for stores: file store, memstore, db
type Interface interface {
	Find(uint) *Doc
	Add(*Doc, bool) uint
	Read() ([]Doc, interface{}, error)
	Write([]Doc, interface{}) error
}
