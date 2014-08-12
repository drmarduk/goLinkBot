package main

var ctxIrc *Irc
var ctxLog *Log

func main() {
	// inmemory linklist
	ctxLog = new(Log)

	//go crawl

	// irc bot
	ctxIrc = new(Irc)
	ctxIrc.Channels = append(ctxIrc.Channels, "#g0")
	ctxIrc.Network = "tardis.nerdlife.de"
	ctxIrc.Port = 6697
	ctxIrc.Run()
}
