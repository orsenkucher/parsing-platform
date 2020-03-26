package bot

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *Bot) UpdateMsg(msg tgbotapi.MessageConfig) {
	prevID, ok := b.UsersMsg[msg.ChatID]
	if ok {
		updatemsg := tgbotapi.NewEditMessageText(msg.ChatID, prevID, msg.Text)
		_, err := b.api.Send(updatemsg)
		if err != nil {
			fmt.Print(err)
		}
	} else {
		msgtg, err := b.api.Send(msg)
		b.UsersMsg[msg.ChatID] = msgtg.MessageID
		if err != nil {
			fmt.Print(err)
		}
	}
}
