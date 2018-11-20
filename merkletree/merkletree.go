package merkletree

type MerkleTree struct {
	nodes *[]Nodes
}

type Node struct {
	value []string
	ElementMethodos
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
	Equal(el Element) bool
	Hash() []byte
	Concat(empty Element, emptySecond Element) Element
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
	arr := e.Elements
	var temporaryArr [][2]string // array for nodes - [ [el1, el2] [el3, el4] ... ]

	if len(e.Elements)%2 != 0 {
		info := make(chan []Node)

		go func() {
			lastElements := arr[len(arr)-1]
			emp := makeEmptyElement(*empty)

			arr = arr[:len(arr)-1]

			for i := 0; i < len(arr); {
				twoEls := [2]string{arr[i].Element, arr[i+1].Element}
				temporaryArr = append(temporaryArr, twoEls)
				i += 2
			}
			temporaryArr = append(temporaryArr, [2]string{lastElements.Element, emp.Element}) // add to the end of array -> [...[lastElement, empty]]

			for i, j := range temporaryArr {
				n := new(Node)
				sum := j[i][0] + j[i][1] // [ [ el1, el2]...]; el1 + el2
				// keccak := sha3.NewKeccak256()
				// keccak.Write([]byte(string(sum)))
				n.value = []string{string(sum)}
				nodeGroup = append(nodeGroup, *n)
			}
			info <- nodeGroup
		}()

		result := <-info

		return &result

	}

	info := make(chan []Node)

	go func() {
		for i := 0; i < len(arr); {
			twoEls := [2]string{arr[i].Element, arr[i+1].Element}
			temporaryArr = append(temporaryArr, twoEls)
			i += 2
		}

		for i, j := range temporaryArr {
			n := new(Node)
			sum := j[i][0] + j[i][1]
			// keccak := sha3.NewKeccak256()
			// keccak.Write([]byte(string(sum)))
			n.value = []string{string(sum)}
			nodeGroup = append(nodeGroup, *n)
		}
		info <- nodeGroup
	}()

	result := <-info
	return &result

}

func makeEmptyElement(empty Element) Element {
	empty = empty.Concat(empty, empty)
	return empty
}
