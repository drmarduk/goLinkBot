package main

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type TblLinks struct {
	Id     int
	User   string
	Url    string
	Status int
	Tstamp time.Time
	Src    []byte
	Post   string
	db     *Db
}

func (tbl *TblLinks) Open(id int) error {
	tbl.db = &Db{}
	tbl.db.Open()
	defer tbl.db.Close()
	if id == 0 {
		return errors.New("id is zero.")
	}

	query := "select id, user, url, status, tstamp, src, post from links where id = ?"
	err := tbl.db.con.QueryRow(query, id).Scan(&tbl.Id, &tbl.User, &tbl.Url, &tbl.Status, &tbl.Tstamp,
		&tbl.Src, &tbl.Post)

	if err != nil {
		fmt.Printf("TblLinks.Open: %s\n", err.Error())
		return err
	}
	return nil
}

func (tbl *TblLinks) Save() error {
	var err error
	var result sql.Result
	var query string

	tbl.db = &Db{}
	tbl.db.Open()
	defer tbl.db.Close()

	switch {
	case tbl.Id == 0: // insert
		fmt.Printf("Insert new\n")
		query = "insert into links(id, user, url, status, tstamp, src, post) values(null, ?, ?, ?, ?, ?, ?)"
		result, err = tbl.db.con.Exec(query, tbl.User, tbl.Url, tbl.Status, tbl.Tstamp, tbl.Src, tbl.Post)

	case tbl.Id > 0: // update
		query = "update links set user = ?, url = ?, status = ?, tstamp = ?, src = ?, post = ? where id = ?"
		result, err = tbl.db.con.Exec(query, tbl.User, tbl.Url, tbl.Status, tbl.Tstamp, tbl.Src,
			tbl.Post, tbl.Id)
	}

	if err != nil {
		fmt.Printf("TblLinks.Save: %s\n", err.Error())
		return err
	}

	id, _ := result.LastInsertId()

	tbl.Id = int(id)
	return nil
}

func (tbl *TblLinks) GetTags() []string {
	if tbl.Id == 0 {
		return nil
	}

	tbl.db = &Db{}
	tbl.db.Open()
	defer tbl.db.Close()

	var result []string

	query := "select tag from hastags as h join tags as t on t.id = h.tagid where h.linkid = ?"
	rows, err := tbl.db.con.Query(query, tbl.Id)

	for rows.Next() {
		var tmp string
		err = rows.Scan(&tmp)
		if err != nil {
			fmt.Printf("tblLinks.GetTags: %s\n", err.Error())
			continue
		}
		result = append(result, tmp)
	}
	return result
}

func (tbl *TblLinks) Addtag(tag string) bool {
	tbl.db = &Db{}
	tbl.db.Open()
	defer tbl.db.Close()

	//query := "insert into"
	return false
}

func LinksSearch(q string) (result []TblLinks, err error) {
	tbl := &Db{}
	tbl.Open()
	defer tbl.Close()

	query := "select * from links where src like '%" + q + "%' order by tstamp asc"
	rows, err := tbl.con.Query(query)
	if err != nil {
		fmt.Printf("tblLinks.Search: %s\n", err.Error())
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
