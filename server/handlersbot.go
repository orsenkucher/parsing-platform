package server

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) NewLocation(chatid int64, loc string) {
	nl := &NewLocation{Location: loc, ChatID: chatid}
	// if b.Updates == nil {
	// 	fmt.Println("asfsdsgs")
	// }
	b.Updates <- nl
}

func (b *Bot) handleCallback(update tgbotapi.Update) {
	data := strings.Split(update.CallbackQuery.Data, "\n")
	ChatID := update.CallbackQuery.Message.Chat.ID
	fmt.Println(data, ChatID)
	if data[0] == "change" {
		b.Updates <- &ChangeState{ChatID: ChatID, Path: data[1]}
	}
	if data[0] == "add" {
		b.Updates <- &Add{ChatID: ChatID, Path: data[1]}
	}
	if data[0] == "sub" {
		b.Updates <- &Sub{ChatID: ChatID, Path: data[1]}
	}
}
