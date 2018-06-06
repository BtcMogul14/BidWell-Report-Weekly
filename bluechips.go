package main

import (
	"fmt"
	"time"

	"github.com/Decarium/go-coinmarketcap/coinmarketcap"
	spreadsheet "gopkg.in/Iwark/spreadsheet.v2"
)

var (
	blueChipSheet *spreadsheet.Sheet
)

type BlueChip struct {
	Name            string
	Slug            string
	Price           string
	PriceGrowth7D   string
	PriceGrowth24H  string
	Volume7D        string
	Volume24H       string
	VolumeGrowth7D  string
	VolumeGrowth24H string
	MarketCap       string
}

type BlueChipsData struct {
	chips []BlueChip
}

func CreateBlueChips() {

	var err error
	blueChipSheet, err = mainSheet.SheetByID(565576139)

	if err != nil {
		fmt.Println(err)
	}

	// t := template.New("bluechips.tmpl")
	// t, err := t.ParseFiles("./templates/bluechips.tmpl") // Parse template file.

	// if err != nil {
	// 	log.Fatal(err)
	// }

	data := GetBlueChipsData()

	// f, err := os.Create("./pdf/sections/bluechips/bluechips.pug")
	// if err != nil {
	// 	log.Println("create file: ", err)
	// 	return
	// }

	// err = t.Execute(f, data)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// f.Close()

	BuildBlueSpreadSheet(data)
}

func BuildBlueSpreadSheet(data BlueChipsData) {

	newRows := len(blueChipSheet.Rows) + 7

	newColumns := len(blueChipSheet.Columns)

	err := service.ExpandSheet(blueChipSheet, uint(newRows), uint(newColumns)) // Expand the sheet to 20 rows and 10 columns

	if err != nil {
		fmt.Println(err)
	}

	row := len(blueChipSheet.Rows) + 1

	column := 0

	//Get Today's Date
	now := time.Now()

	date := fmt.Sprintf("%d/%d/%d", now.Day(), int(now.Month()), now.Year())

	blueChipSheet.Update(row, column, date)

	//Make this into an array of strings, and just iterate through.
	rowHeaders := []string{"Price", "24H Price Change", "24H Volume", "24H Volume Change", "7D Price Change", "7D Volume", "7D Volume Change"}

	for _, header := range rowHeaders {
		column++

		blueChipSheet.Update(row, column, header)

	}

	//Go down a row
	row++

	for _, chip := range data.chips {

		column = 0

		blueChipSheet.Update(row, column, chip.Slug)

		column++

		//Price
		blueChipSheet.Update(row, column, chip.Price)

		column++

		blueChipSheet.Update(row, column, chip.PriceGrowth7D)

		column++

		blueChipSheet.Update(row, column, chip.Volume24H)

		column++

		blueChipSheet.Update(row, column, chip.VolumeGrowth24H)

		column++

		blueChipSheet.Update(row, column, chip.PriceGrowth7D)

		column++

		blueChipSheet.Update(row, column, chip.Volume7D)

		column++

		blueChipSheet.Update(row, column, chip.VolumeGrowth7D)

		//End by iterating the row
		row++

	}

	// Make sure call Synchronize to reflect the changes.
	err = blueChipSheet.Synchronize()

	if err != nil {
		fmt.Println(err)
	}

}

func GetBlueChipsData() BlueChipsData {

	var data BlueChipsData

	data = GetPrice(data)

	return data

}

