package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Crawler struct{}

func (c *Crawler) Crawl(id int, url string) {
	// crawl aka download and save in db
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("crawler.Crawl: %s\n", err.Error())
	}
	defer resp.Body.Close()

	src, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("crawler.Crawl: %s\n", err.Error())
		return
	}

	link := new(TblLinks)
	link.Open(id)
	link.Src = src

	id2, err := link.Save()
	if err != nil {
		fmt.Printf("crawler.Crawl: Fehler beim Speichern der Source. %s\n", err.Error())
		return
	}

	if id != id2 {
		fmt.Printf("crawler.Crawl: Eintrag wurde unter einer anderen Id gespeichert \n")
		return
	}

}
