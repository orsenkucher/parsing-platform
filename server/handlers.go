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
	Location string `json:"location"`
}

func (s *Server) GetLocation(w http.ResponseWriter, r *http.Request) {
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
