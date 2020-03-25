package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/orsenkucher/parsing-platform/bot"
)

type Server struct {
	Bot     *bot.Bot
	Queries map[int64]Query
	Updates chan Update
}

func (s *Server) Listen() {
	for {
		upd := <-s.Updates
		upd.Update(s)
	}
}

func StartServer(bot *bot.Bot) {
	s := Server{Bot: bot, Queries: make(map[int64]Query), Updates: make(chan Update)}
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	http.HandleFunc("/", s.GetLocation)
	hsrv := &http.Server{
		Addr:    ":9094",
		Handler: nil, // use default mux
	}
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
