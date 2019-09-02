/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sqlx.DB
}

func (d *DB) Open(name string) error {
	var err error

	d.conn, err = sqlx.Connect("sqlite3", fmt.Sprintf("file:%s?_journal=wal", name))
	if err != nil {
		return err
	}
	return nil
}
