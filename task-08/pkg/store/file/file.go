package file

import (
	"dersnek/task-8/pkg/index"
	"dersnek/task-8/pkg/store"
	"encoding/gob"
	"os"
)

// Service is a file store service.
type Service struct {
	tree *tree
}

// Tree - binary search tree
type tree struct {
	root      *Node
	currentID uint
}

// Node - a tree node
type Node struct {
	left, right *Node
	Value       *store.Doc
}

// IndexStore - format in which the index and docs are serialized
// to be stored in a file later
type IndexStore struct {
	Docs  []store.Doc
	Index index.Data
}

// New creates and returns a new file store service.
func New() *Service {
	s := Service{
		new(tree),
	}
	return &s
}

const filePath = "./cache.gob"

// Write writes serialized documents and index to a file
func (s *Service) Write(docs []store.Doc, ind interface{}) error {
	is := IndexStore{
		Docs:  docs,
		Index: ind.(index.Data),
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
func (s *Service) Read() ([]store.Doc, interface{}, error) {
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

// Add a web document to the store. Returns ID assigned to this document.
func (s *Service) Add(doc *store.Doc, assignID bool) uint {
	if assignID {
		doc.ID = s.tree.currentID
		s.tree.currentID++
	} else if !assignID && doc.ID > s.tree.currentID {
		s.tree.currentID = doc.ID
	}

	n := &Node{Value: doc}
	if s.tree.root == nil {
		s.tree.root = n
		return doc.ID
	}

	return s.insert(s.tree.root, n)
}

func (s *Service) insert(node, new *Node) uint {
	if new.Value.ID < node.Value.ID {
		if node.left == nil {
			node.left = new
			return new.Value.ID
		}

		return s.insert(node.left, new)
	}

	if node.right == nil {
		node.right = new
		return new.Value.ID
	}

	return s.insert(node.right, new)
}

// Find - find a web document in a tree by its ID
func (s *Service) Find(id uint) *store.Doc {
	return search(s.tree.root, id)
}

func search(n *Node, id uint) *store.Doc {
	if n == nil {
		return nil
	}

	if n.Value.ID == id {
		return n.Value
	}

	if n.Value.ID < id {
		return search(n.right, id)
	}

	return search(n.left, id)
}
