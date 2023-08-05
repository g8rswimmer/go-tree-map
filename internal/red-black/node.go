package rb

type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

type Node struct {
	Pair   Pair[int, string]
	left   *Node
	right  *Node
	parent *Node
	black  bool
}
