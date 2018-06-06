package main

import (
	"fmt"
	"time"

	"github.com/Decarium/go-coinmarketcap/coinmarketcap"
	"github.com/davecgh/go-spew/spew"
	spreadsheet "gopkg.in/Iwark/spreadsheet.v2"
)

const (
	BILLION = 1000000000
)

var (
	marketCapSheet *spreadsheet.Sheet
)

type MarketCapData struct {
	TotalMarketCap                       string
	MarketCapGrowth                      string
	MarketCapGrowthPercentage7D          string
	MarketCapGrowthPercentage1M          string
	MarketCapGrowthPercentageYTD         string
	TotalWeeklyVolume                    string
	TotalWeeklyVolumeGrowth              string
	TotalWeeklyVolumeGrowthPercentage7D  string
	TotalWeeklyVolumeGrowthPercentage1M  string
	TotalWeeklyVolumeGrowthPercentageYTD string
	AltcoinMarketCap                     string
	AltcoinMarketCapGrowth               string
	AltcoinMarketCapGrowthPercentage7D   string
	AltcoinMarketCapGrowthPercentage1M   string
	AltcoinMarketCapGrowthPercentageYTD  string
}

//Create functions for each section

//Quick thought here which is that we are taking the average marketcap of the day, we should move this
//To be the market cap at the end of the day UTC time. This way we don't skew the numbers by average.
//We can fix this when we move this to a more standard thing
func CreateMarketCap() {

	var err error
	marketCapSheet, err = mainSheet.SheetByID(1220239967)

	if err != nil {
		fmt.Println(err)
	}

	// t := template.New("marketcap.tmpl")
	// t, err := t.ParseFiles("./templates/marketcap.tmpl") // Parse template file.

	// if err != nil {
	// 	log.Fatal(err)
	// }

	data := GetMarketCapData()

	spew.Dump(data)

	// f, err := os.Create("./pdf/sections/marketcap/marketcap.pug")
	// if err != nil {
	// 	log.Println("create file: ", err)
	// 	return
	// }

	// err = t.Execute(f, data)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// f.Close()

	BuildMarketCapSpreadSheet(data)
}

func BuildMarketCapSpreadSheet(data MarketCapData) {

	newRows := len(marketCapSheet.Rows) + 6

	newColumns := len(marketCapSheet.Columns)

	err := service.ExpandSheet(marketCapSheet, uint(newRows), uint(newColumns)) // Expand the sheet to 20 rows and 10 columns

	row := len(marketCapSheet.Rows) + 1

	spew.Dump(row)

	column := 0

	//Get Today's Date
	now := time.Now()

	date := fmt.Sprintf("%d/%d/%d", now.Day(), int(now.Month()), now.Year())

	sheet.Update(row, column, date)

	rowHeaders := []string{"Total", "24H", "7D", "1M", "YTD"}

	for _, header := range rowHeaders {

		column++

		sheet.Update(row, column, header)

	}

	//Put these into an array of structs

	//bump a row
	row++
	//reset column
	column = 0

	marketCapSheet.Update(row, column, "Market Cap")

	column++

	//Row 1 = total market cap
	marketCapSheet.Update(row, column, data.TotalMarketCap)

	column++

	//24H growth
	marketCapSheet.Update(row, column, "N/A")

	column++

	marketCapSheet.Update(row, column, data.MarketCapGrowthPercentage7D)

	column++

	marketCapSheet.Update(row, column, data.MarketCapGrowthPercentage1M)

	column++

	marketCapSheet.Update(row, column, data.MarketCapGrowthPercentageYTD)

	//bump a row
	row++
	//reset column
	column = 0

	marketCapSheet.Update(row, column, "Total Volume")

	column++

	//Row 1 = total market cap
	marketCapSheet.Update(row, column, data.TotalWeeklyVolume)

	column++

	//24H growth
	marketCapSheet.Update(row, column, "N/A")

	column++

	marketCapSheet.Update(row, column, data.TotalWeeklyVolumeGrowthPercentage7D)

	column++

	marketCapSheet.Update(row, column, data.TotalWeeklyVolumeGrowthPercentage1M)

	column++

	marketCapSheet.Update(row, column, data.TotalWeeklyVolumeGrowthPercentageYTD)

	//bump a row
	row++
	//reset column
	column = 0

	marketCapSheet.Update(row, column, "Total Market Cap W/out Bitcoin")

	column++

	//Row 1 = total market cap
	marketCapSheet.Update(row, column, data.AltcoinMarketCap)

	column++

	//24H growth
	marketCapSheet.Update(row, column, "N/A")

	column++

	marketCapSheet.Update(row, column, data.AltcoinMarketCapGrowthPercentage7D)

	column++

	marketCapSheet.Update(row, column, data.AltcoinMarketCapGrowthPercentage1M)

	column++

	marketCapSheet.Update(row, column, data.AltcoinMarketCapGrowthPercentageYTD)

	// Make sure call Synchronize to reflect the changes.
	err = marketCapSheet.Synchronize()

	if err != nil {
		fmt.Println(err)
	}

}

