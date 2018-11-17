package main

import (
	"./merkletree"
	"strings"
)

func main() {
	//
	//elems := []byte{0x0, 0x1, 0x4}
	//elemss := []byte{0x9, 0x6, 0x3}
	//el := make([][]byte, 0)
	//el = append(el, elems)
	//el = append(el, elemss)
	//el = append(el, elems)
	//el = append(el, elems)
	//el = append(el, elemss)
	//el = append(el, elemss)

	//el1 := "1"
	//el2 := "2"
	//el3 := "3"
	//el4 := "4"
	//
	//elsdf := make([]string, 0)
	//elsdf = append(elsdf, el1)
	//elsdf = append(elsdf, el2)
	//elsdf = append(elsdf, el3)
	//elsdf = append(elsdf, el4)

	//MerkleTree
	//tree := merkletree.CreateTreeTest(elsdf)
	//tree := merkletree.MerkleTree{}
	//tree.LoadElements( []string{"1", "2", "3", "4"});
	//tree.GetLayers()

	data := []string{"1", "2", "3", "4"}
	tree := merkletree.NewMerkleTree(data)

	//el := initElements()
	//tree, _ := merkletree.CreateTree(el)
	//root := merkletree.GetRoot(tree)
	//println(hex.EncodeToString(root))

	//println(strings.Join([]string{"uuu","1","2"}, ", "))
	//println("^")

	for _, level := range tree.Layers {
		// index is the index where we are
		//println(level)

		println(strings.Join(level, ", "))
		//for _, element := range level {
		//	// index is the index where we are
		//	fmt.Printf("%s ", element)
		//}
	}

}

//func initElements() [][]byte {
//	stringArr := make([]string, 0)
//
//	for o := 0; o < 32; o++ {
//		stringArr = append(stringArr, string(o)+"t")
//	}
//
//	elements := make([][]byte, 0)
//
//	for i := 0; i < len(stringArr); i++ {
//		val, _ := hex.DecodeString(stringArr[i])
//		elements = append(elements, val)
//		println(elements[i])
//	}
//
//	return elements
//}
