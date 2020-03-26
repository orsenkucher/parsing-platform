package server

type Update interface {
	Update(s *Server)
}

type NewLocation struct {
	Location string
	ChatID   int64
}

func (u NewLocation) Update(s *Server) {
	_, ok := s.Queries[u.ChatID]
	if !ok {
		s.Queries[u.ChatID] = &Query{Location: u.Location, State: 0, Purchases: []Purchase{}, Sum: 0, ChatID: u.ChatID}
	} else {
		s.Queries[u.ChatID].Location = u.Location
	}
	s.Bot.SendMessage(u.ChatID, u.Location+"\n"+s.Queries[u.ChatID].ToString())
}
