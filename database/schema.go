/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package database

import "log"

func (d *DB) checkSchema() bool {
	var counted int

	query := "SELECT count(*) FROM sqlite_master"

	if err := d.conn.QueryRow(query).Scan(&counted); err != nil {
		log.Fatalln("Database not opened: " + err.Error())
	}

	if counted > 0 {
		return true
	}

	return false
}

func (d *DB) createSchema() {
	// TODO create schema
}
