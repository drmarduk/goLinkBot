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

	// adds link to db and crawls it
	ctxLog.ParseContent(user, content)

	// command handling
	if strings.Contains(content, "!search ") {
		var query string = strings.Replace(content, "!search ", "", 1)

		search(query)
	}

	if strings.Contains(content, "!linkinfo ") {
		id, _ := strconv.Atoi(strings.Replace(content, "!linkinfo ", "", 1))
		linkinfo(id)
	}

	if strings.Contains(content, "!addTag ") {
		arr := strings.Split(content, " ")
		id, err := strconv.Atoi(arr[1])
		if err != nil {
			return
		}

		for _, t := range arr[2:] {
			addTag(id, t, user)
		}
	}
}
