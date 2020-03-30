package ppdrop

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) HandleMessage(chatid int64, text string) {
	fmt.Print("CheckUser")
	b.Updates <- &CheckUser{ChatID: chatid}
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
	if data[0] == "basket" {
		b.Updates <- &BasketReq{ChatID: ChatID}
	}
	if data[0] == "newbasket" {
		b.Updates <- &NewBasket{ChatID: ChatID}
	}
	if data[0] == "menu" {
		b.Updates <- &MenuReq{ChatID: ChatID}
	}
	if data[0] == "reset" {
		b.Updates <- &Reset{ChatID: ChatID}
	}
	_, err := b.api.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
	if err != nil {
		fmt.Println("Callback: ", err)
	}
}
