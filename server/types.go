package server

import (
	"fmt"
	"strconv"
)

type Product struct {
	Name  string
	Price float64
}

type Purchase struct {
	Product *Product
	Count   int
}

type Query struct {
	State     int
	Location  string
	Purchases []Purchase
	Sum       float64
	ChatID    int64
}

func (q *Query) ToString() string {
	fmt.Println()
	str := ""
	for _, p := range q.Purchases {
		str += p.ToString()
	}
	str += "sum: " + strconv.FormatFloat(q.Sum, 'f', 2, 64) + "\n"
	return str
}

func (p *Purchase) ToString() string {
	return p.Product.Name + " " + string(p.Count) + "\n"
}
