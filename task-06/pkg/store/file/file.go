package file

import (
	"dersnek/task-6/pkg/index"
	"dersnek/task-6/pkg/store"
	"encoding/gob"
	"os"
)

// Service is a file store service.
type Service struct{}

// IndexStore - format in which the index and docs are serialized
// to be stored in a file later
type IndexStore struct {
	Docs  []store.Doc
	Index index.Data
}

// New creates and returns a new file store service.
func New() *Service {
	s := Service{}
	return &s
}

const filePath = "./cache.gob"

// Write writes serialized documents and index to a file
func (s *Service) Write(docs []store.Doc, indx index.Data) error {
	is := IndexStore{
		Docs:  docs,
		Index: indx,
	}
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(is)
	if err != nil {
		return err
	}
	return nil
}

// Read returns deserialized documents and index read from files
func (s *Service) Read() ([]store.Doc, index.Data, error) {
	is := IndexStore{}

	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			return is.Docs, is.Index, nil
		}
		return is.Docs, is.Index, err
	}
	defer file.Close()
	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&is)
	if err != nil {
		return is.Docs, is.Index, err
	}
	return is.Docs, is.Index, nil
}
