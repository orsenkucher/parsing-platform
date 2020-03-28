package ppdrop

import "strconv"

type Update interface {
	Update(s *Server)
}

type NewLocation struct {
	Location uint64
	ChatID   int64
}

func (u *NewLocation) Update(s *Server) {
	locname := strconv.FormatUint(u.Location, 10)
	if _, ok := s.Tree.Next[locname]; ok {
		_, ok := s.UsersStates[u.ChatID]
		if !ok {
			s.UsersStates[u.ChatID] = &UsersState{Location: s.Locaitons[u.Location].Name, State: s.Tree.Next[locname], Purchases: []*Purchase{}, Sum: 0, ChatID: u.ChatID}
		} else {
			s.UsersStates[u.ChatID].Location = locname
			s.UsersStates[u.ChatID].State = s.Tree.Next[locname]
		}
		s.ReloadMsg(u.ChatID)
	} else {
		s.ReloadMsg(u.ChatID)
	}
	//s.Bot.SendMessage(u.ChatID, u.Location+"\n"+s.Queries[u.ChatID].ToString())
}

type ChangeState struct {
	Path   string
	ChatID int64
}

func (u *ChangeState) Update(s *Server) {
	query := s.GetQuery(u.ChatID)
	query.State = s.Tree.GetNode(u.Path)
	s.ReloadMsg(u.ChatID)
}

type Add struct {
	Path   string
	ChatID int64
}

func (u *Add) Update(s *Server) {
	query := s.GetQuery(u.ChatID)
	product := s.Tree.GetNode(u.Path)

	for _, purch := range query.Purchases {
		if product == purch.Product {
			purch.Count++
			query.Sum += product.Product.Price
			s.ReloadMsg(u.ChatID)
			return
		}
	}
	query.Purchases = append(query.Purchases, &Purchase{Product: product, Count: 1})
	query.Sum += product.Product.Price
	s.ReloadMsg(u.ChatID)
}

type Sub struct {
	Path   string
	ChatID int64
}

func (u *Sub) Update(s *Server) {
	query := s.GetQuery(u.ChatID)
	product := s.Tree.GetNode(u.Path)

	for _, purch := range query.Purchases {
		if product == purch.Product {
			if purch.Count > 0 {
				purch.Count--
				query.Sum -= product.Product.Price
				s.ReloadMsg(u.ChatID)
			}
			return
		}
	}
}

type BasketReq struct {
	ChatID int64
}

func (u *BasketReq) Update(s *Server) {
	query := s.GetQuery(u.ChatID)
	query.State = nil
	s.ReloadMsg(u.ChatID)
}

type MenuReq struct {
	ChatID int64
}

func (u *MenuReq) Update(s *Server) {
	query := s.GetQuery(u.ChatID)
	query.State = s.Tree.Next[query.Location]
	s.ReloadMsg(u.ChatID)
}

type Reset struct {
	ChatID int64
}

func (u *Reset) Update(s *Server) {
	delete(s.UsersStates, u.ChatID)
	s.ReloadMsg(u.ChatID)
}

func (s *Server) GetQuery(ChatID int64) *UsersState {
	query, ok := s.UsersStates[ChatID]
	if !ok {
		query = &UsersState{Location: "", State: s.Tree, Purchases: []*Purchase{}, Sum: 0, ChatID: ChatID}
		s.UsersStates[ChatID] = query
	}
	return query
}
