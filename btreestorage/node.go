package btreestorage

type (
	Node struct {
		ID    string      `json:"id"`
		Data  interface{} `json:"data"`
		Left  *Node       `json:"left"`
		Right *Node       `json:"right"`

		bf int
	}
)

func newNode(id string, data interface{}) *Node {
	n := Node{}
	n.ID = id
	n.Data = data
	return &n
}

func (n *Node) insert(id string, data interface{}) bool {
	switch {
	case id == n.ID:
		n.Data = data
		return false
	case id < n.ID:

		if n.Left == nil {
			n.Left = newNode(id, data)
			if n.Right == nil {
				n.bf = -1
			} else {
				n.bf = 0
			}
		} else {
			if n.Left.insert(id, data) {
				if !n.Left.isBalanced() {
					n.rebalance(n.Left)
				} else {
					n.bf = n.bf - 1
				}
			}
		}
	case id > n.ID:
		if n.Right == nil {
			n.Right = newNode(id, data)
			if n.Left == nil {
				n.bf = 1
			} else {
				n.bf = 0
			}
		} else {
			if n.Right.insert(id, data) {
				if !n.Right.isBalanced() {
					n.rebalance(n.Right)
				} else {
					n.bf = n.bf + 1
				}
			}
		}
	}

	if n.bf != 0 {
		return true
	}

	return false
}

func (n *Node) isBalanced() bool {
	return n.bf > -2 && n.bf < 2
}

func (n *Node) rebalance(c *Node) {
	switch {
	case c.bf == -2 && c.Left.bf == -1:
		n.rotateR(c)
	case c.bf == 2 && c.Right.bf == 1:
		n.rotateL(c)
	case c.bf == -2 && c.Left.bf == 1:
		n.rotateLR(c)
	case c.bf == 2 && c.Right.bf == -1:
		n.rotateRL(c)
	}
}

func (n *Node) rotateL(c *Node) {
	cr := c.Right
	c.Right = cr.Left
	cr.Left = c

	if c == n.Left {
		n.Left = cr
	} else {
		n.Right = cr
	}

	c.bf = 0
	cr.bf = 0
}

func (n *Node) rotateR(c *Node) {
	cl := c.Left
	c.Left = cl.Right
	cl.Right = c

	if c == n.Left {
		n.Left = cl
	} else {
		n.Right = cl
	}

	c.bf = 0
	cl.bf = 0
}

func (n *Node) rotateLR(c *Node) {
	c.Right.Left.bf = 1
	c.rotateR(c.Right)
	c.Right.bf = 1
	n.rotateL(c)
}

func (n *Node) rotateRL(c *Node) {
	c.Left.Right.bf = -1
	c.rotateL(c.Left)
	c.Left.bf = -1
	n.rotateR(c)
}

func (n *Node) search(key string) (interface{}, bool) {
	if n == nil {
		return nil, false
	}

	switch {
	case key == n.ID:
		return n.Data, true
	case key < n.ID:
		return n.Left.search(key)
	case key > n.ID:
		return n.Right.search(key)
	}

	return nil, false
}
