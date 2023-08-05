package tree

import (
	"errors"
	"io"

	rb "github.com/g8rswimmer/go-tree-map/internal/red-black"
	"golang.org/x/exp/constraints"
)

type Map[K constraints.Ordered, V any] struct {
	rbt  *rb.Tree[K, V]
	size int
}

func NewMap[K constraints.Ordered, V any]() *Map[K, V] {
	return &Map[K, V]{
		rbt: &rb.Tree[K, V]{},
	}
}

func (m *Map[K, V]) Put(k K, v V) {
	m.rbt.Insert(rb.Pair[K, V]{Key: k, Value: v})
	m.size++
}

func (m *Map[K, V]) Get(k K) (V, bool) {
	return m.rbt.Search(k)
}

func (m *Map[_, _]) Len() int {
	return m.size
}

func (m *Map[K, V]) Iterate(handler func(k K, v V) error) error {
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
