package data

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/orsenkucher/parsing-platform/server"
)

type ProdTree struct {
	Product server.Product
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
			tree.Next[nodes[j]] = &ProdTree{Product: server.Product{Name: nodes[j]}, Prev: tree, Next: make(map[string]*ProdTree)}
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
