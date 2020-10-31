package crawler

// Document - webpage, scanned by a crawler. Page body is not saved.
type Document struct {
	ID    uint
	URL   string
	Title string
}

// Interface определяет контракт поискового робота.
type Interface interface {
	Scan(url string, depth int) ([]Document, error)
}
