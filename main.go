package main

import (
	"fmt"
	"log"

	"github.com/orsenkucher/nothing/encio"
)

func main() {
	fmt.Println("go works")
	key := encio.NewEncIO("password")
	cfg, err := key.GetConfig("bot.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg)
}
