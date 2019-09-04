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
		"CREATE TABLE categories(category, name, PRIMARY KEY(category))",
		"CREATE VIRTUAL TABLE posts USING fts5(posted, topic, post, author, title, description)",
		"CREATE TABLE hashes (topic, post, hash, size, category, PRIMARY KEY(topic, post))",
	}

	var data = []string{
		"INSERT INTO categories VALUES(1, 'Film TV e programmi')",
		"INSERT INTO categories VALUES(2, 'Musica')",
		"INSERT INTO categories VALUES(3, 'E Books')",
		"INSERT INTO categories VALUES(4, 'Film')",
		"INSERT INTO categories VALUES(6, 'Linux')",
		"INSERT INTO categories VALUES(7, 'Anime')",
		"INSERT INTO categories VALUES(8, 'Cartoni')",
		"INSERT INTO categories VALUES(9, 'Macintosh')",
		"INSERT INTO categories VALUES(10, 'Windows Software')",
		"INSERT INTO categories VALUES(11, 'Pc Game')",
		"INSERT INTO categories VALUES(12, 'Playstation')",
		"INSERT INTO categories VALUES(13, 'Students Releases')",
		"INSERT INTO categories VALUES(14, 'Documentari')",
		"INSERT INTO categories VALUES(21, 'Video Musicali')",
		"INSERT INTO categories VALUES(22, 'Sport')",
		"INSERT INTO categories VALUES(23, 'Teatro')",
		"INSERT INTO categories VALUES(24, 'Wrestling')",
		"INSERT INTO categories VALUES(25, 'Varie')",
		"INSERT INTO categories VALUES(26, 'Xbox')",
		"INSERT INTO categories VALUES(27, 'Immagini sfondi')",
		"INSERT INTO categories VALUES(28, 'Altri Giochi')",
		"INSERT INTO categories VALUES(29, 'Serie TV')",
		"INSERT INTO categories VALUES(30, 'Fumetteria')",
		"INSERT INTO categories VALUES(31, 'Trash')",
		"INSERT INTO categories VALUES(32, 'Nintendo')",
		"INSERT INTO categories VALUES(34, 'A Book')",
		"INSERT INTO categories VALUES(35, 'Podcast')",
		"INSERT INTO categories VALUES(36, 'Edicola')",
		"INSERT INTO categories VALUES(37, 'Mobile')",
	}

	tx, _ := d.conn.Begin()
	for _, table := range tables {
		_, err := tx.Exec(table)
		if err != nil {
			log.Fatalln("Error creating tables schema:", err)
		}
	}
	tx.Commit()

	tx, _ = d.conn.Begin()
	for _, val := range data {
		_, err := tx.Exec(val)
		if err != nil {
			log.Fatalln("Error inserting values:", err)
		}
	}
	tx.Commit()
}
