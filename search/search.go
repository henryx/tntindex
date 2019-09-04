/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package search

import (
	"fmt"
	"tntindex/database"
)

func Search(db *database.DB, key *string) error {
	type res struct {
		Title string
		Hash  string
	}

	posts, err := db.SearchPost(*key)
	if err != nil {
		return err
	}

	for _, post := range posts {
		hash, err := db.SearchHash(post)
		if err != nil {
			return err
		}

		fmt.Printf("magnet:?xt=urn:btih:%s - %s\n", hash, post.Title)
	}

	return nil
}
