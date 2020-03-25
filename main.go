package main

import (
	"flag"
	"fmt"

	"github.com/orsenkucher/nothing/encio"
	"github.com/orsenkucher/parsing-platform/bot"
	"github.com/orsenkucher/parsing-platform/server"
)

func main() {
	var s = flag.String("s", "", "provide encio password")
	flag.Parse()
	key := encio.NewEncIO(*s)
	bot := bot.NewBot(key)
	bot.Listen()
	serv := server.Server{Bot: bot}
	fmt.Print(serv)
}