func GetMarketCapData() MarketCapData {

	//Total Market Cap
	totalMarketCap := GetTotalMarketCap()

	marketCap7D := GetTotalMarketCapLastWeek()
	marketCap1M := GetTotalMarketCapLastMonth()
	marketCapYTD := GetTotalMarketCapLastYear()

	// Market Cap Growth

	growth7D := totalMarketCap - marketCap7D

	growth1M := totalMarketCap - marketCap1M

	growthYTD := totalMarketCap - marketCapYTD

	growthPercentage7D := (growth7D / marketCap7D) * 100

	growthPercentage1M := (growth1M / marketCap1M) * 100

	growthPercentageYTD := (growthYTD / marketCapYTD) * 100

	//Total Market Cap
	altcoinMarketCap := GetAltcoinMarketCap()

	altcoinMarketCap7D := GetAltcoinMarketCapLastWeek()
	altcoinMarketCap1M := GetAltcoinMarketCapLastMonth()
	altcoinMarketCapYTD := GetAltcoinMarketCapLastYear()

	// Market Cap Growth

	altGrowth7D := altcoinMarketCap - altcoinMarketCap7D

	altGrowth1M := altcoinMarketCap - altcoinMarketCap1M

	altGrowthYTD := altcoinMarketCap - altcoinMarketCapYTD

	altGrowthPercentage7D := (altGrowth7D / altcoinMarketCap7D) * 100

	altGrowthPercentage1M := (altGrowth1M / altcoinMarketCap1M) * 100

	altGrowthPercentageYTD := (altGrowthYTD / altcoinMarketCapYTD) * 100

	// Total Weekly Volume

	totalWeeklyVolume := GetTotalWeeklyVolume()

	weeklyVolume7D := GetTotalWeeklyVolumeLastWeek()

	weeklyVolume1M := GetTotalWeeklyVolumeLastMonth()

	weeklyVolumeYTD := GetTotalWeeklyVolumeLastYear()

	// Weekly Volume Growth

	volumeGrowth7D := totalWeeklyVolume - weeklyVolume7D

	volumeGrowth1M := totalWeeklyVolume - weeklyVolume1M

	volumeGrowthYTD := totalWeeklyVolume - weeklyVolumeYTD

	volumeGrowthPercentage7D := (volumeGrowth7D / weeklyVolume7D) * 100

	volumeGrowthPercentage1M := (volumeGrowth1M / weeklyVolume1M) * 100

	volumeGrowthPercentageYTD := (volumeGrowthYTD / weeklyVolumeYTD) * 100

	//Format everything
	totalMarketCapFormatted := fmt.Sprintf("%.2f", (totalMarketCap / BILLION))

	growthFormatted := fmt.Sprintf("%.2f", (growth7D / BILLION))

	growthPercentageFormatted7D := fmt.Sprintf("%.2f", growthPercentage7D)
	growthPercentageFormatted1M := fmt.Sprintf("%.2f", growthPercentage1M)
	growthPercentageFormattedYTD := fmt.Sprintf("%.2f", growthPercentageYTD)

	//Format everything
	altcoinMarketCapFormatted := fmt.Sprintf("%.2f", (altcoinMarketCap / BILLION))

	altcoinGrowthFormatted := fmt.Sprintf("%.2f", (altGrowth7D / BILLION))

	altcoinGrowthPercentageFormatted7D := fmt.Sprintf("%.2f", altGrowthPercentage7D)
	altcoinGrowthPercentageFormatted1M := fmt.Sprintf("%.2f", altGrowthPercentage1M)
	altcoinGrowthPercentageFormattedYTD := fmt.Sprintf("%.2f", altGrowthPercentageYTD)

	totalWeeklyVolumeFormatted := fmt.Sprintf("%.2f", (totalWeeklyVolume / BILLION))

	volumeGrowthFormatted := fmt.Sprintf("%.2f", (volumeGrowth7D / BILLION))

	volumeGrowthPercentageFormatted7D := fmt.Sprintf("%.2f", volumeGrowthPercentage7D)

	volumeGrowthPercentageFormatted1M := fmt.Sprintf("%.2f", volumeGrowthPercentage1M)

	volumeGrowthPercentageFormattedYTD := fmt.Sprintf("%.2f", volumeGrowthPercentageYTD)

	//If growth is positive, we add the + sign
	if growth7D > 0 {
		growthFormatted = "+" + growthFormatted
	}

	mcd := MarketCapData{
		TotalMarketCap:                       totalMarketCapFormatted,
		MarketCapGrowth:                      growthFormatted,
		MarketCapGrowthPercentage7D:          growthPercentageFormatted7D,
		MarketCapGrowthPercentage1M:          growthPercentageFormatted1M,
		MarketCapGrowthPercentageYTD:         growthPercentageFormattedYTD,
		TotalWeeklyVolume:                    totalWeeklyVolumeFormatted,
		TotalWeeklyVolumeGrowth:              volumeGrowthFormatted,
		TotalWeeklyVolumeGrowthPercentage7D:  volumeGrowthPercentageFormatted7D,
		TotalWeeklyVolumeGrowthPercentage1M:  volumeGrowthPercentageFormatted1M,
		TotalWeeklyVolumeGrowthPercentageYTD: volumeGrowthPercentageFormattedYTD,
		AltcoinMarketCap:                     altcoinMarketCapFormatted,
		AltcoinMarketCapGrowth:               altcoinGrowthFormatted,
		AltcoinMarketCapGrowthPercentage7D:   altcoinGrowthPercentageFormatted7D,
		AltcoinMarketCapGrowthPercentage1M:   altcoinGrowthPercentageFormatted1M,
		AltcoinMarketCapGrowthPercentageYTD:  altcoinGrowthPercentageFormattedYTD,
	}

	//We want to format this data so that it looks clean on the excel

	return mcd

}

