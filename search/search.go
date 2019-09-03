/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package search

import "tntindex/database"

func Search(db *database.DB, key *string) error {
	// TODO: manage extracted data
	_, err := db.SearchPost(*key)
	if err != nil {
		return err
	}

	return nil
}
