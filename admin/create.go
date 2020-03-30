package admin

import (
	"fmt"
	"log"

	"github.com/orsenkucher/nothing/encio"
)

// Такс, сначала пишем просто. Потом выносим композицию
func AdminBot(key encio.EncIO) *Bot {
	fmt.Println("============AdminBot============")
	cfg := cfg(key, "creds/admin.bot.json")
	binder := NewState()
	bot := NewBot(cfg, binder)
	return bot
}

func ClientBot(key encio.EncIO, binder Binder) *Bot {
	fmt.Println("============ClientBot============")
	cfg := cfg(key, "creds/client.bot.json")
	bot := NewBot(cfg, binder)
	return bot
}

func cfg(key encio.EncIO, path string) encio.Config {
	cfg, err := key.GetConfig(path)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}
