package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Bot         *Bot
	Admin       *Bot
	UsersStates map[int64]*UsersState
	Updates     chan Update
	Tree        *ProdTree
	Locaitons   map[uint64]*Location
}

func (s *Server) Listen() {
	for {
		upd := <-s.Updates
		upd.Update(s)
	}
}

func StartServer(bot *Bot, admin *Bot) {
	s := Server{Bot: bot, UsersStates: make(map[int64]*UsersState), Locaitons: make(map[uint64]*Location), Updates: make(chan Update), Tree: GenerateTree()}
	s.LoadData()
	for store := range s.Tree.Next {
		fmt.Println(store)
	}
	s.Bot.Updates = s.Updates
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	http.HandleFunc("/", s.GetLocation)
	http.HandleFunc("/locations", s.GiveLocations)
	hsrv := &http.Server{
		Addr:    ":9094",
		Handler: nil, // use default mux
	}
	go s.Listen()
	go func() {
		if err := hsrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Print("Server Started")
	<-done
	log.Print("Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := hsrv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
