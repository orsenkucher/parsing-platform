package main

import (
	"flag"

	"github.com/orsenkucher/nothing/encio"
	"github.com/orsenkucher/parsing-platform/ppdrop"
)

func main() {
	var s = flag.String("s", "", "provide encio password")
	flag.Parse()
	key := encio.NewEncIO(*s)
	bot := ppdrop.NewBot(key)
	bot.Listen()
}
