package admin

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/orsenkucher/nothing/encio"
)

// type Binder interface {
// 	Bind(tgbotapi.UpdatesChannel)
// }

type Bot struct {
	api              *tgbotapi.BotAPI
	cfg              encio.Config
	messagesMaster   chan deferredMessage
	messagesBackup   chan deferredMessage
	lastMessageTimes map[int64]int64
}

type deferredMessage struct {
	chatID int64
	text   string
}

// Такс, сначала пишем просто. Потом выносим композицию
func AdminBot(key encio.EncIO) *Bot {
	fmt.Println("============AdminBot============")
	cfg := cfg(key, "creds/admin.bot.json")
	bot := &Bot{cfg: cfg}
	go bot.processMessages()
	return bot
}

func cfg(key encio.EncIO, path string) encio.Config {
	cfg, err := key.GetConfig(path)
	if err != nil {
		log.Fatalln(err)
	}
	return cfg
}

func (b *Bot) SendMessage(chatID int64, text string) {
	b.messagesMaster <- deferredMessage{chatID, text}
}

func (b *Bot) sendMessage(msg deferredMessage) (tgbotapi.Message, error) {
	b.lastMessageTimes[msg.chatID] = time.Now().UnixNano()
	return b.api.Send(tgbotapi.NewMessage(msg.chatID, msg.text))
}

// TODO handle errors
func (b *Bot) processMessages() {
	timer := time.NewTicker(time.Second / 30)
	for range timer.C {
		select {
		case msgMaster := <-b.messagesMaster:
			if ok, delta := b.userCanReceiveMessage(msgMaster.chatID); !ok {
				go func() {
					time.Sleep(time.Duration(delta))
					b.messagesBackup <- msgMaster
				}()
			} else {
				b.sendMessage(msgMaster)
			}
		case msgBackup := <-b.messagesBackup:
			b.sendMessage(msgBackup)
		}
	}
}

// TODO do this better
func (b *Bot) userCanReceiveMessage(userId int64) (can bool, delta int64) {
	if t, ok := b.lastMessageTimes[userId]; ok {
		delta = time.Now().UnixNano() - t
		can = delta >= int64(time.Second)
		return
	} else {
		return !can, delta
	}
}
