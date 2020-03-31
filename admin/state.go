package admin

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type StateFn func(tgbotapi.Update) StateFn

type State struct {
	sender Sender
	users  map[int64]int
}

func NewState(sender Sender) *State {
	s := &State{
		sender: sender,
		users:  make(map[int64]int),
	}
	sender.Bind(s.bind)
	return s
}

func (s *State) bind(upds tgbotapi.UpdatesChannel) {
	state := s.start
	for upd := range upds {
		state = state(upd)
	}
}

func (s *State) start(upd tgbotapi.Update) StateFn {
	msg := tgbotapi.NewMessage(chatID(upd), "Жду твой номер, бро🤫")
	// btn := tgbotapi.NewKeyboardButtonLocation("Send to bot")
	btn := tgbotapi.NewKeyboardButtonContact("Отправить номер")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(btn))
	s.sender.WriteMessages(msg)
	return s.phone
}

func (s *State) phone(upd tgbotapi.Update) StateFn {
	// msg := tgbotapi.NewMessage(chatID(upd), "Great")
	// edt := tgbotapi.NewEditMessageText()
	// s.sender.EditMessages()
	cont := upd.Message.Contact
	fmt.Println(cont)
	msg := tgbotapi.NewMessage(chatID(upd), "Отлично")
	btn := tgbotapi.NewRemoveKeyboard(false)
	msg.ReplyMarkup = btn
	s.sender.WriteMessages(msg, tgbotapi.NewMessage(chatID(upd), fmt.Sprint(cont)))
	return s.start
}

func chatID(u tgbotapi.Update) int64 {
	if u.Message != nil {
		return u.Message.Chat.ID
	} else {
		return u.CallbackQuery.Message.Chat.ID
	}
}
