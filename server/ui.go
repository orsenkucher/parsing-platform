package server

import (
	"fmt"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (s *Server) ReloadMsg(ChatID int64) {
	query := s.GetQuery(ChatID)
	var tgmsg tgbotapi.MessageConfig
	if query.State != s.Tree {
		text := query.Location + "\n" + query.ToString()
		tgmsg = tgbotapi.NewMessage(ChatID, text)
	} else {
		text := "Choose your store"
		tgmsg = tgbotapi.NewMessage(ChatID, text)
	}
	mkp := s.GenerateButtons(query.State)
	tgmsg.ReplyMarkup = &mkp
	s.Bot.UpdateMsg(tgmsg)
}

func (s *Server) GenerateButtons(state *ProdTree) tgbotapi.InlineKeyboardMarkup {
	fmt.Println(state.Product.Name)
	if state != s.Tree {
		nodes := make([]*ProdTree, 0, len(state.Next)+1)
		for _, v := range state.Next {
			nodes = append(nodes, v)
		}

		rows := [][]tgbotapi.InlineKeyboardButton{}
		for _, node := range nodes {
			text := node.Product.Name
			path := node.GetPath()
			if len(node.Next) > 0 {
				button := tgbotapi.NewInlineKeyboardButtonData(text, "change\n"+path)
				rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
			} else {
				button := tgbotapi.NewInlineKeyboardButtonData(text+" "+strconv.FormatFloat(node.Product.Price, 'f', 2, 64), "\n"+path)
				addButton := tgbotapi.NewInlineKeyboardButtonData("+", "add\n"+path)
				subButton := tgbotapi.NewInlineKeyboardButtonData("-", "sub\n"+path)
				rows = append(rows, []tgbotapi.InlineKeyboardButton{addButton, button, subButton})
			}
		}

		if back := state.Prev; back != nil && back.Product.Name != "root" {
			rows = append(rows, []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData(back.Product.Name, "change\n"+back.GetPath())})
		}

		return tgbotapi.NewInlineKeyboardMarkup(rows...)
	} else {
		button := tgbotapi.NewInlineKeyboardButtonData("OpenMap", "\n")
		rows := [][]tgbotapi.InlineKeyboardButton{[]tgbotapi.InlineKeyboardButton{button}}
		return tgbotapi.NewInlineKeyboardMarkup(rows...)
	}
}
