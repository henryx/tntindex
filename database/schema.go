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
	var tables = []string{
		"CREATE TABLE categories(id, name)",
		"CREATE VIRTUAL TABLE posts USING fts5(posted, topic, post, author, title, description)",
		"CREATE TABLE hashes (topic, post, hash, size, category)",
	}

	tx, _ := d.conn.Begin()
	for _, table := range tables {
		_, err := tx.Exec(table)
		if err != nil {
			log.Fatalln("Error creating tables schema:", err)
		}
	}
	tx.Commit()
}
