/*
   Copyright (C) 2019-2022 Enrico Bianchi (enrico.bianchi@gmail.com)
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

// DB struct define a structure to manage the database
type DB struct {
	conn *sqlx.DB
}

// Open open the connection to the database
func (d *DB) Open(name string) error {
	var err error

	d.conn, err = sqlx.Connect("sqlite3", fmt.Sprintf("file:%s?_journal=wal&cache=shared", name))
	if err != nil {
		return err
	}
	d.conn.SetMaxOpenConns(1)

	if !d.checkSchema() {
		d.createSchema()
	}

	return nil
}

// Close close the connection to the database
func (d *DB) Close() {
	d.conn.Close()
}
