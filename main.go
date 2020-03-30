package main

import (
	"flag"
	"fmt"

	"github.com/orsenkucher/nothing/encio"
	"github.com/orsenkucher/parsing-platform/admin"
)

func main() {
	var s = flag.String("s", "", "provide encio password")
	flag.Parse()
	key := encio.NewEncIO(*s)

	// bot := ppdrop.NewBot(key, "creds/client.bot.json")
	// admin := ppdrop.NewBot(key, "creds/admin.bot.json")
	// go bot.Listen()
	// //go admin.Listen()

	a := admin.AdminBot(key)
	fmt.Println(a)
	fmt.Scanln()
}
