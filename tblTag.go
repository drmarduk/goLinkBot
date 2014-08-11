package main

import (
	"database/sql"
	"fmt"
)

type TblTags struct {
	Id  int
	Tag string

	db *sql.DB
}

func (tbl *TblTags) Open(id int) error {
	tbl.db = new(*Db)
	tbl.db.Open()
	defer tbl.db.Close()
	query := "select id, tag from tags where id = ?"
	err := tbl.db.QueryRow(query, id).Scan(&tbl.Id, &tbl.Tag)
	if err != nil {
		fmt.Printf("TblTags.Open: %s\n", err.Error())
		return err
	}
	return nil
}

func (tbl *TblTags) Save() (int, error) {
	tbl.db = new(*Db)
	tbl.db.Open()
	defer tbl.db.Close()
	var result *sql.Result

	if tbl.Id == 0 {
		query := "insert into tags(tag) values(?)"
		result, err = tbl.db.Exec(query, tbl.Tag)
	}
	if tbl.id > 0 {
		query := "update tags set tag = ? where id = ?"
		result, err = tbl.db.Exec(query, tbl.Tag, tbl.Id)
	}

	if err != nil {
		fmt.Printf("TblTags.Save: %s\n", err.Error())
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Printf("TblTags.Save: %s\n", err.Error())
		return 0, err
	}
	return int(id), nil
}

func (tbl *TblTags) Search(q string) ([]TblTags, error) {
	tbl.db = new(*Db)
	tbl.db.Open()
	defer tbl.db.Close()
	var result []TblTags = make([]TblTags)
	query := "select id, tag from tags where tag like '%" + q + "%'"

	rows, err := tbl.db.Query(query)

	for rows.Next() {
		t := TblTags{}
		err := rows.Scan(&t.Id, &t.Tag)
		if err != nil {
			fmt.Printf("TblTag.Search: %s\n", err.Error())
			continue
		}

		result = append(result, t)
	}

	return result, nil
}
