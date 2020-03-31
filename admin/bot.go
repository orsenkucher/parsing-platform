package admin

import (
	"fmt"
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
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			}
			delcfg := tgbotapi.NewDeleteMessage(update.Message.Chat.ID, update.Message.MessageID)
			if _, err := b.api.DeleteMessage(delcfg); err != nil {
				log.Println(err)
			}
		}
		b.updc <- update
	}
}

var _ Sender = (*Bot)(nil)

func (b *Bot) Bind(bindFn func(tgbotapi.UpdatesChannel)) {
	bindFn(b.updc)
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
	log := b.shipLog[cid]
	prev := log[len(log)-1]
	// em := tgbotapi.NewEditMessageText(cid, prev, m.Text)
	// if m.ReplyMarkup != nil {
	// 	em.ReplyMarkup = m.ReplyMarkup.(*tgbotapi.InlineKeyboardMarkup)
	// }
	dm := tgbotapi.NewDeleteMessage(cid, prev)
	b.shipToGrid(defferedShipment{chatID: m.ChatID, cargo: dm})
	b.shipToGrid(defferedShipment{chatID: m.ChatID, cargo: m})
	// 	dss = append(dss, ds)
	// }
}

func (b *Bot) WriteMessages(mm ...tgbotapi.MessageConfig) {
	for _, m := range mm {
		ds := defferedShipment{chatID: m.ChatID, cargo: m}
		fmt.Println("ADDED")
		b.shipToGrid(ds)
		// time.Sleep(time.Second / 4)
		// time.Sleep(time.Second - time.Millisecond*700)
		// time.Sleep(time.Millisecond)
	}
}

func (b *Bot) shipToGrid(ds defferedShipment) {
	c, ok := b.shipGrid[ds.chatID]
	if !ok {
		c = make(chan defferedShipment, 10)
		b.shipGrid[ds.chatID] = c
		fmt.Println("NEWCHANNEL")
	}
	// time.Sleep(time.Millisecond * 10) // DONT
	c <- ds
	if !ok {
		// c <- ds
		go func() {
			for {
				// fmt.Println("IN FOR")
				//wg.wait
				select {
				case ds := <-c:
					fmt.Println("DS CASE")
					ok, delay := b.ready(ds.chatID)
					if !ok {
						fmt.Println("SLEEP", time.Duration(delay))
						time.Sleep(time.Duration(delay))
					}
					fmt.Println("TO HIGHWAY")
					b.shipTime[ds.chatID] = time.Now().UnixNano()
					b.shipHighway <- ds
				default:
					delete(b.shipGrid, ds.chatID)
					close(c) // You'd better not send on this chan
					fmt.Println("\tEXITED")
					return
				}
			}
		}()
	} // else {
	// fmt.Println("TOCHANNEL")
	// if len(c) == 0 // USE waitGroups?
	// fmt.Println("TOCHANNEL")
	// time.Sleep(time.Millisecond)
	// wg.one
	// time.Sleep(time.Second * 2) // WILL SKIP ALL
	// c <- ds // WHAT IF c is voided??? TODO
	// PANIC for now
	// YEEEEA(((
	// wg.release
	// }
}

// Grid is ready to deliver user's cargo
func (b *Bot) ready(chatID int64) (ok bool, delta int64) {
	if t, ok := b.shipTime[chatID]; ok {
		delta = int64(time.Second/3) + t - time.Now().UnixNano()
		return delta <= 0, delta
	}
	return true, 0
}

// TODO handle errors
func (b *Bot) processMessages() {
	timer := time.NewTicker(time.Second / 30)
	for range timer.C {
		// fmt.Println("TICK")
		ds := <-b.shipHighway
		err := b.deliver(ds)
		if err != nil {
			log.Println(err)
		}
	}
}

func (b *Bot) deliver(ds defferedShipment) error {
	// b.shipTime[ds.chatID] = time.Now().UnixNano()
	m, err := b.api.Send(ds.cargo)
	b.shipLog[ds.chatID] = append(b.shipLog[ds.chatID], m.MessageID)
	return err
}
