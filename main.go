/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package main

import (
	"log"
	"tntindex/database"
	"tntindex/index"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var db database.DB

	// Index command
	cmdIndex := kingpin.Command("index", "Index a file")
	indexFile := cmdIndex.Arg("filename", "Filename to index").Required().String()

	db.Open("tntindex.db")

	switch kingpin.Parse() {
	case cmdIndex.FullCommand():
		if err := index.Index(indexFile); err != nil {
			log.Fatalln("Error when indexing data:", err)
		}
	}
}
