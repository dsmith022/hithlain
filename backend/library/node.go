package library

type node struct {
	depth    int
	body     string
	parents  []node
	children []node
}

func (n node) getBody() string {
	return n.body
}

func (n node) getParents() []node {
	return n.parents
}

func (n node) getChildren() []node {
	return n.children
}
