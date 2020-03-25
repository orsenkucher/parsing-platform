package main

import (
	"flag"

	"github.com/orsenkucher/nothing/encio"
	"github.com/orsenkucher/parsing-platform/bot"
)

func main() {
	var s = flag.String("s", "", "provide encio password")
	flag.Parse()
	key := encio.NewEncIO(*s)
	bot := bot.NewBot(key)
	bot.Listen()
}
