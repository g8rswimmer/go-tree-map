package rb

func search(n *Node, k int) (string, bool) {
	switch {
	case n == nil:
		return "", false
	case n.Pair.Key == k:
		return n.Pair.Value, true
	case n.Pair.Key < k:
		return search(n.right, k)
	default:
		return search(n.left, k)
	}
}
