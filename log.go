package main

import (
	"fmt"
	"regexp"
	"time"
)

type Link struct {
	Timestamp time.Time
	User      string
	Url       string
	Status    int
}

func NewLink(user, content string) *Link {
	return &Link{Timestamp: time.Now(), User: user, Url: content, Status: 0}
}

type Log struct {
	LinkLog []*Link
}

func (l *Log) AddLink(user, url string) {
	// entweder oder?!
	link := NewLink(user, url)
	l.LinkLog = append(l.LinkLog, link)
	id, err := ctxDb.NewLink(link.User, link.Url, link.Status, link.Timestamp) // in db adden
	if err == false {
		fmt.Printf("log.AddLink: %s\n", "Could not add new link")
	}
	// crawl

	c := &Crawler{Url: url}
	c.Crawl(id)
}

func (l *Log) ParseContent(user, content string) {
	r := regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)`)

	if r.MatchString(content) {
		links := r.FindAllString(content, -1)
		for _, x := range links {
			l.AddLink(user, x)
		}
	}

}
