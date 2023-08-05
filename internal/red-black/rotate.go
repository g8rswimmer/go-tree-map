package rb

import "golang.org/x/exp/constraints"

func rotateLeft[K constraints.Ordered, V any](n *Node[K, V]) *Node[K, V] {
	r := n.right
	r.left, n.right = n, r.left
	n.parent = r
	if n.right != nil {
		n.right.parent = n
	}
	return r
}

func rotateRight[K constraints.Ordered, V any](n *Node[K, V]) *Node[K, V] {
	r := n.left
	r.right, n.left = n, r.right
	n.parent = r
	if n.left != nil {
		n.left.parent = n
	}
	return r
}
