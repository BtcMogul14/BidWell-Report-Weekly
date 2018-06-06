package main

import (
	"fmt"

	spreadsheet "gopkg.in/Iwark/spreadsheet.v2"
)

//TODO make sure that everything is in UTC time specifically, the searches on coinmarketcap

var (
	service   *spreadsheet.Service
	mainSheet spreadsheet.Spreadsheet
)

//Set up the service
func initSpreadsheet() {
	var err error
	service, err = spreadsheet.NewService()

	if err != nil {
		fmt.Println(err)
	}

	//Return Empty so that spreadsheet gets the empty cells

	spreadsheetID := "1JSXgIRI4_aw52ChegBmVmiGZjuCil5RL7U4z-h09N_M"
	mainSheet, err = service.FetchSpreadsheet(spreadsheetID)
}

func main() {
	initSpreadsheet()
	// CreateMarketCap()
	// CreateMarketCapChart()
	// CreateBlueChips()
	CreateTop100()
	// CreateSectors()
}
