package main

import (
	"crypto/tls"
	"github.com/thoj/go-ircevent"
	"strconv"
	"strings"
)

type Irc struct {
	Con      *irc.Connection
	Network  string
	Port     int
	Channels []string
}

func (i *Irc) Run() {
	i.Con = irc.IRC("Datenkrake", "Datenkrake")
	i.Con.VerboseCallbackHandler = false
	i.Con.UseTLS = true
	i.Con.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	i.Con.Connect(i.Network + ":" + strconv.Itoa(i.Port))

	i.Con.AddCallback("001", func(e *irc.Event) {
		i.Con.Join(i.Channels[0])
	})
	i.Con.AddCallback("PRIVMSG", parseIrcMsg)
	i.Con.AddCallback("CTCP_ACTION", parseIrcMsg)

	i.Con.Loop()
}

func (i *Irc) WriteToChannel(content string) {
	i.Con.Privmsg(i.Channels[0], content)
}

func parseIrcMsg(e *irc.Event) {
	var user, content string
	user = e.Nick
	content = e.Arguments[1]
	text := strings.Split(content, " ")

	// adds link to db and crawls it
	ctxLog.ParseContent(user, content) // generell immer machen bei link
	switch {
	case text[1] == "!search":
		search(strings.Join(text[2:], " "))
	case text[1] == "!linkinfo":
		id, err := strconv.Atoi(text[1])
		if err != nil {
			return
		}
		linkinfo(id)
	case text[1] == "!addtag":
		id, err := strconv.Atoi(text[1])
		if err != nil {
			return
		}
		for _, t := range text[2:] {
			addTag(id, t, user)
		}
	default:
		return
	}
}
