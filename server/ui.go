package server

import (
	"fmt"
	"sort"
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
	mkp := s.GenerateButtons(query)
	tgmsg.ReplyMarkup = &mkp
	s.Bot.UpdateMsg(tgmsg)
}

func (s *Server) GenerateButtons(query *Query) tgbotapi.InlineKeyboardMarkup {
	fmt.Println(query.State.Product.Name)
	if query.State != s.Tree {
		nodes := make([]*ProdTree, 0, len(query.State.Next)+1)
		for _, v := range query.State.Next {
			nodes = append(nodes, v)
		}

		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Product.Priority < nodes[j].Product.Priority
		})

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
				rows = append(rows, []tgbotapi.InlineKeyboardButton{subButton, button, addButton})
			}
		}

		if back := query.State.Prev; back != nil && back.Product.Name != "root" {
			rows = append(rows, []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData("go back", "change\n"+back.GetPath())})
		}

		return tgbotapi.NewInlineKeyboardMarkup(rows...)
	} else {
		urlbutton := tgbotapi.NewInlineKeyboardButtonURL("OpenMap", fmt.Sprintf("https://scheduleuabot.firebaseapp.com/#/?chatid=%v", query.ChatID))
		//locbutton := tgbotapi.NewKeyboardButtonLocation("Give your Location")

		rows := [][]tgbotapi.InlineKeyboardButton{[]tgbotapi.InlineKeyboardButton{urlbutton}}
		return tgbotapi.NewInlineKeyboardMarkup(rows...)
	}
}
