package merkletree

type MerkleTree struct {
	Layers [][]string
}

func NewMerkleTree(data []string) *MerkleTree {
	tree := new(MerkleTree)

	lambda := func(x string) string { return hashFunction(x) }
	tree.Layers = getLayers(Map(lambda, data))

	return tree
}

func hashFunction(data string) string {
	return "hash(" + data + ")"
}

type Tree [][]string

func Map(f func(string) string, xs []string) []string {
	ys := make([]string, len(xs))
	for i, x := range xs {
		ys[i] = f(x)
	}
	return ys
}

func concat(a string, b string) string {
	return "concat(" + a + "," + b + ")"
}

func getLayers(elements []string) Tree {
	emptyLeveled := hashFunction("0")
	if (len(elements) % 2) > 0 {
		elements = append(elements, emptyLeveled)
	}

	myTree := Tree{elements}
	const maxLevel = 31

	for level := 1; level < 10; level++ {
		var current []string
		for i := 0; i < len(myTree[level-1])/2; i++ {
			a := myTree[level-1][i*2]
			b := myTree[level-1][i*2+1]
			hash := hashFunction(concat(a, b))

			current = append(current, hash)
			// current.push(hash)
		}

		if (len(current)%2 > 0) && level < maxLevel {
			//current.push(emptyLeveled);
			current = append(current, emptyLeveled)
		}

		emptyLeveled = hashFunction(concat(emptyLeveled, emptyLeveled))

		// tree.push(current);
		myTree = append(myTree, current)
	}

	//for (let level = 1; level <= maxLevel; level++) {
	//}

	return myTree
}

func (tree *MerkleTree) getRoot(elements []string) string {
	return tree.Layers[len(tree.Layers)-1][0]
}

//func getHexRoot () {
//	return bufferToHex(this.getRoot());
//}
//
//getProof (idx) {
//	if (idx == = -1) {
//		throw new Error('Element does not exist in Merkle tree');
//	}
//
//	return this.layers.reduce((proof, layer) = > {
//	const pairElement = this.getPairElement(idx, layer);
//	if (pairElement) {
//		proof.push(pairElement);
//	}
//
//	idx = Math.floor(idx / 2);
//	return proof;
//	}, []);
//}
//
//getHexProof (idx) {
//	const result = this.getProof(idx);
//	return this.bufArrToHexArr(result);
//}
//
//getPairElement (idx, layer) {
//	const pairIdx = idx % 2 == = 0 ? idx + 1: idx - 1;
//	if (pairIdx < layer.length) {
//		return layer[pairIdx];
//	} else {
//		return null;
//	}
//}
//
//bufIndexOf (el, arr) {
//let hash;
//
//// Convert element to 32 byte hash if it is not one already
//if (el.length != = 32 || !Buffer.isBuffer(el)) {
//hash = this.hashFunction(el);
//} else {
//hash = el;
//}
//
//for (let i = 0; i < arr.length; i++) {
//if (hash.equals(arr[i])) {
//return i;
//}
//}
//
//return -1;
//}
//
//bufArrToHexArr (arr) {
//if (arr.some(el = > !Buffer.isBuffer(el))) {
//throw new Error('Array is not an array of buffers');
//}
//
//return arr.map(el = > '0x' + el.toString('hex'));
//}
//}
//
//function keccak160 (element) {
//return keccak256(element).slice(12,32);
//}
//
//module.exports = {
//MerkleTree,
//keccak160
//};
