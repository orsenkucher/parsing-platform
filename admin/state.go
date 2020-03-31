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
	basket  string
}

func NewState(sender Sender) *State {
	s := &State{
		sender:  sender,
		workers: make(map[int64]int),
	}
	go sender.Bind(s.bind)
	return s
}

func (s *State) Basket(basket string) {
	s.basket = basket
	fmt.Println("Basket set")
	fmt.Println(basket)
	for cid := range s.workers {
		s.showBasketToUser(cid)
	}
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
	return s.showBasket(upd)
}

func (s *State) showBasket(upd tgbotapi.Update) StateFn {
	chatID := chatID(upd)
	return s.showBasketToUser(chatID)
}

func (s *State) showBasketToUser(chatID int64) StateFn {
	if s.basket == "" {
		s.sender.WriteMessages(tgbotapi.NewMessage(chatID, "Новых заказов нет"))
		return s.showBasket
	}
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Текущие заказы\n%s", s.basket))
	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Выполнить", "confirm"),
			// tgbotapi.NewInlineKeyboardButtonData("Отменить","reject"),
		),
	)
	s.sender.WriteMessages(msg)
	return s.confirm
}

func (s *State) confirm(upd tgbotapi.Update) StateFn {
	if upd.CallbackQuery == nil {
		return s.confirm
	}
	if upd.CallbackQuery.Data != "confirm" {
		return s.confirm
	}
	cid := chatID(upd)
	fmt.Println(cid)
	return s.showBasketToUser(cid)
}

func (s *State) loop(upd tgbotapi.Update) StateFn {
	return s.loop
}

func chatID(u tgbotapi.Update) int64 {
	if u.Message != nil {
		return u.Message.Chat.ID
	} else {
		return u.CallbackQuery.Message.Chat.ID
	}
}
