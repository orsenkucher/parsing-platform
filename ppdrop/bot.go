package ppdrop

import (
	"fmt"
	"log"

	"github.com/orsenkucher/nothing/encio"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type Bot struct {
	api      *tgbotapi.BotAPI
	cfg      encio.Config
	UsersMsg map[int64]int
	Updates  chan Update
}

func NewBot(key encio.EncIO, cfgPath string) *Bot {
	fmt.Println("============CFG============")
	cfg := botCfg(key, cfgPath)
	fmt.Println(cfg)
	bot := Bot{cfg: cfg}
	bot.initAPI()
	return &bot
}

func botCfg(key encio.EncIO, path string) encio.Config {
	cfg, err := key.GetConfig(path)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
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
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	b.UsersMsg = make(map[int64]int)

	updates, err := b.api.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.CallbackQuery != nil {
			b.handleCallback(update)
			continue
		}

		if update.Message != nil {
			if update.Message.Text != "" {
				b.HandleMessage(update.Message.Chat.ID, update.Message.Text)
				delcfg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
				if _, err := b.api.DeleteMessage(delcfg); err != nil {
					log.Println(err)
				}
			}

			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			continue
		}
	}
}

func (b *Bot) UpdateMsg(msg tgbotapi.MessageConfig) {
	prevID, ok := b.UsersMsg[msg.ChatID]
	if ok {
		updatemsg := tgbotapi.NewEditMessageText(msg.ChatID, prevID, msg.Text)
		updatemsg.ReplyMarkup = msg.ReplyMarkup.(*tgbotapi.InlineKeyboardMarkup)
		_, err := b.api.Send(updatemsg)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		msgtg, err := b.api.Send(msg)
		b.UsersMsg[msg.ChatID] = msgtg.MessageID
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (b *Bot) ResendMsg(msg tgbotapi.MessageConfig) {
	prevID, ok := b.UsersMsg[msg.ChatID]
	if ok {
		delcfg := tgbotapi.NewDeleteMessage(msg.ChatID, prevID)
		_, err := b.api.DeleteMessage(delcfg)
		if err != nil {
			fmt.Println(err)
		}
	}
	msgtg, err := b.api.Send(msg)
	b.UsersMsg[msg.ChatID] = msgtg.MessageID
	if err != nil {
		fmt.Println(err)
	}
}
