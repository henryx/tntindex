/*
   Copyright (C) 2019-2022 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package index

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"tntindex/data"
	"tntindex/database"

	"github.com/jszwec/csvutil"
)

// Index index a file in database
func Index(db *database.DB, filename *string) error {
	fd, err := os.Open(*filename)
	if err != nil {
		return err
	}
	defer fd.Close()

	reader := csv.NewReader(bufio.NewReader(fd))
	dec, err := csvutil.NewDecoder(reader)
	for {
		d := data.Data{}

		if err := dec.Decode(&d); err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		if err := db.IndexData(&d); err != nil {
			return err
		}
	}

	return nil
}
