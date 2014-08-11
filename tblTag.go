package main

import (
	"database/sql"
	"fmt"
)

type TblTags struct {
	Id  int
	Tag string
}

/*
	Fügt einen neuen Tag ein und gibt die Id zurück
*/
func (db *Db) NewTag(tag string) int64 {
	db.Open()
	defer db.Close()
	query := "insert into tags(tag) value(?)"
	result, err := db.con.Exec(query, tag)
	if err != nil {
		fmt.Printf("db.NewTag: %s\n", err.Error())
		return 0
	}
	id, _ := result.LastInsertId()
	return id
}

/*
	Liefert die Id zu einem speziellen Tag zurück
*/
func (db *Db) GetTagId(tag string) int64 {
	db.Open()
	defer db.Close()
	query := "select id from tags where tag = ?"
	var id int
	err := db.con.QueryRow(query, tag).Scan(&id)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("db.GetTagId: No rows found.\n")
		return db.NewTag(tag) // bei 0 recodrds diretk einen einfügen // kA ob das geht
		return 0
	case err != nil:
		fmt.Printf("db.GetTagId: %s\n!", err.Error())
		return 0
	default:
		return int64(id) // on success
	}
}

/*func (db *Db) DeleteTag(id int) bool {
	query1 := "delete from tags where id = ?"
	query2 := "delete from hastags where tagid = ?"
}
*/