//Done
func GetTotalMarketCap() float64 {

	//Assuming we are running this on a Tuesday. If we can make this work on
	//Any day of the week that would be great.
	t := time.Now().AddDate(0, 0, -1)

	//Start date is going to be midnight 7 days ago
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	spew.Dump("Start time: ")
	spew.Dump(start)

	//End will be 1 day forward so we get 24hrs of ticks
	end := start.AddDate(0, 0, 1)

	spew.Dump("End Time: ")
	spew.Dump(end)

	global, err := coinmarketcap.GetGlobalHistoricalTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	return global.MarketCapByAvailableSupply[0].Amount
}

//Done
func GetTotalMarketCapLastWeek() float64 {

	//Assuming we are running this on a Tuesday. If we can make this work on
	//Any day of the week that would be great.
	t := time.Now().AddDate(0, 0, -8)

	//Start date is going to be midnight 7 days ago
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//End will be 1 day forward so we get 24hrs of ticks
	end := start.AddDate(0, 0, 1)

	global, err := coinmarketcap.GetGlobalHistoricalTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	return global.MarketCapByAvailableSupply[0].Amount
}

func GetTotalMarketCapLastMonth() float64 {

	t := time.Now().AddDate(0, 0, -31)

	//Start date is going to be midnight 7 days ago
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//End will be 1 day forward so we get 24hrs of ticks
	end := start.AddDate(0, 0, 1)

	global, err := coinmarketcap.GetGlobalHistoricalTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	return global.MarketCapByAvailableSupply[0].Amount
}

func GetTotalMarketCapLastYear() float64 {

	yearDay := time.Now().YearDay()

	t := time.Now().AddDate(0, 0, -(yearDay))

	//Start date is going to be midnight 7 days ago
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//End will be 1 day forward so we get 24hrs of ticks
	end := start.AddDate(0, 0, 1)

	global, err := coinmarketcap.GetGlobalHistoricalTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	return global.MarketCapByAvailableSupply[0].Amount
}

