/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package index

import "time"

type data struct {
	Date     time.Time `csv:"data"`
	Hash     string    `csv:"hash"`
	Topic    int       `csv:"topic"`
	Post     int       `csv:"post"`
	Author   string    `csv:"autore"`
	Title    string    `csv:"titolo"`
	Desc     string    `csv:"descrizione"`
	Size     int       `csv:"dimensione"`
	Category int       `csv:"categoria"`
}
