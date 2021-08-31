package main

import (
	"fmt"
	"flag"
	"time"
)

func main() {
	defer end(time.Now())

	fmt.Println("------------------ Program is running ------------------")
	sheet := flag.String("sheet", "Feuille 1", "Sheetname to open")
	flag.Parse()

	rows, filename := scanXlxsOnly("./scan/", *sheet)

	// test with less lines from rows 
	// testRows := [][]string{(*rows)[0], (*rows)[22]}
	// testRows = [][]string{{"titre 1", "titre 2", "titre 3"}, {"val 1", "val 2", "val 3"}, {"VAL1", "VAL2", "VAL3"}} 
	headers := extractHeaders(rows)
	mapRows := mapRows(rows, headers)
	sumMapping := mapSum(mapRows, headers)
	generateXlxx(sumMapping, filename)
	
}

func end(start time.Time) {
	fmt.Println(time.Since(start))
}