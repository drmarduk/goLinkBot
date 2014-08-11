package main

import (
	"database/sql"
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
	tbl.db = new(Db)
	tbl.db.Open()
	defer tbl.db.Close()
	query := "select id, user, url, status, tstamp, src, post from links where id = ?"
	err := tbl.db.con.QueryRow(query, id).Scan(&tbl.Id, &tbl.User, &tbl.Url, &tbl.Status, &tbl.Tstamp,
		&tbl.Src, &tbl.Post)

	if err != nil {
		fmt.Printf("TblLinks.Open: %s\n", err.Error())
		return err
	}
	return nil
}

func (tbl *TblLinks) Save() (int, error) {
	var err error
	var result *sql.Result
	tbl.db = new(Db)
	tbl.db.Open()
	defer tbl.db.Close()
	var query string
	if tbl.Id == 0 { // insert
		query = "insert into links(user, url, status, tstamp, src, post) values(?, ?, ?, ?, ?, ?)"
		result, err := tbl.db.con.Exec(query, tbl.User, tbl.Url, tbl.Status, tbl.Tstamp, tbl.Src,
			tbl.Post)
	}
	if tbl.Id > 0 { // update
		query = "update links set user = ?, url = ?, status = ?, tstamp = ?, src = ?, post = ? where id = ?"
		result, err := tbl.db.con.Exec(query, tbl.User, tbl.Url, tbl.Status, tbl.Tstamp, tbl.Src,
			tbl.Post, tbl.Id)
	}

	if err != nil {
		fmt.Printf("TblLinks.Save: %s\n", err.Error())
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("TblLinks.Save: %s\n", err.Error())
		return 0, err
	}
	return int(id), nil
}

func /*(tbl *TblLinks) */ LinksSearch(q string) (result []TblLinks, err error) {
	tbl := new(&Db)
	tbl.Open()
	defer tbl.Close()

	query := "select * from links where src like '%" + q + "%' order by tstamp asc"
	rows, err = tbl.con.Query(query)
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
