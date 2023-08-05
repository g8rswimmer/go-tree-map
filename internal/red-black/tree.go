package rb

import "golang.org/x/exp/constraints"

type Tree[K constraints.Ordered, V any] struct {
	root *Node[K, V]
}

func (t *Tree[K, V]) Insert(p Pair[K, V]) {
	if t.root == nil {
		t.root = &Node[K, V]{
			Pair:  p,
			black: true,
		}
		return
	}

	t.root = insert(t.root, t.root, p, &rotation[K, V]{})
}

func (t *Tree[K, V]) Inorder() []Pair[K, V] {
	return inorderTraversal(t.root, []Pair[K, V]{})
}

func (t *Tree[K, V]) Search(k K) (V, bool) {
	return search(t.root, k)
}
