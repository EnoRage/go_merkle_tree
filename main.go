package main

import (
	"./merkletree"
	"encoding/hex"
)

func main() {

	el := initElements()
	tree, _ := merkletree.CreateTree(el)
	//root := merkletree.GetRoot(tree)
	//println(hex.EncodeToString(root))
	println(tree)
}

func initElements() [][]byte {
	stringArr := make([]string, 0)
	stringArr = append(stringArr, "Te")
	stringArr = append(stringArr, "Te")
	stringArr = append(stringArr, "Te")
	stringArr = append(stringArr, "Te")

	elements := make([][]byte, 0)

	for i := 0; i < len(stringArr); i++ {
		val, _ := hex.DecodeString(stringArr[i])
		elements = append(elements, val)
	}

	return elements
}
