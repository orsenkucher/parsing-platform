package server

type Update interface {
	Update(s *Server)
}

type NewLocation struct {
	Location string
	ChatID   int64
}

func (u *NewLocation) Update(s *Server) {
	if _, ok := s.Tree.Next[u.Location]; ok {
		_, ok := s.Queries[u.ChatID]
		if !ok {
			s.Queries[u.ChatID] = &Query{Location: u.Location, State: s.Tree.Next[u.Location], Purchases: []*Purchase{}, Sum: 0, ChatID: u.ChatID}
		} else {
			s.Queries[u.ChatID].Location = u.Location
			s.Queries[u.ChatID].State = s.Tree.Next[u.Location]
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
	delete(s.Queries, u.ChatID)
	s.ReloadMsg(u.ChatID)
}

func (s *Server) GetQuery(ChatID int64) *Query {
	query, ok := s.Queries[ChatID]
	if !ok {
		query = &Query{Location: "", State: s.Tree, Purchases: []*Purchase{}, Sum: 0, ChatID: ChatID}
		s.Queries[ChatID] = query
	}
	return query
}
