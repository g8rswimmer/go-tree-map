package rb

func rotateLeft(n *Node) *Node {
	r := n.right
	r.left, n.right = n, r.left
	n.parent = r
	if n.right != nil {
		n.right.parent = n
	}
	return r
}

func rotateRight(n *Node) *Node {
	r := n.left
	r.right, n.left = n, r.right
	n.parent = r
	if n.left != nil {
		n.left.parent = n
	}
	return r
}
