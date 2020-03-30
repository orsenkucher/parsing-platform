package ppdrop

import (
	"fmt"
	"strconv"
)

type Update interface {
	Update(s *Server)
}

type NewLocation struct {
	Location uint64
	ChatID   int64
}

func (u *NewLocation) Update(s *Server) {
	locstr := strconv.FormatUint(u.Location, 10)
	if _, ok := s.Tree.Next[locstr]; ok {
		state, ok := s.UsersStates[u.ChatID]
		if !ok {
			state = &UsersState{Current: u.Location, State: s.Tree.Next[locstr], Baskets: make(map[uint64]*Basket), ChatID: u.ChatID}
			s.UsersStates[u.ChatID] = state
		}
		_, ok = state.Baskets[u.Location]
		if !ok {
			state.Baskets[u.Location] = &Basket{Location: u.Location, Purchases: []*Purchase{}}
			state.Current = u.Location
		}
		state.State = s.Tree.Next[locstr]
		s.Bot.UpdateMsg(state.GenerateMsg())
	}
	//s.Bot.SendMessage(u.ChatID, u.Location+"\n"+s.Queries[u.ChatID].ToString())
}

type ChangeState struct {
	Path   string
	ChatID int64
}

func (u *ChangeState) Update(s *Server) {
	state := s.GetQuery(u.ChatID)
	state.State = s.Tree.GetNode(u.Path)
	s.Bot.UpdateMsg(state.GenerateMsg())
}

type Add struct {
	Path   string
	ChatID int64
}

func (u *Add) Update(s *Server) {
	state := s.GetQuery(u.ChatID)
	basket := state.Baskets[state.Current]
	product := s.Tree.GetNode(u.Path)

	for _, purch := range basket.Purchases {
		if product == purch.Product {
			purch.Count++
			basket.Sum += product.Product.Price
			s.Bot.UpdateMsg(state.GenerateMsg())
			return
		}
	}
	basket.Purchases = append(basket.Purchases, &Purchase{Product: product, Count: 1})
	basket.Sum += product.Product.Price
	s.Bot.UpdateMsg(state.GenerateMsg())
}

type Sub struct {
	Path   string
	ChatID int64
}

func (u *Sub) Update(s *Server) {
	state := s.GetQuery(u.ChatID)
	basket := state.Baskets[state.Current]
	product := s.Tree.GetNode(u.Path)

	for _, purch := range basket.Purchases {
		if product == purch.Product {
			if purch.Count > 0 {
				purch.Count--
				basket.Sum -= product.Product.Price
				s.Bot.UpdateMsg(state.GenerateMsg())
			}
			return
		}
	}
}

type BasketReq struct {
	ChatID int64
}

func (u *BasketReq) Update(s *Server) {
	state := s.GetQuery(u.ChatID)
	state.State = s.Tree.Next["basket"]
	s.Bot.UpdateMsg(state.GenerateMsg())
}

type MenuReq struct {
	ChatID int64
}

func (u *MenuReq) Update(s *Server) {
	state := s.GetQuery(u.ChatID)
	state.State = s.Tree.Next[strconv.FormatUint(state.Baskets[state.Current].Location, 10)]
	s.Bot.UpdateMsg(state.GenerateMsg())
}

type Reset struct {
	ChatID int64
}

func (u *Reset) Update(s *Server) {
	delete(s.UsersStates, u.ChatID)
	//s.Bot.UpdateMsg(state.GenerateMsg())
}

type CheckUser struct {
	ChatID int64
}

func (u *CheckUser) Update(s *Server) {
	if state, ok := s.UsersStates[u.ChatID]; !ok {
		fmt.Print(state)
		state = &UsersState{State: s.Tree.Next["home"], Baskets: make(map[uint64]*Basket), ChatID: u.ChatID}
		s.UsersStates[u.ChatID] = state
		s.Bot.UpdateMsg(state.GenerateMsg())
	}
}

type NewBasket struct {
	ChatID int64
}

func (u *NewBasket) Update(s *Server) {
	state := s.UsersStates[u.ChatID]
	state.State = s.Tree
	s.Bot.UpdateMsg(state.GenerateMsg())
}

func (s *Server) GetQuery(ChatID int64) *UsersState {
	state, ok := s.UsersStates[ChatID]
	if !ok {
		state = &UsersState{State: s.Tree, Baskets: make(map[uint64]*Basket), ChatID: ChatID}
		s.UsersStates[ChatID] = state
	}
	return state
}
