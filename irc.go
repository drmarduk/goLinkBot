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

	ctxLog.ParseContent(user, content)

	if strings.Contains(content, "!search") {
		search(strings.SplitAfter(content, "!search")[1])
	}
}
