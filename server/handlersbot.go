package server

func (b *Bot) NewLocation(chatid int64, loc string) {
	nl := NewLocation{Location: loc, ChatID: chatid}
	// if b.Updates == nil {
	// 	fmt.Println("asfsdsgs")
	// }
	b.Updates <- nl
}
