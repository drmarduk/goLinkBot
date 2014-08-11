package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"time"
)

type TblLinks struct {
	Id     int
	User   string
	Url    string
	Post   string
	Status int
	Tstamp time.Time
	Src    []byte
}

func (db *Db) AddLink(user, url, post string, status int, timestmap time.Time) (int64, bool) {
	var err error
	var result sql.Result
	db.Open()
	defer db.Close()

	// escape
	user = template.HTMLEscapeString(user)
	url = template.HTMLEscapeString(url)
	post = template.HTMLEscapeString(post)

	query := "insert into links (user, url, post, tstamp, status) values(?, ?, ?, ?, ?);"
	result, err = db.con.Exec(query, user, url, post, timestmap, status)
	if err != nil {
		fmt.Printf("db.NewLink: %s\n", err.Error())
		return 0, false
	}

	id, _ := result.LastInsertId()

	return id, true
}

func (db *Db) UpdateSource(id int64, src []byte) bool {
	var err error
	db.Open()
	defer db.Close()

	query := "update links set status = 1, src = ? where id = ?"
	_, err = db.con.Exec(query, src, id)
	if err != nil {
		fmt.Printf("db.UpdateSource: %s\n", err.Error())
		return false
	}
	return true
}

func (db *Db) SearchBlob(q string) (result []TblLinks, err error) {
	var rows *sql.Rows
	db.Open()
	defer db.Close()

	query := "select * from links where src like '%" + q + "%' order by tstamp asc"
	rows, err = db.con.Query(query)
	if err != nil {
		fmt.Printf("db.SearchBlob: %s\n", err.Error())
		return nil, err
	}

	for rows.Next() {
		l := TblLinks{}
		err = rows.Scan(&l.Id, &l.User, &l.Url, &l.Status,
			&l.Tstamp, &l.Src, &l.Post)

		if err != nil {
			fmt.Printf("db.Search: %s\n", err.Error())
			continue
		}
		result = append(result, l)
	}
	return result, nil
}
