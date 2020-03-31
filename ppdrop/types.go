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

func (s *Status) ToString() string {
	if *s == New {
		return "Новый"
	} else if *s == Sent {
		return "Отправлен"
	} else if *s == Received {
		return "В очереди"
	} else if *s == Processing {
		return "Собирается"
	}
	return "Готово!"
}

type Basket struct {
	Status    Status
	Location  uint64
	Purchases []*Purchase
	Sum       float64
}

func (basket *Basket) ToString() string {
	str := Locations[basket.Location].Name + "\n\n"
	for _, p := range basket.Purchases {
		if p.Count > 0 {
			str += p.ToString()
		}
	}
	str += "\nВартість: " + strconv.FormatFloat(basket.Sum, 'f', 2, 64) + "\n\n"
	if basket.Status != New {
		str += basket.Status.ToString() + "\n"
	}
	return str
}

func (state *UsersState) ToString() string {
	fmt.Println(len(state.Baskets), " ", state.Current)
	basket := state.Baskets[state.Current]
	return basket.ToString()
}

func (p *Purchase) ToString() string {
	return strconv.Itoa(p.Count) + p.Product.Product.Name + " " + "\n"
}
