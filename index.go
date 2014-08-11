package main

import (
	_ "database/sql"
	"fmt"
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
	return id == id
}

func addTag(id int, tag string) bool {
	return id == id && tag == tag

}
