/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package main

import (
	"tntindex/index"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	// Index command
	cmdIndex := kingpin.Command("index", "Index a file")
	indexFile := cmdIndex.Arg("filename", "Filename to index").Required().String()

	switch kingpin.Parse() {
	case cmdIndex.FullCommand():
		index.Index(indexFile)
	}
}
