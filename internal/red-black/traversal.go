package rb

import "golang.org/x/exp/constraints"

func inorderTraversal[K constraints.Ordered, V any](n *Node[K, V], pairs []Pair[K, V]) []Pair[K, V] {
	if n == nil {
		return pairs
	}

	pairs = inorderTraversal(n.left, pairs)
	pairs = append(pairs, n.Pair)
	return inorderTraversal(n.right, pairs)
}
