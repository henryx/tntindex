/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package database

import "tntindex/data"

func (d *DB) IndexData(val *data.Data) error {
	post := "INSERT INTO posts VALUES(?, ?, ?, ?, ?, ?)"
	hash := "INSERT INTO hashes VALUES(?, ?,?, ?, ?)"

   tx := d.conn.MustBegin()
   if _, err := tx.Exec(post, val.Date.String(), val.Topic, val.Post, val.Author, val.Title, val.Desc);     err != nil {
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
