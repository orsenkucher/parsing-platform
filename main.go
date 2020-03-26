package main

import (
	"flag"

	"github.com/orsenkucher/nothing/encio"
	"github.com/orsenkucher/parsing-platform/bot"
	"github.com/orsenkucher/parsing-platform/data"
	"github.com/orsenkucher/parsing-platform/server"
)

func main() {
	tree := data.ProdTree{Product: server.Product{Name: "root"}, Prev: nil, Next: make(map[string]*data.ProdTree)}
	tree.AddFile("./data/Products.csv")
	tree.Print("")
	var s = flag.String("s", "", "provide encio password")
	flag.Parse()
	key := encio.NewEncIO(*s)
	bot := bot.NewBot(key)
	bot.Listen()
	server.StartServer(bot)
}
