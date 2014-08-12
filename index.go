package main

import (
	_ "database/sql"
	"fmt"
	"strings"
)

func search(query string) bool {
	links, err := LinksSearch(query)
	if err != nil {
		fmt.Printf("index.Search: %s\n", err.Error())
		return false
	}

	for _, l := range links {
		ctxIrc.WriteToChannel(fmt.Sprintf("[%s] {%s} %s: %s", l.Tstamp.Format("02.01.2006 15:04:05"), l.Id, l.User, l.Url))
	}
	return true
}

func linkinfo(id int) bool {
	l := &TblLinks{}
	err := l.Open(id)
	if err != nil {
		fmt.Printf("index.linkinfo: Fehler beim Öffnen von Orm Objekt. %s\n", err.Error())
		return false
	}

	ctxIrc.WriteToChannel(fmt.Sprintf("Id: %s", l.Id))
	ctxIrc.WriteToChannel(fmt.Sprintf("Link: %s", l.Url))
	ctxIrc.WriteToChannel(fmt.Sprintf("User: %s", l.User))
	ctxIrc.WriteToChannel(fmt.Sprintf("Timestamp: %s", l.Tstamp.Format("02.01.2006 15:04:05")))
	ctxIrc.WriteToChannel(fmt.Sprintf("Original Message: '%s'", l.Post))
	ctxIrc.WriteToChannel(fmt.Sprintf("Tags: %s", strings.Join(l.GetTags(), ", ")))

	return true
}

func addTag(id int, tag string) bool {
	return id == id && tag == tag

}
