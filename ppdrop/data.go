package ppdrop

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func (s *Server) LoadData() {
	path := "./data/"

	files, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}
	for i, dir := range files {
		if dir.IsDir() {
			s.AddLocationsFromFile(path+dir.Name()+"/Locations.csv", uint64(i+1))
			s.Tree.AddProductFromFile(path+dir.Name()+"/Products.csv", uint64(i+1))
		}
	}
	for id, loc := range s.Locaitons {
		fmt.Println(id, " ", loc.Adress, " ", loc.Lat, " ", loc.Lng)
	}
}

func (s *Server) AddLocationsFromFile(path string, netid uint64) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	lines := strings.Split(string(data), "\n")
	for i := 1; i < len(lines)-1; i++ {
		s.AddLocation(&lines[i], netid)
	}
}

func (s *Server) AddLocation(line *string, netid uint64) {
	parts := strings.Split(strings.TrimSpace(*line), ";")
	storeid, err := strconv.Atoi(parts[0])
	id := netid<<32 + uint64(storeid)
	lat := parseFloat(parts[3])
	long := parseFloat(parts[4])
	if err != nil {
		fmt.Println("location from file err:", err)
	}
	location := Location{ID: id, Lat: lat, Lng: long, Name: parts[1], Adress: parts[2]}
	s.Locaitons[location.ID] = &location
}

func (tree *ProdTree) AddProductFromFile(path string, netid uint64) *ProdTree {
	data, _ := ioutil.ReadFile(path)
	lines := strings.Split(string(data), "\n")

	for i := 1; i < len(lines)-1; i++ {
		tree.AddProduct(&lines[i], netid)
	}

	return tree
}

func (tree *ProdTree) AddProduct(line *string, netid uint64) {
	nodes := strings.Split(strings.TrimSpace(*line), ";")
	id, err := strconv.Atoi(nodes[0])
	if err != nil {
		fmt.Print(err)
	}
	locid := strconv.FormatUint(netid<<32+uint64(id), 10)

	_, ok := tree.Next[locid]
	if !ok {
		tree.Next[locid] = &ProdTree{Product: Product{Name: locid}, Prev: tree, Next: make(map[string]*ProdTree)}
	}
	tree = tree.Next[locid]
	for j := 6; j < len(nodes); j++ {
		_, ok := tree.Next[nodes[j]]
		if !ok {
			tree.Next[nodes[j]] = &ProdTree{Product: Product{Name: nodes[j], Priority: len(tree.Next) + tree.Product.Priority*100}, Prev: tree, Next: make(map[string]*ProdTree)}
		}
		tree = tree.Next[nodes[j]]
	}

	prodname := nodes[2]
	prodid, _ := strconv.Atoi(nodes[1])
	price := parseFloat(nodes[5])
	_, ok = tree.Next[prodname]
	if !ok {
		tree.Next[prodname] = &ProdTree{Product: Product{Name: prodname, Priority: len(tree.Next) + tree.Product.Priority*100, ID: prodid, Price: price}, Prev: tree, Next: make(map[string]*ProdTree)}
	}

	if err != nil {
		fmt.Println(err)
	}
}

//func LoadProducts

func parseFloat(str string) float64 {
	res, err := strconv.ParseFloat(strings.ReplaceAll(str, ",", "."), 64)
	if err != nil {
		fmt.Println(err)
	}
	return res
}
