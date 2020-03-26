package server

import (
	"fmt"
	"sort"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (s *Server) ReloadMsg(ChatID int64) {
	query := s.GetQuery(ChatID)
	if query.State != nil {
		var tgmsg tgbotapi.MessageConfig
		if query.State != s.Tree {
			text := "-------------------------------------\n" + query.Location + "\n" + query.ToString()
			tgmsg = tgbotapi.NewMessage(ChatID, text)
		} else {
			text := "Choose your store"
			tgmsg = tgbotapi.NewMessage(ChatID, text)
		}
		mkp := s.ReloadButtons(query)
		tgmsg.ReplyMarkup = &mkp
		s.Bot.UpdateMsg(tgmsg)
	} else {
		s.ShowBasket(query)
	}
}

func (s *Server) ReloadButtons(query *Query) tgbotapi.InlineKeyboardMarkup {
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
				button := tgbotapi.NewInlineKeyboardButtonData(text+" "+strconv.FormatFloat(node.Product.Price, 'f', 2, 64), "\n")
				addButton := tgbotapi.NewInlineKeyboardButtonData("+", "add\n"+path)
				subButton := tgbotapi.NewInlineKeyboardButtonData("-", "sub\n"+path)
				count := 0
				for _, p := range query.Purchases {
					if p.Product == node {
						count = p.Count
					}
				}
				countButton := tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(count), "\n")

				rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
				rows = append(rows, []tgbotapi.InlineKeyboardButton{subButton, countButton, addButton})
			}
		}

		if back := query.State.Prev; back != nil && back.Product.Name != "root" {
			//if back := query.State.Prev; back != nil {
			rows = append(rows, []tgbotapi.InlineKeyboardButton{tgbotapi.NewInlineKeyboardButtonData("go back", "change\n"+back.GetPath())})
		}
		rows = append(rows, LowButtoms(query))

		return tgbotapi.NewInlineKeyboardMarkup(rows...)
	} else {
		urlbutton := tgbotapi.NewInlineKeyboardButtonURL("OpenMap", fmt.Sprintf("https://scheduleuabot.firebaseapp.com/#?chatid=%v", query.ChatID))
		//locbutton := tgbotapi.NewKeyboardButtonLocation("Give your Location")

		rows := [][]tgbotapi.InlineKeyboardButton{[]tgbotapi.InlineKeyboardButton{urlbutton}}
		return tgbotapi.NewInlineKeyboardMarkup(rows...)
	}
}

func (s *Server) ShowBasket(query *Query) {
	var tgmsg tgbotapi.MessageConfig
	if query.State != s.Tree {
		text := "-------------------------------------\n" + query.Location + "\n" + query.ToString()
		tgmsg = tgbotapi.NewMessage(query.ChatID, text)
	} else {
		text := "Choose your store"
		tgmsg = tgbotapi.NewMessage(query.ChatID, text)
	}
	mkp := s.ShowBasketButtons(query)
	tgmsg.ReplyMarkup = &mkp
	s.Bot.UpdateMsg(tgmsg)
}

func (s *Server) ShowBasketButtons(query *Query) tgbotapi.InlineKeyboardMarkup {
	fmt.Println("basket")
	if query.State != s.Tree {
		nodes := make([]*Purchase, 0, len(query.Purchases))
		for _, v := range query.Purchases {
			if v.Count > 0 {
				nodes = append(nodes, v)
			}
		}

		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Product.Product.Priority < nodes[j].Product.Product.Priority
		})

		rows := [][]tgbotapi.InlineKeyboardButton{}
		for _, node := range nodes {
			text := node.Product.Product.Name
			path := node.Product.GetPath()
			button := tgbotapi.NewInlineKeyboardButtonData(text+" "+strconv.FormatFloat(node.Product.Product.Price, 'f', 2, 64), "\n"+path)
			addButton := tgbotapi.NewInlineKeyboardButtonData("+", "add\n"+path)
			subButton := tgbotapi.NewInlineKeyboardButtonData("-", "sub\n"+path)
			count := 0
			for _, p := range query.Purchases {
				if p == node {
					count = p.Count
				}
			}
			countButton := tgbotapi.NewInlineKeyboardButtonData(strconv.Itoa(count), "\n")
			rows = append(rows, []tgbotapi.InlineKeyboardButton{button})
			rows = append(rows, []tgbotapi.InlineKeyboardButton{subButton, countButton, addButton})
		}
		rows = append(rows, LowButtoms(query))

		return tgbotapi.NewInlineKeyboardMarkup(rows...)
	} else {
		urlbutton := tgbotapi.NewInlineKeyboardButtonURL("OpenMap", fmt.Sprintf("https://scheduleuabot.firebaseapp.com/#?chatid=%v", query.ChatID))
		//locbutton := tgbotapi.NewKeyboardButtonLocation("Give your Location")

		rows := [][]tgbotapi.InlineKeyboardButton{[]tgbotapi.InlineKeyboardButton{urlbutton}}
		return tgbotapi.NewInlineKeyboardMarkup(rows...)
	}
}

func LowButtoms(q *Query) []tgbotapi.InlineKeyboardButton {
	menu := tgbotapi.NewInlineKeyboardButtonData("🏠", "menu\n")
	location := tgbotapi.NewInlineKeyboardButtonData("📍", "reset\n")
	basket := tgbotapi.NewInlineKeyboardButtonData("🧺 "+strconv.FormatFloat(q.Sum, 'f', 2, 64), "basket\n")
	return []tgbotapi.InlineKeyboardButton{menu, location, basket}
}
