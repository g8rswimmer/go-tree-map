package rb

type Tree struct {
	root *Node
}

func (t *Tree) Insert(p Pair[int, string]) {
	if t.root == nil {
		t.root = &Node{
			Pair:  p,
			black: true,
		}
		return
	}

	t.root = insert(t.root, t.root, p, &rotation{})
}

func (t *Tree) Inorder() []Pair[int, string] {
	return inorderTraversal(t.root, []Pair[int, string]{})
}

func (t *Tree) Search(k int) (string, bool) {
	return search(t.root, k)
}
