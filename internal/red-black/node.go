package rb

type Pair struct {
	Key   int
	Value string
}

type Node struct {
	Pair   Pair
	left   *Node
	right  *Node
	parent *Node
	black  bool
}
