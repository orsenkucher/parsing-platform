package ppdrop

import (
	"fmt"
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

	} else if state.State.Product.Name == "home" {
		tgmsg = tgbotapi.NewMessage(state.ChatID, state.HomeMsg())
		btms := state.HomeBtm()
		tgmsg.ReplyMarkup = &btms
	} else if state.State.Product.Name == "root" {
		tgmsg = tgbotapi.NewMessage(state.ChatID, state.LocationMsg())
		btms := state.LocationBtm()
		tgmsg.ReplyMarkup = &btms
	} else {
		tgmsg = tgbotapi.NewMessage(state.ChatID, state.BasketMsg())
		btms := state.TreeBtm()
		tgmsg.ReplyMarkup = &btms
	}
	return tgmsg
}

func (state *UsersState) HomeMsg() string {
	return "Ваши закази:"
}

func (state *UsersState) HomeBtm() tgbotapi.InlineKeyboardMarkup {
	rows := [][]tgbotapi.InlineKeyboardButton{}

	for _, basket := range state.Baskets {
		locstr := strconv.FormatUint(basket.Location, 10)
		button := tgbotapi.NewInlineKeyboardButtonData(Locations[basket.Location].Name, "basket\n"+locstr)
		rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
	}

	button := tgbotapi.NewInlineKeyboardButtonData("+", "newbasket\n")
	rows = append(rows, []tgbotapi.InlineKeyboardButton{button})

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (state *UsersState) LocationMsg() string {
	return "Выберите новую локацию на карте:"
}

func (state *UsersState) LocationBtm() tgbotapi.InlineKeyboardMarkup {
	urlbutton := tgbotapi.NewInlineKeyboardButtonURL("🗺 OpenMap", fmt.Sprintf("https://map-bot.abmcloud.com?chatid=%v", state.ChatID))

	rows := [][]tgbotapi.InlineKeyboardButton{[]tgbotapi.InlineKeyboardButton{urlbutton}}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (state *UsersState) BasketMsg() string {
	return state.ToString()
}

func (state *UsersState) TreeBtm() tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	products := make([]*ProdTree, 0, len(state.State.Next))
	for _, prod := range state.State.Next {
		products = append(products, prod)
	}

	rows = state.productButtons(products)
	rows = append(rows, state.lowButtons())

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func (state *UsersState) productButtons(products []*ProdTree) [][]tgbotapi.InlineKeyboardButton {
	rows := [][]tgbotapi.InlineKeyboardButton{}
	basket := state.Baskets[state.Current]
	for _, node := range products {
		text := node.Product.Name
		path := node.GetHash()
		if len(node.Next) > 0 {
			button := tgbotapi.NewInlineKeyboardButtonData(text, "change\n"+path)
			rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
		} else {
			button := tgbotapi.NewInlineKeyboardButtonData(text+" "+strconv.FormatFloat(node.Product.Price, 'f', 2, 64), "\n")
			addButton := tgbotapi.NewInlineKeyboardButtonData("+", "add\n"+path)
			subButton := tgbotapi.NewInlineKeyboardButtonData("-", "sub\n"+path)
			count := 0
			for _, p := range basket.Purchases {
				if p.Product == node {
					count = p.Count
				}
			}
			countButton := tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(count), "\n")

			rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
			rows = append(rows, []tgbotapi.InlineKeyboardButton{subButton, countButton, addButton})
		}
	}
	return rows
}

func (state *UsersState) lowButtons() []tgbotapi.InlineKeyboardButton {
	menu := tgbotapi.NewInlineKeyboardButtonData("🏠", "menu\n")
	location := tgbotapi.NewInlineKeyboardButtonData("🗺", "reset\n")
	basket := tgbotapi.NewInlineKeyboardButtonData("🧺 "+strconv.FormatFloat(state.Baskets[state.Current].Sum, 'f', 2, 64), "basket\n")
	return []tgbotapi.InlineKeyboardButton{menu, location, basket}
}
