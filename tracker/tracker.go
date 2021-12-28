/*
   Copyright (C) 2019-2022 Enrico Bianchi (enrico.bianchi@gmail.com)
   Project       tntindex
   Description   An TNTVillage dump indexer and search tool
   License       GPL version 2 (see LICENSE for details)
*/

package tracker

import (
	"bufio"
	"fmt"
	"os"
)

func Generate(fileName, magnet *string) error {
	f, err := os.Open(*fileName)
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	line := fmt.Sprintf("%s", magnet)
	for scanner.Scan() {
		line = fmt.Sprintf("%s&%s", line, scanner.Text())
	}
	fmt.Println("Magnet with trackers:")
	fmt.Println(line)

	return nil
}
