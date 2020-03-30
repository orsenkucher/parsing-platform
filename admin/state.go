package admin

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

type StateFn func(tgbotapi.Update) StateFn

type State struct {
	sender Sender
	users  map[int64]int
}

func NewState() *State {
	return &State{users: make(map[int64]int)}
}

var _ Binder = (*State)(nil)

func (s *State) Bind(upds tgbotapi.UpdatesChannel, sender Sender) {
	s.sender = sender
	state := s.start
	for upd := range upds {
		state = state(upd)
	}
}

func (s *State) start(upd tgbotapi.Update) StateFn {
	s.sender.WriteMessages(tgbotapi.NewMessage(chatID(upd), "Send location"))
	return s.location
}

func (s *State) location(upd tgbotapi.Update) StateFn {
	// msg := tgbotapi.NewMessage(chatID(upd), "Great")
	// edt := tgbotapi.NewEditMessageText()
	// s.sender.EditMessages()
	s.sender.WriteMessages(tgbotapi.NewMessage(chatID(upd), "Great"))
	return s.start
}

func chatID(u tgbotapi.Update) int64 {
	if u.Message != nil {
		return u.Message.Chat.ID
	} else {
		return u.CallbackQuery.Message.Chat.ID
	}
}
