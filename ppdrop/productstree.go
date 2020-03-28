package ppdrop

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/orsenkucher/parsing-platform/hsh"
)

type ProdTree struct {
	Product Product
	Next    map[string]*ProdTree
	Prev    *ProdTree
}

func (tree *ProdTree) Print(s string) {
	fmt.Print(s)
	if tree.Product.Price > 0 {
		fmt.Print(strconv.FormatFloat(tree.Product.Price, 'f', 2, 64), " ")
	}
	fmt.Println(tree.Product.Name)

	for _, node := range tree.Next {
		node.Print(s + "\t")
	}
}

func GenerateTree() *ProdTree {
	tree := &ProdTree{Product: Product{Name: "root"}, Prev: nil, Next: make(map[string]*ProdTree)}
	return tree
}

func (tree *ProdTree) GetNode(hash string) *ProdTree {
	nodes := strings.Split(hsh.Hashs[hash], ";")
	for i := 1; i < len(nodes); i++ {
		tree = tree.Next[nodes[i]]
	}
	return tree
}

func (tree *ProdTree) GetHash() string {
	path := tree.Product.Name
	tree = tree.Prev
	for tree != nil {
		path = tree.Product.Name + ";" + path
		tree = tree.Prev
	}
	return hsh.Hash(path)
}
