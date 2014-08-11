package main

import (
	_ "database/sql"
	"fmt"
)

func search(query string) bool {
	links, err := ctxDb.SearchBlob(query)
	if err != nil {
		fmt.Printf("index.Search: %s\n", err.Error())
		return false
	}

	for _, l := range links {
		ctxIrc.WriteToChannel(fmt.Sprintf("[%s] {%s} %s: %s", l.Tstamp.Format("02.01.2006 15:04:05"), l.Id, l.User, l.Url))
	}
	return true
}
