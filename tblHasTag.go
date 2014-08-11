package main

import (
	_ "database/sql"
	"fmt"
	"time"
)

type TblHasTags struct {
	LinkId int
	TagId  int
	User   string
	Tstamp time.Time
	Post   string
}

/*
	Fügt einem Link ein Tag hinzu
*/
func (db *Db) AddTagToLink(linkid, tagid int, user string) bool {
	db.Open()
	defer db.Close()
	query := "insert into hastags(linkid, tagid, user, tstamp) values(?, ?, ?, ?)"
	_, err := db.con.Exec(query, linkid, tagid, user, time.Now())
	if err != nil {
		fmt.Printf("db.AddTagToLink: %s\n", err.Error())
		return false
	}
	return true
}

/*
	Liefert alle Tags für einen speziellen Link zurück
*/
func (db *Db) GetTagsFromLink(linkid int) (result []TblHasTags, err error) {
	db.Open()
	defer db.Close()
	query := "select hastags.user, hastags.tstamp, tags.id, tags.tag from hastags join tags on hastags.tagid = tags.id where hastags.linkid = ?"
	rows, err := db.con.Query(query, linkid)
	if err != nil {
		fmt.Printf("db.GetTagsFromLink: %s\n", err.Error())
		return nil, err
	}

	for rows.Next() {
		r := TblHasTags{}
		err = rows.Scan(&r.User, &r.Tstamp, &r.TagId, &r.Post)
		if err != nil {
			fmt.Printf("db.GetTagsfromLink: %s\n", err.Error())
			continue
		}
		result = append(result, r)
	}

	return result, nil
}

func (db *Db) GetLinksFromTags(tagid int) (result []TblLinks, err error) {
	db.Open()
	defer db.Close()
	query := "select links.id, links.user, links.url, links.tstamp from hastags join links on hastags.linkid = links.id where hastags.tagid = ?"
	rows, err := db.con.Query(query, tagid)
	if err != nil {
		fmt.Printf("db.GetLinks.FromTags: %s\n", err.Error())
		return nil, err
	}

	for rows.Next() {
		r := TblLinks{}
		err = rows.Scan(&r.Id, &r.User, &r.Url, &r.Tstamp)
		if err != nil {
			fmt.Printf("db.GetLinksFromTags: %s\n", err.Error())
			continue
		}

		result = append(result, r)
	}
	return result, nil

	return result, nil

}
