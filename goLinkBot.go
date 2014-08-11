package main

var ctxIrc *Irc
var ctxLog *Log
var ctxDb *Db

func main() {
	// database
	ctxDb = new(Db)
	ctxDb.DbFile = "goLinkBot.db"
	// inmemory linklist
	ctxLog = new(Log)
	// irc bot
	ctxIrc = new(Irc)
	ctxIrc.Channels = append(ctxIrc.Channels, "#g0")
	ctxIrc.Network = "tardis.nerdlife.de"
	ctxIrc.Port = 6697
	ctxIrc.Run()
}
