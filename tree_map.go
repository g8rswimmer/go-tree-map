package tree

import (
	"errors"
	"io"

	rb "github.com/g8rswimmer/go-tree-map/internal/red-black"
)

type Map struct {
	rbt  *rb.Tree
	size int
}

func NewMap() *Map {
	return &Map{
		rbt: &rb.Tree{},
	}
}

func (m *Map) Put(k int, v string) {
	m.rbt.Insert(rb.Pair{Key: k, Value: v})
	m.size++
}

func (m *Map) Get(k int) (string, bool) {
	return m.rbt.Search(k)
}

func (m *Map) Len() int {
	return m.size
}

func (m *Map) Iterate(handler func(k int, v string) error) error {
	pairs := m.rbt.Inorder()

	for _, p := range pairs {
		err := handler(p.Key, p.Value)
		switch {
		case errors.Is(err, io.EOF):
			return nil
		case err != nil:
			return err
		default:
		}
	}
	return nil
}
