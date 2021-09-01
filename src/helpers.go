package main

import (
	"io/ioutil"
	"log"
	"io/fs"

	"github.com/xuri/excelize/v2"
)

func scanDir(path string) ([]fs.FileInfo) {
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func getOnlyFile(files []fs.FileInfo) (fs.FileInfo) {
	if (len(files) < 1 || len(files) > 1) {
		log.Fatal("One file is require")
	}
	log.Println("Open file : " + files[0].Name())

	return files[0]
} 

func openXlsxFile(pathfile string) (*excelize.File) {
	f, err := excelize.OpenFile(pathfile)
	if err != nil {
		log.Fatal(err)
	}

	return f
}

func getRowsFromExcelize(f *excelize.File, sheetName string) (*[][]string) {
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	return &rows
}

func addKeyCodsNoDuplicata(arr *map[string]map[string]float64, key string) () {
	if _, found := (*arr)[key]; !found {
		(*arr)[key] = make(map[string]float64)
	}
}

func addKeyRubSum(arr *map[string]map[string]float64, key1 string, key2 string, val float64) () {
	val1 := (*arr)[key1][key2]
	(*arr)[key1][key2] = val1 + val
} 