package rb

func inorderTraversal(n *Node, pairs []Pair) []Pair {
	if n == nil {
		return pairs
	}

	pairs = inorderTraversal(n.left, pairs)
	pairs = append(pairs, n.Pair)
	return inorderTraversal(n.right, pairs)
}
