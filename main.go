/*
   Copyright (C) 2019-2022 Enrico Bianchi (enrico.bianchi@gmail.com)
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
	"tntindex/tracker"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var db database.DB
	var fileName *string

	// Index command
	cmdIndex := kingpin.Command("index", "Index a file")
	fileName = cmdIndex.Arg("filename", "Filename to index").Required().String()

	// Search command
	cmdSearch := kingpin.Command("search", "Search a torrent")
	term := cmdSearch.Arg("term", "Term to search").Required().String()

	// Tracker command
	cmdTracker := kingpin.Command("tracker", "Add trackers to magnet link from file")
	fileName = cmdTracker.Flag("file", "A file that contains a list of trackers").Short('f').Required().String()
	magnet := cmdTracker.Arg("magnet", "A torrent magnet link").Required().String()

	kingpin.HelpFlag.Short('h')

	parsed := kingpin.Parse()

	db.Open("tntindex.db")
	defer db.Close()

	switch parsed {
	case cmdIndex.FullCommand():
		if err := index.Index(&db, fileName); err != nil {
			log.Fatalln("Error when indexing data:", err)
		}
	case cmdSearch.FullCommand():
		if err := search.Search(&db, term); err != nil {
			log.Fatalln("Error when searching data:", err)
		}
	case cmdTracker.FullCommand():
		if err := tracker.Generate(fileName, magnet); err != nil {
			log.Fatalln("Error when opening tracker list file:", err)
		}
	}
}
