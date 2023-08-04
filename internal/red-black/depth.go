package rb

func depth(n *Node, curr int) int {
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
