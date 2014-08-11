package main

import (
	"fmt"
	"regexp"
	"time"
)

type Log struct {
	//LinkLog []*Link
}

func (l *Log) AddLink(user, url, content string) {
	id, err := ctxDb.AddLink(user, url, content, 0, time.Now()) // in db adden
	if err == false {
		fmt.Printf("log.AddLink: %s\n", "Could not add new link")
	}
	// crawl
	c := &Crawler{} // neu machen, mit channels evtl
	c.Crawl(id, url)
}

func (l *Log) ParseContent(user, content string) {
	r := regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)`)

	if r.MatchString(content) {
		links := r.FindAllString(content, -1)
		for _, x := range links {
			l.AddLink(user, x, content)
		}
	}

}
