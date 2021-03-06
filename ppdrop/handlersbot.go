package ppdrop

import (
	"fmt"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) HandleMessage(chatid int64, text string) {
	if text != "/start" {
		b.Updates <- &CheckUser{ChatID: chatid}
	} else {
		tgmsg := tgbotapi.NewMessage(chatid, `Здравствуйте, я бот "Забирай сам"
		Здесь можно заказать товар в ближайшем к вам магазине.
		Выбирайте нужный магазин, заказывайте товар.  Сотрудник магазина получит и соберет ваш заказ. После этого вы сможете забрать пакет с товаром без очереди или у окна выдачи магазина по номеру заказа.`)
		button := tgbotapi.NewInlineKeyboardButtonData("Заберу сам", "newbasket\n")
		rows := [][]tgbotapi.InlineKeyboardButton{}
		rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
		buttons := tgbotapi.NewInlineKeyboardMarkup(rows...)
		tgmsg.ReplyMarkup = &buttons
		b.ResendMsg(tgmsg)
	}
}

func (b *Bot) handleCallback(update tgbotapi.Update) {
	if id, ok := b.UsersMsg[update.CallbackQuery.Message.Chat.ID]; !ok || id != update.CallbackQuery.Message.MessageID {
		b.Updates <- &CheckUser{ChatID: update.CallbackQuery.Message.Chat.ID}
		return
	}
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
	if data[0] == "home" {
		b.Updates <- &HomeReq{ChatID: ChatID}
	}
	if data[0] == "agree" {
		b.Updates <- &AgreeHome{ChatID: ChatID}
	}
	if data[0] == "newbasket" {
		b.Updates <- &NewBasket{ChatID: ChatID}
	}
	if data[0] == "sendbasket" {
		b.Updates <- &SendBasket{ChatID: ChatID}
	}
	if data[0] == "catalog" {
		b.Updates <- &CatalogReq{ChatID: ChatID}
	}
	if data[0] == "reset" {
		loc, _ := strconv.ParseUint(data[1], 10, 64)
		b.Updates <- &Reset{ChatID: ChatID, BasketID: loc}
	}
	if data[0] == "location" {
		loc, _ := strconv.ParseUint(data[1], 10, 64)
		fmt.Println("change location to ", loc)
		b.Updates <- &NewLocation{Location: loc, ChatID: ChatID}
	}
	_, err := b.api.AnswerCallbackQuery(tgbotapi.NewCallback(update.CallbackQuery.ID, ""))
	if err != nil {
		fmt.Println("Callback: ", err)
	}
}
