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
	"tntindex/search"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var db database.DB

	// Index command
	cmdIndex := kingpin.Command("index", "Index a file")
	indexFile := cmdIndex.Arg("filename", "Filename to index").Required().String()

	// Search command
	cmdSearch := kingpin.Command("search", "Search a torrent")
	term := cmdSearch.Arg("term", "Term to search").Required().String()

	kingpin.HelpFlag.Short('h')

	db.Open("tntindex.db")
	defer db.Close()

	switch kingpin.Parse() {
	case cmdIndex.FullCommand():
		if err := index.Index(&db, indexFile); err != nil {
			log.Fatalln("Error when indexing data:", err)
		}
	case cmdSearch.FullCommand():
	}
}
