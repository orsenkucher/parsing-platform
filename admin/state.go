package admin

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type StateFn func(tgbotapi.Update) StateFn

type State struct {
	sender  Sender
	workers map[int64]int
}

func NewState(sender Sender) *State {
	s := &State{
		sender:  sender,
		workers: make(map[int64]int),
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
	// txt := fmt.Sprintf("[%v] Жду твой номер, бро🤫", i)

	msg := tgbotapi.NewMessage(chatID(upd), "Подтверждение личности🗄")
	// btn := tgbotapi.NewKeyboardButtonLocation("Локация")
	btn := tgbotapi.NewKeyboardButtonContact("Отправить номер")
	msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(tgbotapi.NewKeyboardButtonRow(btn))
	// s.sender.WriteMessages(msg, msg, msg, msg)
	s.sender.WriteMessages(msg)
	return s.phone
}

var workers = map[string]bool{
	"380962475522": true,
}

func (s *State) phone(upd tgbotapi.Update) StateFn {
	cont := upd.Message.Contact
	if cont == nil {
		return s.start(upd)
	}

	fmt.Println(cont)
	btn := tgbotapi.NewRemoveKeyboard(false)
	if !workers[cont.PhoneNumber] {
		log.Println("Worker not registered")
		msg := tgbotapi.NewMessage(chatID(upd), "Вы тут не работаете🤨.\nНо очень советуем заглянуть в @ppdropbot😉")
		msg.ReplyMarkup = btn
		s.sender.WriteMessages(msg)
		return s.start
	}
	log.Println("Woker connected!")
	s.workers[chatID(upd)] = 1
	msg := tgbotapi.NewMessage(chatID(upd), fmt.Sprintf("%v🤟", cont.FirstName))
	msg.ReplyMarkup = btn
	s.sender.WriteMessages(msg)

	msg = tgbotapi.NewMessage(chatID(upd), fmt.Sprintf("Текущие заказы\n%s", cont.FirstName))
	s.sender.WriteMessages(msg)

	return s.start
}

func chatID(u tgbotapi.Update) int64 {
	if u.Message != nil {
		return u.Message.Chat.ID
	} else {
		return u.CallbackQuery.Message.Chat.ID
	}
}
