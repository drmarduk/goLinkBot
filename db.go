package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Db struct {
	err    error
	DbFile string
	con    *sql.DB
}

func (db *Db) Open() {
	db.con, db.err = sql.Open("sqlite3", db.DbFile)
	if db.err != nil {
		fmt.Printf("db.NewLink: %s\n", db.err.Error())
	}
}

func (db *Db) Close() {
	db.con.Close()
}
