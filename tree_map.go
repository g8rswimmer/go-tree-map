/*
This is an implemenation of a tree map using a red-black tree.  This can be used when
ranging through the key value pairs needs to be deterministic.
*/
package tree

import (
	"errors"
	"io"

	rb "github.com/g8rswimmer/go-tree-map/internal/red-black"
	"golang.org/x/exp/constraints"
)

// Map is the tree map structure
type Map[K constraints.Ordered, V any] struct {
	rbt  *rb.Tree[K, V]
	size int
}

// NewMap is used to create a new tree map.  The keys must be
// an Ordered type and the value can be anything.
func NewMap[K constraints.Ordered, V any]() *Map[K, V] {
	return &Map[K, V]{
		rbt:  &rb.Tree[K, V]{},
		size: 0,
	}
}

// Store will store the value with the key.  If the
// key already has an value, the value passed will replace it.
func (m *Map[K, V]) Store(k K, v V) {
	m.rbt.Insert(rb.Pair[K, V]{Key: k, Value: v})
	m.size++
}

// Load will return the key's value.
func (m *Map[K, V]) Load(k K) (V, bool) {
	return m.rbt.Search(k)
}

// Len is the number of keys in the map.
func (m *Map[_, _]) Len() int {
	return m.size
}

// Range will pass the key value pair to the handler in the key's decending
// order.
//
// If the handler would like to break out of the range, return io.EOF.
func (m *Map[K, V]) Range(handler func(k K, v V) error) error {
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
