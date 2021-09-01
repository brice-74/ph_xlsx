package main

import (
	"regexp"
	"strconv"
	"log"
	"strings"

	"github.com/xuri/excelize/v2"
)

type Headers struct {
	codification 			[]string
	rubrique_paie			[]string
	global						[]string
}

func scanXlsxOnly(path string, sheetName string) (*[][]string, string) {
	file 	:= getOnlyFile(scanDir(path))
	fxlsx := openXlsxFile(path + file.Name())
	rows 	:= getRowsFromExcelize(fxlsx, sheetName)
	
	return rows, file.Name()
}

func extractHeaders(rows *[][]string) (*Headers) {
	var headers *Headers
	headers = new(Headers)
	// use regex to separate "codifications" & "rubriques paie"
	var codificationRegex = regexp.MustCompile(`(Code analytique [0-9]*)\w+|(^%$)`)
	// extract headers and create unique keys for headers name %
	index := 1
	for _, title := range (*rows)[0] {
		if codificationRegex.MatchString(title) {
			if title == "%" {
				title = "%" + strconv.Itoa(index)
				index++
			}
			headers.codification = append(headers.codification, title)
		}else{
			headers.rubrique_paie = append(headers.rubrique_paie, title)
		}
	}
	// list all headers
	headers.global = append(headers.codification, headers.rubrique_paie...)

	return headers
}

func mapRows(rows *[][]string, headers *Headers) (*map[int]map[string]string) {
	var rowsMap = make(map[int]map[string]string)
	type colsMap = map[string]string
	// map rows using headers 
	for i, row := range *rows {
		rowsMap[i] = make(colsMap)
		for i2, col := range row {
			// not mapping 0 values 
			if col != "" && col != "0" {
				rowsMap[i][headers.global[i2]] = col
			}
		}
	}
	// delete headers mapping headers
	delete(rowsMap, 0); 

	return &rowsMap
}

func mapSum(mapRows *map[int]map[string]string, headers *Headers) (*map[string]map[string]float64) {
	var sumMapping = make(map[string]map[string]float64)
	var sumMappingPointer = &sumMapping
	extractNumeric := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)

	for _, row := range *mapRows {
		for _, hc := range headers.codification {	
			if (row[hc] != "" && row[hc] != "0") {
				if hc[0:1] != "%" {
					// get % associated to the code analytique 
					strPourcentage := strings.Replace(row["%" + extractNumeric.FindString(hc)], ",", ".", -1)
					pourcentage, err := strconv.ParseFloat(strPourcentage, 64)
					if err != nil {
						log.Fatal("Parse float error : "+ row[hc] +" | "+ row["%" + extractNumeric.FindString(hc)])
					}
					addKeyCodsNoDuplicata(sumMappingPointer, row[hc])
					for _, hr := range headers.rubrique_paie {
						if (row[hr] != "" && row[hr] != "0"){
							float, err := strconv.ParseFloat(strings.Replace(row[hr], ",", ".", -1), 64)
							if err != nil {
								log.Fatal("Parse float error : "+ row[hc] +" | "+hr+" | "+row[hr])
							}
							float = float * pourcentage / 100
							addKeyRubSum(sumMappingPointer, row[hc], hr, float)
						}
					}
				}
			}
		}
	}

	return &sumMapping
}  

func generateXlsx(sumMapping *map[string]map[string]float64, fileName string) () {
	f := excelize.NewFile()
	sheet := f.NewSheet("Sheet1")

	i := 1
	for codification, row := range *sumMapping {
		for rubrique_paie, montant_proratise := range row {
			istr := strconv.Itoa(i)
			f.SetCellValue("Sheet1", "A" + istr, codification)
			f.SetCellValue("Sheet1", "B" + istr, rubrique_paie)
			f.SetCellValue("Sheet1", "C" + istr, montant_proratise)
			i++
		}
	}

	f.SetActiveSheet(sheet)
	if err := f.SaveAs("generate/_"+ fileName); err != nil {
		log.Fatal(err)
	}

	log.Println("Create file : _" + fileName + " generated successfuly")
}