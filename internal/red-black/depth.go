package rb

import "golang.org/x/exp/constraints"

func depth[K constraints.Ordered, V any](n *Node[K, V], curr int) int {
	if n == nil {
		return curr
	}
	curr++
	l := depth(n.left, curr)
	r := depth(n.right, curr)

	if l > r {
		return l
	}
	return r
}
