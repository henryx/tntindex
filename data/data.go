/*
   Copyright (C) 2019-2022 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package data

import "time"

type Time struct {
	time.Time
}

const format = "2006-01-02T15:04:05"

func (t *Time) UnmarshalCSV(data []byte) error {
	tt, err := time.Parse(format, string(data))
	if err != nil {
		return err
	}
	*t = Time{Time: tt}
	return nil
}

type Data struct {
	Date     Time   `csv:"DATA"`
	Hash     string `csv:"HASH"`
	Topic    int    `csv:"TOPIC"`
	Post     int    `csv:"POST"`
	Author   string `csv:"AUTORE"`
	Title    string `csv:"TITOLO"`
	Desc     string `csv:"DESCRIZIONE"`
	Size     int    `csv:"DIMENSIONE"`
	Category int    `csv:"CATEGORIA"`
}
