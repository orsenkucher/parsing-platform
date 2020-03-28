package ppdrop

import (
	"fmt"
	"strconv"
)

type Product struct {
	Name     string
	Priority int
	ID       int
	Price    float64
}

type Purchase struct {
	Product *ProdTree
	Count   int
}

type Location struct {
	ID     uint64  `json:"id"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
	Name   string  `json:"name"`
	Adress string  `json:"adress"`
}

type UsersState struct {
	State     *ProdTree
	Location  string
	Purchases []*Purchase
	Sum       float64
	ChatID    int64
}

func (q *UsersState) ToString() string {
	fmt.Println()
	str := ""
	for _, p := range q.Purchases {
		if p.Count > 0 {
			str += p.ToString()
		}
	}
	str += "sum: " + strconv.FormatFloat(q.Sum, 'f', 2, 64) + "\n"
	return str
}

//func (q *Query)

func (p *Purchase) ToString() string {
	return p.Product.Product.Name + " " + strconv.Itoa(p.Count) + "\n"
}
