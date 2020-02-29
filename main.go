package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/orsenkucher/nothing/encio"
)

func main() {
	var s = flag.String("s", "", "provide encio password")
	flag.Parse()
	key := encio.NewEncIO(*s)
	cfg, err := key.GetConfig("bot.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg)
}
