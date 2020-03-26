package main

import (
	"flag"

	"github.com/orsenkucher/nothing/encio"
	"github.com/orsenkucher/parsing-platform/server"
)

func main() {
	var s = flag.String("s", "", "provide encio password")
	flag.Parse()
	key := encio.NewEncIO(*s)
	bot := server.NewBot(key)
	go bot.Listen()
	server.StartServer(bot)
}