func GetTotalWeeklyVolume() float64 {

	t := time.Now().AddDate(0, 0, -8)

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	spew.Dump("start time:")
	spew.Dump(start)

	//So we do last week + 7 days so if this report is released on a Tuesday, we get Monday to Monday
	end := start.AddDate(0, 0, 8)

	spew.Dump("end time:")
	spew.Dump(end)

	global, err := coinmarketcap.GetGlobalHistoricalTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	var total float64

	//Global is now 7 days worth of volume so we want to iterate through it
	for _, day := range global.VolumeUsd {
		total += day.Amount
	}

	return total
}

func GetTotalWeeklyVolumeLastWeek() float64 {

	t := time.Now().AddDate(0, 0, -15)

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//So we do last week + 7 days so if this report is released on a Tuesday, we get Monday to Monday
	end := start.AddDate(0, 0, 8)

	global, err := coinmarketcap.GetGlobalHistoricalTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	var total float64

	//Global is now 7 days worth of volume so we want to iterate through it
	for _, day := range global.VolumeUsd {
		total += day.Amount
	}

	return total
}

func GetTotalWeeklyVolumeLastMonth() float64 {

	t := time.Now().AddDate(0, 0, -37)

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//So we do last week + 7 days so if this report is released on a Tuesday, we get Monday to Monday
	end := start.AddDate(0, 0, 8)

	global, err := coinmarketcap.GetGlobalHistoricalTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	var total float64

	//Global is now 7 days worth of volume so we want to iterate through it
	for _, day := range global.VolumeUsd {
		total += day.Amount
	}

	return total
}

func GetTotalWeeklyVolumeLastYear() float64 {

	yearDate := time.Now().YearDay()

	t := time.Now().AddDate(0, 0, -(yearDate + 7))

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//So we do last week + 7 days so if this report is released on a Tuesday, we get Monday to Monday
	end := start.AddDate(0, 0, 8)

	global, err := coinmarketcap.GetGlobalHistoricalTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	var total float64

	//Global is now 7 days worth of volume so we want to iterate through it
	for _, day := range global.VolumeUsd {
		total += day.Amount
	}

	return total
}

func GetAltcoinMarketCap() float64 {

	//Assuming we are running this on a Tuesday. If we can make this work on
	//Any day of the week that would be great.
	t := time.Now().AddDate(0, 0, -1)

	//Start date is going to be midnight 7 days ago
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	spew.Dump("Start time: ")
	spew.Dump(start)

	//End will be 1 day forward so we get 24hrs of ticks
	end := start.AddDate(0, 0, 1)

	spew.Dump("End Time: ")
	spew.Dump(end)

	global, err := coinmarketcap.GetGlobalHistoricalAltcoinTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	return global.MarketCapByAvailableSupply[0].Amount
}

//Done
func GetAltcoinMarketCapLastWeek() float64 {

	//Assuming we are running this on a Tuesday. If we can make this work on
	//Any day of the week that would be great.
	t := time.Now().AddDate(0, 0, -8)

	//Start date is going to be midnight 7 days ago
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//End will be 1 day forward so we get 24hrs of ticks
	end := start.AddDate(0, 0, 1)

	global, err := coinmarketcap.GetGlobalHistoricalAltcoinTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	return global.MarketCapByAvailableSupply[0].Amount
}

func GetAltcoinMarketCapLastMonth() float64 {

	t := time.Now().AddDate(0, 0, -31)

	//Start date is going to be midnight 7 days ago
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//End will be 1 day forward so we get 24hrs of ticks
	end := start.AddDate(0, 0, 1)

	global, err := coinmarketcap.GetGlobalHistoricalAltcoinTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	return global.MarketCapByAvailableSupply[0].Amount
}

func GetAltcoinMarketCapLastYear() float64 {

	yearDay := time.Now().YearDay()

	t := time.Now().AddDate(0, 0, -(yearDay))

	//Start date is going to be midnight 7 days ago
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//End will be 1 day forward so we get 24hrs of ticks
	end := start.AddDate(0, 0, 1)

	global, err := coinmarketcap.GetGlobalHistoricalAltcoinTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	return global.MarketCapByAvailableSupply[0].Amount
}
