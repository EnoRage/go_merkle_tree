package merkletree

import (
	"github.com/ethereum/go-ethereum/crypto/sha3"
	"hash"
	"strconv"
)

const (
	maxLevel = 3
)

//type Tree struct {
//}
//
//type Data interface {
//	HashOf() ([]byte, error)
//}

func initHashFunction() hash.Hash {
	return sha3.NewKeccak256()
}

func CreateTree(elements [][]byte) ([][][]byte, error) {
	h := initHashFunction()
	_, _ = h.Write([]byte(string(' ')))
	empty := h.Sum(nil)

	current := make([][]byte, 0)

	tree := make([][][]byte, 0)
	//for i:=0; i < len(elements); i++ {
	tree = append(tree, elements)
	//}

	for level := 1; level <= maxLevel; level++ {



		for i := 0; i < len(tree[level-1])/2; i++ {
			left := tree[level-1][i*2]
			right := tree[level-1][i*2+1]

			f := initHashFunction()
			bytes := make([]byte, 0)

			//_, _ = f.Write(bytes)

			//fmt.Printf("left1 = %x\n", left)

			bytes = append(left, right...)

			//println(left)
			//fmt.Printf("left2 = %x\n", left)

			if _, err := f.Write(bytes); err != nil {
				return nil, err
			}

			//println(hex.EncodeToString(current[level]))
			//fmt.Printf("%x\n", f.Sum(nil))

			current = append(current, f.Sum(nil))

		}

		//current.length % 2 && level < maxLevel
		if (len(current)%2 > 0) && level < maxLevel {
			current = append(current, empty)
		}

		e := make([]byte, 0)
		e = append(empty, empty...)

		l := initHashFunction()
		_, _ = l.Write(e)
		empty = l.Sum(nil)

		tree = append(tree, current)

		//println(hex.EncodeToString(current[level]))

	}

	return tree, nil

}

func GetRoot(tree [][][]byte) []byte {
	return tree[len(tree)-1][0]
}

func CreateTreeTest(elements []string) ([][]string) {

	empty := string('f')

	current := make([]string, 0)



	tree := make([][]string, 0)

	tree = append(tree, elements)

	for level := 1; level <= maxLevel; level++ {
		println("-----------------------")
		println("Level " + strconv.Itoa(level))

		for i := 0; i < len(tree[level-1])/2; i++ {

			left := tree[level-1][i*2]

			right := tree[level-1][i*2+1]

			println("left " + left)
			println("right " + right)


			current = append(current, left+right)

			print("current: ")
			println(current[level-1])


		}

		//current.length % 2 && level < maxLevel
		if (len(current)%2 > 0) && level < maxLevel {
			current = append(current, empty)
		}

		e := empty + empty

		empty = e

		tree = append(tree, current)
		println("-----------------------")
	}


	//println(tree[0][0])
	//println(tree[0][1])
	//println(tree[1][0])
	//println(tree[1][1])
	//println(tree[2][0])
	//println(tree[2][1])
	//println(tree[3][0])

	//println(tree[3][1])
	//println(tree[4][0])
	//println(tree[4][1])
	//println(tree[5][0])
	//println(tree[5][1])

	return tree

}
