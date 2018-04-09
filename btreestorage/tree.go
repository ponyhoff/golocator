package btreestorage

type (
	BTree struct {
		Root *Node `json:"Root"`
	}
)

func (t *BTree) insert(id string, data interface{}) bool {
	if t.Root == nil {
		t.Root = newNode(id, data)
		return true
	}

	ok := t.Root.insert(id, data)
	if !t.isBalanced() {
		t.rebalance()
	}

	return ok
}

func (t *BTree) rebalance() {
	// the Root node of the tree does not have a parent.
	ghostNode := &Node{Left: t.Root, ID: "ghost-parent-node"}
	ghostNode.rebalance(t.Root)

	t.Root = ghostNode.Left
}

func (t *BTree) isBalanced() bool {
	return t.Root.isBalanced()
}

func (t *BTree) find(key string) (interface{}, bool) {
	if t.Root == nil {
		return nil, false
	}

	return t.Root.search(key)
}
