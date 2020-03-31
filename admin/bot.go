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
	EditMessages(...tgbotapi.MessageConfig)
}

type Binder interface {
	Bind(func(tgbotapi.UpdatesChannel))
}

// ChatManager
type Sender interface {
	Writer
	Editor
	Binder
}

type Bot struct {
	api         *tgbotapi.BotAPI
	cfg         encio.Config
	updc        chan tgbotapi.Update            // TODO(multiple bindings)
	shipGrid    map[int64]chan defferedShipment // 1 drip/sec on highway per user
	shipHighway chan defferedShipment           // 30 shipments/sec
	shipTime    map[int64]int64                 // last sent time
	shipLog     map[int64][]int                 // sent history
}

type defferedShipment struct {
	chatID int64
	cargo  tgbotapi.Chattable
}

func NewBot(cfg encio.Config) *Bot {
	bot := &Bot{
		cfg:         cfg,
		updc:        make(chan tgbotapi.Update),
		shipLog:     make(map[int64][]int),
		shipTime:    make(map[int64]int64),
		shipGrid:    make(map[int64]chan defferedShipment),
		shipHighway: make(chan defferedShipment, 1000),
	}
	bot.init()
	go bot.processMessages()
	go bot.listen()
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

func (b *Bot) listen() {
	ucfg := tgbotapi.NewUpdate(0)
	ucfg.Timeout = 60
	updates, err := b.api.GetUpdatesChan(ucfg)
	if err != nil {
		log.Fatalln(err)
	}
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
		b.updc <- update
	}
}

var _ Sender = (*Bot)(nil)

func (b *Bot) Bind(bindFn func(tgbotapi.UpdatesChannel)) {
	bindFn(b.updc)
}

func (b *Bot) WriteMessages(mm ...tgbotapi.MessageConfig) {
	for _, m := range mm {
		ds := defferedShipment{chatID: m.ChatID, cargo: m}
		b.shipToGrid(ds)
	}
}

func (b *Bot) EditMessages(mm ...tgbotapi.MessageConfig) {
	if len(mm) > 1 {
		panic(">=2 messages - not supported now")
	}

	i := 0
	m := mm[i]
	// Заодно и пофиксит накопление лога
	// var dss []defferedShipment
	// for i, m := range mm {
	cid := m.ChatID
	prev := b.shipLog[cid][i]
	em := tgbotapi.NewEditMessageText(cid, prev, m.Text)
	em.ReplyMarkup = m.ReplyMarkup.(*tgbotapi.InlineKeyboardMarkup)
	ds := defferedShipment{chatID: m.ChatID, cargo: em}
	b.shipToGrid(ds)
	// 	dss = append(dss, ds)
	// }
}

func (b *Bot) shipToGrid(ds defferedShipment) {
	c, ok := b.shipGrid[ds.chatID]
	if !ok {
		c = make(chan defferedShipment, 10)
		b.shipGrid[ds.chatID] = c
	}
	if len(c) == 0 {
		go func() {
			timer := time.NewTicker(time.Second)
			for range timer.C {
				select {
				case ds := <-c:
					b.shipHighway <- ds
				default:
					// b.shipGrid[ds.chatID] = nil // THINK
					// delete(b.shipGrid, ds.chatID)
					return
				}
			}
		}()
	}
	c <- ds
}

func (b *Bot) deliver(ds defferedShipment) error {
	b.shipTime[ds.chatID] = time.Now().UnixNano()
	m, err := b.api.Send(ds.cargo)
	b.shipLog[ds.chatID] = append(b.shipLog[ds.chatID], m.MessageID)
	return err
}

// TODO handle errors
func (b *Bot) processMessages() {
	timer := time.NewTicker(time.Second / 30)
	for range timer.C {
		ds := <-b.shipHighway
		b.deliver(ds)
		// select {
		// case cargo := <-b.shipMaster:
		// 	if ok, delta := b.userCanReceiveMessage(cargo.chatID); !ok {
		// 		go func() {
		// 			time.Sleep(time.Duration(delta))
		// 			b.shipBackup <- cargo
		// 		}()
		// 	} else {
		// 		b.sendMessage(cargo)
		// 	}
		// case defferedCargo := <-b.shipBackup:
		// 	b.sendMessage(defferedCargo)
		// }
	}
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
