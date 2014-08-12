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
	// orm db obj
	link := &TblLinks{}
	link.User = user
	link.Url = url
	link.Post = content
	link.Status = 0
	link.Tstamp = time.Now()

	/*id, */ err := link.Save()
	if err != nil {
		fmt.Printf("log.AddLink: %s\n", "Could not add new link")
	}

	fmt.Printf("%s added: %s\n", user, url)
	// crawl
	c := &Crawler{} // neu machen, mit channels evtl
	c.Crawl(link.Id, url)
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
