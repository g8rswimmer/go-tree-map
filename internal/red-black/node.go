package rb

import "golang.org/x/exp/constraints"

type Pair[K constraints.Ordered, V any] struct {
	Key   K
	Value V
}

type Node[K constraints.Ordered, V any] struct {
	Pair   Pair[K, V]
	left   *Node[K, V]
	right  *Node[K, V]
	parent *Node[K, V]
	black  bool
}
