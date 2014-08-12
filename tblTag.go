package main

import (
	"database/sql"
	"fmt"
)

type TblTags struct {
	Id  int
	Tag string

	db *Db
}

func (tbl *TblTags) Open(id int) error {
	tbl.db = &Db{}
	tbl.db.Open()
	defer tbl.db.Close()
	query := "select id, tag from tags where id = ?"
	err := tbl.db.con.QueryRow(query, id).Scan(&tbl.Id, &tbl.Tag)
	if err != nil {
		fmt.Printf("TblTags.Open: %s\n", err.Error())
		return err
	}
	return nil
}

func (tbl *TblTags) Save() error {
	tbl.db = &Db{}
	tbl.db.Open()
	defer tbl.db.Close()
	var result sql.Result
	var err error

	query := "insert or replace into tags(id, tag) values((select id from tags where tag = ?), ?)"
	result, err = tbl.db.con.Exec(query, tbl.Tag, tbl.Tag)

	if err != nil {
		fmt.Printf("TblTags.Save: %s\n", err.Error())
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		fmt.Printf("TblTags.Save: %s\n", err.Error())
		return err
	}
	tbl.Id = int(id)
	return nil
}

func (tbl *TblTags) Search(q string) ([]TblTags, error) {
	tbl.db = &Db{}
	tbl.db.Open()
	defer tbl.db.Close()
	var result []TblTags = []TblTags{}
	query := "select id, tag from tags where tag like '%" + q + "%'"

	rows, err := tbl.db.con.Query(query)
	if err != nil {
		fmt.Printf("tblTag.Search: %s\n", err.Error())
		return nil, err
	}

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
