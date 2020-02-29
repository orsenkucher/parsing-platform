package ppdrop

import (
	"fmt"
	"log"

	"github.com/orsenkucher/nothing/encio"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	API   *tgbotapi.BotAPI
	cfg   encio.Config
	Users []int64
}

func NewBot(key encio.EncIO) *Bot {
	cfg, err := key.GetConfig("ppdrop/bot.json")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(cfg)
	b := Bot{cfg: cfg}
	b.initAPI()
	return &b
}
func (b *Bot) initAPI() {
	var err error
	b.API, err = tgbotapi.NewBotAPI(b.cfg["token"].(string))
	if err != nil {
		log.Fatalln(err)
	}

	b.API.Debug = false
	log.Printf("Authorized on account %s\n", b.API.Self.UserName)

	_, err = b.API.RemoveWebhook()
	if err != nil {
		log.Println("Cant remove webhook")
	}
}

func (b *Bot) Listen() {

}
