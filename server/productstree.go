package server

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type ProdTree struct {
	Product Product
	Next    map[string]*ProdTree
	Prev    *ProdTree
}

func (tree *ProdTree) AddFile(path string) *ProdTree {
	data, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(data), "\n")

	for i := 1; i < len(lines)-1; i++ {
		tree.AddProduct(&lines[i])
	}

	return tree
}

func (tree *ProdTree) AddProduct(line *string) {
	nodes := strings.Split(strings.TrimSpace(*line), ",")
	for j := 0; j < len(nodes)-1; j++ {
		_, ok := tree.Next[nodes[j]]
		if !ok {
			tree.Next[nodes[j]] = &ProdTree{Product: Product{Name: nodes[j], Priority: len(tree.Next) + tree.Product.Priority*100}, Prev: tree, Next: make(map[string]*ProdTree)}
		}
		tree = tree.Next[nodes[j]]
	}
	var err error
	tree.Product.Price, err = strconv.ParseFloat(nodes[len(nodes)-1], 32)
	if err != nil {
		fmt.Println(err)
	}
}

func (tree *ProdTree) Print(s string) {
	fmt.Print(s)
	if tree.Product.Price > 0 {
		fmt.Print(tree.Product.Price, " ")
	}
	fmt.Println(tree.Product.Name)

	for _, node := range tree.Next {
		node.Print(s + "\t")
	}
}

func GenerateTree() *ProdTree {
	tree := &ProdTree{Product: Product{Name: "root"}, Prev: nil, Next: make(map[string]*ProdTree)}
	tree.AddFile("./data/Products.csv")
	return tree
}

func (tree *ProdTree) GetNode(path string) *ProdTree {
	nodes := strings.Split(path, ",")
	for i := 1; i < len(nodes); i++ {
		tree = tree.Next[nodes[i]]
	}
	return tree
}

func (tree *ProdTree) GetPath() string {
	path := tree.Product.Name
	tree = tree.Prev
	for tree != nil {
		path = tree.Product.Name + "," + path
		tree = tree.Prev
	}
	return path
}
