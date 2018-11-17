package merkletree

type MerkleTree struct {
	nodes *[]Nodes
}

type Node struct {
	value []byte
}

type Nodes struct {
	node *[]Node
}

type Params struct {
	level *uint64
}

type Elements struct {
	Elements []Element
}

type ElementMethodos interface {
	Equal(el Element) (bool)
	Hash() ([]byte)
	Concat(empty Element, emptySecond Element) (Element)
}

type Element struct {
	Element string

	ElementMethodos
}

func Tree(params *Params, elements *Elements) *MerkleTree {

	empty := new(Element)
	empty.Element = "|"

	nodes := new(Nodes)
	nodes.node = makeNodes(elements, empty)

	nodesGroup := make([]Nodes, 0)
	nodesGroup = append(nodesGroup, *nodes)
	nodesGroup = append(nodesGroup, *nodes)

	tree := new(MerkleTree)
	tree.nodes = &nodesGroup

	return tree
}

func makeNodes(elements *Elements, empty *Element) *[]Node {

	nodeGroup := make([]Node, 0)

	e := *elements

	for i := 0; i < len(e.Elements); i++ {
		n := new(Node)
		n.value = e.Elements[i].Hash()
		nodeGroup = append(nodeGroup, *n)
	}

	if len(nodeGroup)%2 != 0 {
		n := new(Node)
		emp := makeEmptyElement(*empty)
		n.value = emp.Hash()
		nodeGroup = append(nodeGroup, *n)
	}

	return &nodeGroup
}

func makeEmptyElement(empty Element) Element {
	empty = empty.Concat(empty, empty)
	return empty
}
