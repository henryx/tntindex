/*
   Copyright (C) 2019 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package index

import "time"

type data struct {
	Date     time.Time `csv:"DATA"`
	Hash     string    `csv:"HASH"`
	Topic    int       `csv:"TOPIC"`
	Post     int       `csv:"POST"`
	Author   string    `csv:"AUTORE"`
	Title    string    `csv:"TITOLO"`
	Desc     string    `csv:"DESCRIZIONE"`
	Size     int       `csv:"DIMENSIONE"`
	Category int       `csv:"CATEGORIA"`
}
