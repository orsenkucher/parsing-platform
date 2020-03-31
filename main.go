package main

import (
	"flag"
	"fmt"

	"github.com/orsenkucher/nothing/encio"
	"github.com/orsenkucher/parsing-platform/admin"
	"github.com/orsenkucher/parsing-platform/ppdrop"
)

func main() {
	var s = flag.String("s", "", "provide encio password")
	flag.Parse()
	key := encio.NewEncIO(*s)

	if true {
		bot := ppdrop.NewBot(key, "creds/client.bot.json")
		// adm := ppdrop.NewBot(key, "creds/admin.bot.json")
		go bot.Listen()
		//go admin.Listen()
		adm := admin.AdminBot(key)
		ppdrop.StartServer(bot, adm)
	} else {
		a := admin.AdminBot(key)
		fmt.Println(a)
		fmt.Scanln()
	}
}
