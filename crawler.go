package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Crawler struct {
}

func (c *Crawler) Crawl(id int64, url string) {
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

	ctxDb.UpdateSource(id, src)
}
