/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package index

import "fmt"

// Index index a file in database
func Index(filename *string) error {
	fmt.Println("you have passed", *filename, "file")

	return nil
}
