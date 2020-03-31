package admin

import (
	"fmt"
	"log"

	"github.com/orsenkucher/nothing/encio"
)

// Такс, сначала пишем просто. Потом выносим композицию
func AdminBot(key encio.EncIO) *State {
	fmt.Println("============AdminBot============")
	cfg := cfg(key, "creds/admin.bot.json")
	bot := NewBot(cfg)
	state := NewState(bot)
	return state
}

func ClientBot(key encio.EncIO) *State {
	fmt.Println("============ClientBot============")
	cfg := cfg(key, "creds/client.bot.json")
	bot := NewBot(cfg)
	state := NewState(bot)
	return state
}

func cfg(key encio.EncIO, path string) encio.Config {
	cfg, err := key.GetConfig(path)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}
