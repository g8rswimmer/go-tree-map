package rb

import "golang.org/x/exp/constraints"

func insert[K constraints.Ordered, V any](root, n *Node[K, V], p Pair[K, V], r *rotation[K, V]) *Node[K, V] {
	rrConflict := false
	switch {
	case n == nil:
		return &Node[K, V]{
			Pair:  p,
			black: false,
		}
	case p.Key == n.Pair.Key:
		n.Pair.Value = p.Value
		return n
	case p.Key < n.Pair.Key:
		n.left = insert(root, n.left, p, r)
		n.left.parent = n
		if root.Pair.Key != n.Pair.Key {
			if !n.black && !n.left.black {
				rrConflict = true
			}
		}
	default:
		n.right = insert(root, n.right, p, r)
		n.right.parent = n
		if root.Pair.Key != n.Pair.Key {
			if !n.black && !n.right.black {
				rrConflict = true
			}
		}
	}
	n = r.rotate(n)

	if rrConflict {
		handleRedRedConflict(root, n, r)
	}
	return n
}

func handleRedRedConflict[K constraints.Ordered, V any](root, n *Node[K, V], r *rotation[K, V]) {
	if n.parent.right == n {
		if n.parent.left == nil || n.parent.left.black {
			switch {
			case n.left != nil && !n.left.black:
				r.rightLeft = true
			case n.right != nil && !n.right.black:
				r.left = true
			default:
			}
			return
		}
		n.parent.left.black = true
		n.black = true
		if root != n.parent {
			n.parent.black = false
		}
		return
	}

	if n.parent.right == nil || n.parent.right.black {
		switch {
		case n.left != nil && !n.left.black:
			r.right = true
		case n.right != nil && !n.right.black:
			r.leftRight = true
		default:
		}
		return
	}

	n.parent.right.black = true
	n.black = true
	if root != n.parent {
		n.parent.black = false
	}

}
