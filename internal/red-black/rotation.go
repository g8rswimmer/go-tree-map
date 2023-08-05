package rb

import "golang.org/x/exp/constraints"

type rotation[K constraints.Ordered, V any] struct {
	left      bool
	right     bool
	leftRight bool
	rightLeft bool
}

func (r *rotation[K, V]) rotate(n *Node[K, V]) *Node[K, V] {
	switch {
	case r.left:
		n = rotateLeft(n)
		n.black = true
		n.left.black = false
		r.left = false
	case r.right:
		n = rotateRight(n)
		n.black = true
		n.right.black = false
		r.right = false
	case r.rightLeft:
		n.right = rotateRight(n.right)
		n.right.parent = n
		n = rotateLeft(n)
		n.black = true
		n.left.black = false
		r.rightLeft = false
	case r.leftRight:
		n.left = rotateLeft(n.left)
		n.left.parent = n
		n = rotateRight(n)
		n.black = true
		n.right.black = false
		r.leftRight = false
	default:
	}
	return n
}
