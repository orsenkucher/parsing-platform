package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type LocationReq struct {
	UID int64   `json:"uid"`
	X   float64 `json:"x"`
	Y   float64 `json:"y"`
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
	fmt.Println()
	fmt.Println(req.X)
	fmt.Println(req.Y)
	fmt.Println("uid: ", req.UID)
}
