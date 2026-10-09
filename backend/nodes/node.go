package nodes

type node struct {
	body     string
	parents  []node
	children []node
}

func (n node) getBody() string {
	return n.body
}
