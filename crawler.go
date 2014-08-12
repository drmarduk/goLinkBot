package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Crawler struct{}

func (c *Crawler) Crawl(id int, url string) {
	// skip focking cert validation -.-
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	// crawl aka download and save in db
	resp, err := client.Get(url)
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

	err := link.Save()
	if err != nil {
		fmt.Printf("crawler.Crawl: Fehler beim Speichern der Source. %s\n", err.Error())
		return
	}

}

type CrawlObject struct {
	Linkid int
	Url string
}
func crawlchan(chan cs CrawlObject) {

	for {
		select obj, chan_ok := <-cs {
			if chan_ok {
				// chan ok

			} else {
				// chan closed
			}
		}
	}

}
