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

type Status int

const (
	New Status = iota
	Sent
	Received
	Processing
	Ready
)

type Basket struct {
	Status    Status
	Location  uint64
	Purchases []*Purchase
	Sum       float64
}

func (state *UsersState) ToString() string {
	fmt.Println(len(state.Baskets), " ", state.Current)
	basket := state.Baskets[state.Current]
	str := ""
	for _, p := range basket.Purchases {
		if p.Count > 0 {
			str += p.ToString()
		}
	}
	str += "Вартість: " + strconv.FormatFloat(basket.Sum, 'f', 2, 64) + "\n"
	return str
}

func (p *Purchase) ToString() string {
	return strconv.Itoa(p.Count) + p.Product.Product.Name + " " + "\n"
}
