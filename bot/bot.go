package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/orsenkucher/nothing/encio"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api *tgbotapi.BotAPI
	cfg encio.Config
}

func NewBot(key encio.EncIO) *Bot {
	cfg, err := key.GetConfig("bot/bot.json")
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
	b.api, err = tgbotapi.NewBotAPI(b.cfg["token"].(string))
	if err != nil {
		log.Fatalln(err)
	}

	b.api.Debug = false
	log.Printf("Authorized on account %s\n", b.api.Self.UserName)

	_, err = b.api.RemoveWebhook()
	if err != nil {
		log.Println("Cant remove webhook")
	}
}

func (b *Bot) Listen() {

}

func (b *Bot) SpreadMessage(users []int64, msg string) {
	log.Printf("Sending message to %v users\n", len(users))
	for _, u := range users {
		time.Sleep(100 * time.Millisecond)

		log.Printf("Deleting previous msg for %v\n", u)

		log.Printf("Sending to %v\n", u)
		tgmsg := tgbotapi.NewMessage(u, msg)
		_, err := b.api.Send(tgmsg)
		if err != nil {
			log.Println(err)
		}
	}
}
