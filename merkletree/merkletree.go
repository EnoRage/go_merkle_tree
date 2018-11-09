package merkletree

import (
	"crypto/sha256"
	"encoding/hex"
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
	tree = append(tree, elements)

	for level := 1; level <= maxLevel; level++ {

		for i := 0; i < len(tree[level - 1]) / 2; i++ {

			left := tree[level - 1][i * 2]
			right := tree[level - 1][i * 2 + 1]

			h = sha256.New()

			if _, err := h.Write(append(left, right...)); err != nil {
				return nil, err
			}

			current = append(current, h.Sum(nil))
		}

		if len(current) % 2 == 0 && level < maxLevel {
			current = append(current, empty)
		}

		h = sha256.New()
		_, _ = h.Write(append(empty, empty...))
		empty = h.Sum(nil)

		tree = append(tree, current)
		print("Current level ")
		print(level)
		println(" ")
		println(hex.EncodeToString(current[level]))

	}

	return tree, nil
}

func GetRoot (tree [][][]byte) []byte {
	return tree[len(tree) - 1][0]
}
