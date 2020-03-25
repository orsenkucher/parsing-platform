package server

type Location struct {
	X int
	Y int
}

type Product struct {
	Price int
}

type Update interface {
	Update(s *Server)
}

type Query struct {
	State    int
	Location Location
	Products []Product
	Sum      int
}
