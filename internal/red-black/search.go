package rb

import "golang.org/x/exp/constraints"

func search[K constraints.Ordered, V any](n *Node[K, V], k K) (V, bool) {

	switch {
	case n == nil:
		var value V
		return value, false
	case n.Pair.Key == k:
		return n.Pair.Value, true
	case n.Pair.Key < k:
		return search(n.right, k)
	default:
		return search(n.left, k)
	}
}
