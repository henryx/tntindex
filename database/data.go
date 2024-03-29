/*
   Copyright (C) 2019-2022 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package database

import "tntindex/data"

type Post struct {
	Topic int
	Post  int
	Title string
	Size  int
}

func (d *DB) IndexData(val *data.Data) error {
	post := "INSERT INTO posts VALUES(?, ?, ?, ?, ?, ?)"
	hash := "INSERT INTO hashes VALUES(?, ?,?, ?, ?)"

	tx := d.conn.MustBegin()
	if _, err := tx.Exec(post, val.Date.String(), val.Topic, val.Post, val.Author, val.Title, val.Desc); err != nil {
		tx.Rollback()
		return err
	}

	if _, err := tx.Exec(hash, val.Topic, val.Post, val.Hash, val.Size, val.Category); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()

	return nil
}

func (d *DB) SearchPost(val string) ([]Post, error) {
	posts := []Post{}
	query := "SELECT topic, post, title FROM posts WHERE title MATCH ? OR description MATCH ?"

	err := d.conn.Select(&posts, query, val, val)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (d *DB) SearchHash(val Post) (string, int, error) {
	var res string
	var size int

	query := "SELECT hash, size FROM hashes WHERE topic = ? AND post = ?"

	err := d.conn.QueryRow(query, val.Topic, val.Post).Scan(&res, &size)
	if err != nil {
		return "", -1, err
	}

	return res, size, nil
}
