package main

import (
	"database/sql"
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
	query := "insert into hastags(linkid, tagid, user, tstamp) values(?, ?, ?, ?)"
	result, err := db.con.Exec(query, linkid, tagid, user, time.Now())
	if err != nil {
		fmt.Printf("db.AddTagToLink: %s\n", err.Error())
		return false
	}
	return true
}

/*
	Liefert alle Tags für einen speziellen Link zurück
*/
func (db *Db) GetTagsFromLink(linkid int) (result []TblHasTags, error) {
	query := "select hastags.user, hastags.tstamp, tags.id, tags.tag from hastags join tags on hastags.tagid = tags.id where hastags.linkid = ?"
	rows, err := db.con.Query(query, linkid)
	if err != nil {
		fmt.Printf("db.GetTagsFromLink: %s\n", err.Error())
		return nil, false
	}

	for rows.Next() {
		r := &TblHasTags{}
		err = rows.Scan(&r.User, &r.Tstamp, &r.TagId, &r.Post)
		if err != nil {
			fmt.Printf("db.GetTagsfromLink: %s\n", err.Error())
			continue
		}
		result = append(result, r)
	}


}

func (db *Db) GetLinksFromTags(tagid int) []string {
	query := "select * from hastags join links on hastags.linkid = links.id where hastags.tagid = ?"
}
