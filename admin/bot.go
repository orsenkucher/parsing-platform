package admin

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/orsenkucher/nothing/encio"
)

type Writer interface {
	WriteMessages(...tgbotapi.MessageConfig)
}

type Editor interface {
	EditMessages(...tgbotapi.EditMessageTextConfig)
}

type Sender interface {
	Writer
	Editor
}

type Binder interface {
	Bind(tgbotapi.UpdatesChannel, Sender)
	// Bind2(func (tgbotapi.Update))
}

func NewBot(cfg encio.Config, binder Binder) *Bot {
	bot := &Bot{
		cfg:        cfg,
		shipLog:    make(map[int64][]int),
		shipTime:   make(map[int64]int64),
		shipMaster: make(chan defferedShipment, 1000),
		shipBackup: make(chan defferedShipment, 1000),
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
	go binder.Bind(pipe, b)
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

	shipMaster chan defferedShipment // main message channel
	shipBackup chan defferedShipment // reserve channel
	shipTime   map[int64]int64       // last sent time
	shipLog    map[int64][]int       // sent history
}

type defferedShipment struct {
	chatID int64
	cargo  tgbotapi.Chattable
}

var _ Sender = (*Bot)(nil)

func (b *Bot) WriteMessages(mm ...tgbotapi.MessageConfig) {
	for _, m := range mm {
		b.shipMaster <- defferedShipment{chatID: m.ChatID, cargo: m}
	}
}

func (b *Bot) EditMessages(mm ...tgbotapi.EditMessageTextConfig) {
	// for m := range mm {
	// 	b.messagesMaster <- defferedShipment{chatID: m.ChatID, cargo: m}
	// }
	panic("TODO")
}

// TODO handle errors
// This wont work
func (b *Bot) processMessages() {
	timer := time.NewTicker(time.Second / 30)
	for range timer.C {
		select {
		case cargo := <-b.shipMaster:
			if ok, delta := b.userCanReceiveMessage(cargo.chatID); !ok {
				go func() {
					time.Sleep(time.Duration(delta))
					b.shipBackup <- cargo
				}()
			} else {
				b.sendMessage(cargo)
			}
		case defferedCargo := <-b.shipBackup:
			b.sendMessage(defferedCargo)
		}
	}
}

func (b *Bot) sendMessage(msg defferedShipment) (tgbotapi.Message, error) {
	b.shipTime[msg.chatID] = time.Now().UnixNano()
	return b.api.Send(msg.cargo)
}

// TODO do this better
func (b *Bot) userCanReceiveMessage(userId int64) (can bool, delta int64) {
	if t, ok := b.shipTime[userId]; ok {
		delta = time.Now().UnixNano() - t
		can = delta >= int64(time.Second)
		return
	} else {
		return !can, delta
	}
}
