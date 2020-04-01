package admin

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type StateFn func(tgbotapi.Update) StateFn

type State struct {
	sender  Sender
	workers map[int64]string
	bask    basket
	sfn     map[int64]StateFn
}

type basket struct {
	status   int // 2-rec, 3-proc, 4-rdy
	content  string
	callback func(string)
}

func NewState(sender Sender) *State {
	s := &State{
		sender:  sender,
		workers: make(map[int64]string),
		sfn:     make(map[int64]StateFn),
	}
	go sender.Bind(s.bind)
	return s
}

func (s *State) Basket(status int, content string, callback func(confimedBy string)) {
	s.bask = basket{
		status:   status,
		content:  content,
		callback: callback,
	}
	fmt.Println("Basket set")
	fmt.Println(content)
	for cid := range s.workers {
		s.sfn[cid] = s.showBasketToUser(cid)
	}
}

func (s *State) bind(upds tgbotapi.UpdatesChannel) {
	for upd := range upds {
		cid := chatID(upd)
		state, ok := s.sfn[cid]
		if !ok {
			state = s.start
		}
		s.sfn[cid] = state(upd)
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
	"380962475522":  true,
	"+380962475522": true,
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
		s.sender.EditMessages(msg)
		return s.start
	}
	log.Println("Woker connected!")
	s.workers[chatID(upd)] = cont.FirstName
	msg := tgbotapi.NewMessage(chatID(upd), fmt.Sprintf("%v🤟", cont.FirstName))
	msg.ReplyMarkup = btn
	s.sender.EditMessages(msg)
	return s.showBasket(upd)
}

func (s *State) showBasket(upd tgbotapi.Update) StateFn {
	chatID := chatID(upd)
	return s.showBasketToUser(chatID)
}

func (s *State) showBasketToUser(chatID int64) StateFn {
	if s.bask.content == "" {
		s.sender.WriteMessages(tgbotapi.NewMessage(chatID, "Новых заказов нет"))
		return s.showBasket
	}
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Текущие заказы\n%s", s.bask.content))

	txt := "Выполнить"
	if s.bask.status == 3 {
		txt = "Готово"
	}
	btn := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(txt, "confirm"),
			// tgbotapi.NewInlineKeyboardButtonData("Отменить","reject"),
		),
	)
	if s.bask.status != 4 {
		msg.ReplyMarkup = btn
	}
	s.sender.EditMessages(msg)
	fmt.Println("waiting confirm")
	return s.confirm
}

func (s *State) confirm(upd tgbotapi.Update) StateFn {
	fmt.Println("In confirm()")
	if upd.CallbackQuery == nil {
		return s.confirm
	}
	fmt.Println(upd.CallbackQuery.Data)
	if upd.CallbackQuery.Data != "confirm" {
		return s.confirm
	}
	fmt.Println("\tCONFIRMED")
	log.Println("\tCONFIRMED")
	cid := chatID(upd)
	fmt.Println(cid)
	if s.bask.callback != nil {
		s.bask.callback(s.workers[cid])
	}
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
