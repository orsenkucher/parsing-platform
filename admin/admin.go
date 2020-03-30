package admin

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/orsenkucher/nothing/encio"
)

type Binder interface {
	Bind(tgbotapi.UpdatesChannel)
}

func NewBot(cfg encio.Config, binder Binder) *Bot {
	bot := &Bot{
		cfg:            cfg,
		usersMsg:       make(map[int64]int),
		messagesMaster: make(chan deferredMessage, 1000),
		messagesBackup: make(chan deferredMessage, 1000),
	}
	bot.init()
	go bot.processMessages()
	go bot.listen(binder)
	return bot
}

func (b *Bot) init() {
	var err error
	b.api, err = tgbotapi.NewBotAPI(b.cfg["token"].(string))
	if err != nil {
		log.Fatalln(err)
	}
	if flag, ok := b.cfg["debug"]; ok {
		b.api.Debug = flag.(bool)
	}
	log.Printf("Authorized on account %s\n", b.api.Self.UserName)
	_, err = b.api.RemoveWebhook()
	if err != nil {
		log.Println("Cant remove webhook")
	}
}

func (b *Bot) listen(binder Binder) {
	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := b.api.GetUpdatesChan(ucfg)
	if err != nil {
		log.Fatalln(err)
	}

	pipe := make(chan tgbotapi.Update)
	binder.Bind(pipe)
	for update := range updates {
		if update.Message != nil {
			if update.Message.Text != "" {
				delcfg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
				if _, err := b.api.DeleteMessage(delcfg); err != nil {
					log.Println(err)
				}
			}
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		}
		pipe <- update
	}
}

type Bot struct {
	api *tgbotapi.BotAPI
	cfg encio.Config

	messagesMaster   chan deferredMessage
	messagesBackup   chan deferredMessage
	lastMessageTimes map[int64]int64

	usersMsg map[int64]int
}

type deferredMessage struct {
	chatID int64
	text   string // Chattable
}

// Такс, сначала пишем просто. Потом выносим композицию
func AdminBot(key encio.EncIO, binder Binder) *Bot {
	fmt.Println("============AdminBot============")
	cfg := cfg(key, "creds/admin.bot.json")
	bot := NewBot(cfg, binder)
	return bot
}

func ClientBot(key encio.EncIO, binder Binder) *Bot {
	fmt.Println("============ClientBot============")
	cfg := cfg(key, "creds/client.bot.json")
	bot := NewBot(cfg, binder)
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
