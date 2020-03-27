package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type LocationReq struct {
	ChatID   string `json:"chatid"`
	Location uint64 `json:"location"`
}

func (s *Server) GetLocation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received location")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	var req LocationReq
	err = json.Unmarshal(body, &req)
	if err != nil {
		log.Println(err)
	}
	chatid, _ := strconv.ParseInt(req.ChatID, 10, 64)
	s.Updates <- &NewLocation{Location: req.Location, ChatID: chatid}
	fmt.Println()
	fmt.Println(req.Location)
	fmt.Println("uid: ", req.ChatID)
}

func (s *Server) GiveLocations(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Give locations")
	locs := []*Location{}
	for _, loc := range s.Locaitons {
		locs = append(locs, loc)
	}
	data, err := json.Marshal(locs)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
	fmt.Fprint(w, string(data))
}
