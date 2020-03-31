package ppdrop

import (
	"fmt"
	"sort"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type UsersState struct {
	State   *ProdTree
	Current uint64
	Baskets map[uint64]*Basket
	ChatID  int64
}

func (state *UsersState) GenerateMsg() tgbotapi.MessageConfig {
	fmt.Println("State: ", state.State.Product.Name)
	var tgmsg tgbotapi.MessageConfig
	if state.State.Product.Name == "basket" {
		tgmsg = tgbotapi.NewMessage(state.ChatID, state.BasketMsg())
		Btns := state.BasketBtn()
		tgmsg.ReplyMarkup = &Btns
	} else if state.State.Product.Name == "home" {
		tgmsg = tgbotapi.NewMessage(state.ChatID, state.HomeMsg())
		Btns := state.HomeBtn()
		tgmsg.ReplyMarkup = &Btns
	} else if state.State.Product.Name == "root" {
		tgmsg = tgbotapi.NewMessage(state.ChatID, state.LocationMsg())
		Btns := state.LocationBtn()
		tgmsg.ReplyMarkup = &Btns
	} else if state.State.Product.Name == "agree" {
		tgmsg = tgbotapi.NewMessage(state.ChatID, state.AgreeMsg())
		Btns := state.AgreeBtn()
		tgmsg.ReplyMarkup = &Btns
	} else if state.Baskets[state.Current].Status != New {
		tgmsg = tgbotapi.NewMessage(state.ChatID, state.BasketStatusMsg())
		Btns := state.BasketStatusBtn()
		tgmsg.ReplyMarkup = &Btns
	} else {
		tgmsg = tgbotapi.NewMessage(state.ChatID, state.BasketMsg())
		Btns := state.TreeBtn()
		tgmsg.ReplyMarkup = &Btns
	}
	return tgmsg
}

func (state *UsersState) BasketStatusMsg() string {
	return state.ToString()
}

func (state *UsersState) BasketStatusBtn() tgbotapi.InlineKeyboardMarkup {
	rows := [][]tgbotapi.InlineKeyboardButton{}

	home := tgbotapi.NewInlineKeyboardButtonData("В меню", "home\n")
	reset := tgbotapi.NewInlineKeyboardButtonData("Отменить", "reset\n"+strconv.FormatUint(state.Current, 10))

	rows = append(rows, []tgbotapi.InlineKeyboardButton{home})
	rows = append(rows, []tgbotapi.InlineKeyboardButton{reset})

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (state *UsersState) AgreeMsg() string {
	return "Вы уверенны, что хотите выйти в главное меню?\n При выходе из неотправленной корзины, он удалится!!!"
}

func (state *UsersState) AgreeBtn() tgbotapi.InlineKeyboardMarkup {
	rows := [][]tgbotapi.InlineKeyboardButton{}

	catalog := tgbotapi.NewInlineKeyboardButtonData("Вернуться к заказу", "catalog\n")
	home := tgbotapi.NewInlineKeyboardButtonData("В меню", "home\n")
	rows = append(rows, []tgbotapi.InlineKeyboardButton{catalog})
	rows = append(rows, []tgbotapi.InlineKeyboardButton{home})

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (state *UsersState) HomeMsg() string {
	return "Ваши заказы:"
}

func (state *UsersState) HomeBtn() tgbotapi.InlineKeyboardMarkup {
	rows := [][]tgbotapi.InlineKeyboardButton{}

	for _, basket := range state.Baskets {
		locstr := strconv.FormatUint(basket.Location, 10)
		button := tgbotapi.NewInlineKeyboardButtonData(Locations[basket.Location].Name, "location\n"+locstr)
		rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
	}

	button := tgbotapi.NewInlineKeyboardButtonData("➕ Новый заказ", "newbasket\n")
	rows = append(rows, []tgbotapi.InlineKeyboardButton{button})

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (state *UsersState) LocationMsg() string {
	return "Выберите магазин:"
}

func (state *UsersState) LocationBtn() tgbotapi.InlineKeyboardMarkup {
	urlbutton := tgbotapi.NewInlineKeyboardButtonURL("🗺 На карте", fmt.Sprintf("https://map-bot.abmcloud.com/#/?chatid=%v", state.ChatID))
	listbutton := tgbotapi.NewInlineKeyboardButtonData("📖 Из списка", "home\n")
	location := tgbotapi.NewInlineKeyboardButtonData("🏠 Назад", "home\n")

	rows := [][]tgbotapi.InlineKeyboardButton{[]tgbotapi.InlineKeyboardButton{urlbutton, listbutton}, []tgbotapi.InlineKeyboardButton{location}}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (state *UsersState) BasketMsg() string {
	return state.ToString()
}

func (state *UsersState) BasketBtn() tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton
	basket := state.Baskets[state.Current]

	products := make([]*ProdTree, 0, len(basket.Purchases))
	for _, purch := range basket.Purchases {
		products = append(products, purch.Product)
	}

	rows = state.productButtons(products)
	rows = append(rows, state.lowButtons()...)

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (state *UsersState) TreeBtn() tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	products := make([]*ProdTree, 0, len(state.State.Next))
	for _, prod := range state.State.Next {
		products = append(products, prod)
	}

	rows = state.productButtons(products)
	rows = append(rows, state.lowButtons()...)

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (state *UsersState) productButtons(products []*ProdTree) [][]tgbotapi.InlineKeyboardButton {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Product.Priority < products[j].Product.Priority
	})
	rows := [][]tgbotapi.InlineKeyboardButton{}
	basket := state.Baskets[state.Current]
	for _, node := range products {
		text := node.Product.Name
		path := node.GetHash()
		if len(node.Next) > 0 {
			button := tgbotapi.NewInlineKeyboardButtonData(text, "change\n"+path)
			rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
		} else {
			button := tgbotapi.NewInlineKeyboardButtonData(text, "\n")
			addButton := tgbotapi.NewInlineKeyboardButtonData("➕", "add\n"+path)
			price := tgbotapi.NewInlineKeyboardButtonData(strconv.FormatFloat(node.Product.Price, 'f', 2, 64), "\n")
			subButton := tgbotapi.NewInlineKeyboardButtonData("➖", "sub\n"+path)
			count := 0
			for _, p := range basket.Purchases {
				if p.Product == node {
					count = p.Count
				}
			}
			countButton := tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(count)+" шт", "\n")

			rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
			rows = append(rows, []tgbotapi.InlineKeyboardButton{subButton, price, countButton, addButton})
		}
	}
	if back := state.State.Prev; back != nil && back.Product.Name != "root" {
		rows = append(rows, []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData("Выбрать другую категорию", "change\n"+back.GetHash())})
	}
	return rows
}

func (state *UsersState) lowButtons() [][]tgbotapi.InlineKeyboardButton {
	rows := [][]tgbotapi.InlineKeyboardButton{}
	menu := tgbotapi.NewInlineKeyboardButtonData("📖", "catalog\n")
	location := tgbotapi.NewInlineKeyboardButtonData("🏠", "agree\n")
	basket := tgbotapi.NewInlineKeyboardButtonData("🧺 "+strconv.FormatFloat(state.Baskets[state.Current].Sum, 'f', 2, 64), "basket\n")
	if state.State.Product.Name == "basket" {
		basket = tgbotapi.NewInlineKeyboardButtonData("Отправить заказ ✅", "sendbasket\n")
		rows = append(rows, []tgbotapi.InlineKeyboardButton{basket})
		rows = append(rows, []tgbotapi.InlineKeyboardButton{menu, location})
	} else {
		rows = append(rows, []tgbotapi.InlineKeyboardButton{menu, location, basket})
	}
	return rows
}