func GetPrice(data BlueChipsData) BlueChipsData {
	tickers, err := coinmarketcap.GetTickers()

	if err != nil {
		fmt.Println(err)
	}

	//Price growth here probably won't be standard.... Let's fix that and make sure all stuff is standard time.
	//I think we want to use the other api from coinmarket cap for this, but we can wait for that.

	var chips []BlueChip

	//We have to return 24H volume as well.
	volume, growth := ReturnVolumeAndGrowth("bitcoin")

	bitcoin := BlueChip{
		Name:           "Bitcoin",
		Slug:           "BTC",
		Price:          fmt.Sprintf("$%.2f", tickers["BTC"].Quotes["USD"].Price),
		PriceGrowth7D:  fmt.Sprintf("%.2f%%", tickers["BTC"].Quotes["USD"].PercentChange7D),
		PriceGrowth24H: fmt.Sprintf("%.2f%%", tickers["BTC"].Quotes["USD"].PercentChange24H),
		MarketCap:      fmt.Sprintf("%d", int(tickers["BTC"].Quotes["USD"].MarketCap/BILLION)),
		Volume7D:       volume,
		VolumeGrowth7D: growth,
	}

	chips = append(chips, bitcoin)

	volume, growth = ReturnVolumeAndGrowth("ethereum")

	eth := BlueChip{
		Name:           "Ethereum",
		Slug:           "ETH",
		Price:          fmt.Sprintf("$%.2f", tickers["ETH"].Quotes["USD"].Price),
		PriceGrowth7D:  fmt.Sprintf("%.2f%%", tickers["ETH"].Quotes["USD"].PercentChange7D),
		PriceGrowth24H: fmt.Sprintf("%.2f%%", tickers["ETH"].Quotes["USD"].PercentChange24H),
		MarketCap:      fmt.Sprintf("%d", int(tickers["ETH"].Quotes["USD"].MarketCap/BILLION)),
		Volume7D:       volume,
		VolumeGrowth7D: growth,
	}

	chips = append(chips, eth)

	volume, growth = ReturnVolumeAndGrowth("ripple")

	ripple := BlueChip{
		Name:           "Ripple",
		Slug:           "XRP",
		Price:          fmt.Sprintf("$%.2f", tickers["XRP"].Quotes["USD"].Price),
		PriceGrowth7D:  fmt.Sprintf("%.2f%%", tickers["XRP"].Quotes["USD"].PercentChange7D),
		PriceGrowth24H: fmt.Sprintf("%.2f%%", tickers["XRP"].Quotes["USD"].PercentChange24H),
		MarketCap:      fmt.Sprintf("%d", int(tickers["XRP"].Quotes["USD"].MarketCap/BILLION)),
		Volume7D:       volume,
		VolumeGrowth7D: growth,
	}

	chips = append(chips, ripple)

	volume, growth = ReturnVolumeAndGrowth("litecoin")

	litecoin := BlueChip{
		Name:           "Litecoin",
		Slug:           "LTC",
		Price:          fmt.Sprintf("$%.2f", tickers["LTC"].Quotes["USD"].Price),
		PriceGrowth7D:  fmt.Sprintf("%.2f%%", tickers["LTC"].Quotes["USD"].PercentChange7D),
		PriceGrowth24H: fmt.Sprintf("%.2f%%", tickers["LTC"].Quotes["USD"].PercentChange24H),
		MarketCap:      fmt.Sprintf("%d", int(tickers["LTC"].Quotes["USD"].MarketCap/BILLION)),
		Volume7D:       volume,
		VolumeGrowth7D: growth,
	}

	chips = append(chips, litecoin)

	data.chips = chips

	return data

}

func ReturnVolumeAndGrowth(currency string) (string, string) {

	start := time.Now().AddDate(0, 0, -8)
	end := start.AddDate(0, 0, 7)

	daily := coinmarketcap.GetHistoricalDailyByDate(currency, start, end)

	var totalVolume float64

	for _, day := range daily {
		volume, _ := day.Volume.Float64()
		totalVolume += volume
	}

	totalVolumeString := fmt.Sprintf("%d", int(totalVolume/BILLION))

	start = time.Now().AddDate(0, 0, -15)
	end = start.AddDate(0, 0, 7)

	daily = coinmarketcap.GetHistoricalDailyByDate(currency, start, end)

	var totalVolumeOld float64

	for _, day := range daily {
		volume, _ := day.Volume.Float64()
		totalVolumeOld += volume
	}

	volumeGrowth := ((totalVolume - totalVolumeOld) / totalVolumeOld) * 100

	volumeGrowthString := fmt.Sprintf("%.2f%%", volumeGrowth)

	return totalVolumeString, volumeGrowthString

}
