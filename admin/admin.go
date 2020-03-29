package admin

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/orsenkucher/nothing/encio"
)

type Binder interface {
	Bind(tgbotapi.UpdatesChannel)
}

type Bot struct {
	api *tgbotapi.BotAPI
	cfg encio.Config
}

// Такс, сначала пишем просто. Потом выносим композицию
func AdminBot(key encio.EncIO) *Bot {
	fmt.Println("============AdminBot============")
	cfg := cfg(key, "creds/admin.bot.json")
	bot := &Bot{cfg: cfg}
	return bot
}

func cfg(key encio.EncIO, path string) encio.Config {
	cfg, err := key.GetConfig(path)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}
