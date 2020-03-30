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

	bot := ppdrop.NewBot(key, "creds/client.bot.json")
	admin := ppdrop.NewBot(key, "creds/admin.bot.json")
	go bot.Listen()
	//go admin.Listen()
	ppdrop.StartServer(bot, admin)
	// a := admin.AdminBot(key)
	// fmt.Println(a)
	// fmt.Scanln()
}
