package merkletree

import (
	"crypto/sha256"
)

const (
	maxLevel = 31
)

//type Tree struct {
//}
//
//type Data interface {
//	HashOf() ([]byte, error)
//}

func CreateTree(elements [][]byte) ([][][]byte, error) {
	h := sha256.New()
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

			f := sha256.New()
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

		l := sha256.New()
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
