package main

import (
	"fmt"
	"math"
	"time"

	"github.com/Decarium/go-coinmarketcap/coinmarketcap"
	"github.com/bradfitz/slice"
	spreadsheet "gopkg.in/Iwark/spreadsheet.v2"
)

//TODO double check that we aren't double posting the date.

//TODO move the date around in the excel so that its month day year.
type SectorData struct {
	Sectors []Sector
}

type Sector struct {
	Name                string
	MarketCap           float64
	MarketCap24H        float64
	MarketCap7D         float64
	MarketCap1M         float64
	MarketCapYTD        float64
	GrowthPercentage24H float64
	GrowthPercentage7D  float64
	GrowthPercentage1M  float64
	GrowthPercentageYTD float64
	Negative            bool
	BarWidth            int
	NameList            []string
}

var (
	sheet         *spreadsheet.Sheet
	currency      []string
	dappPlatforms []string
	settlements   []string
	anon          []string
	digitalAssets []string
	crosschain    []string
)

func CreateSectors() {

	// get a sheet by the ID.
	var err error
	sheet, err = mainSheet.SheetByID(684580903)

	if err != nil {
		fmt.Println(err)
	}

	// t := template.New("sectors.tmpl")
	// t, err = t.ParseFiles("./templates/sectors.tmpl") // Parse template file.

	// if err != nil {
	// 	log.Fatal(err)
	// }

	data := GetSectorData()

	// f, err := os.Create("./pdf/sections/sectors/sectors.pug")
	// if err != nil {
	// 	log.Println("create file: ", err)
	// 	return
	// }

	// err = t.Execute(f, data)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// f.Close()

	BuildSpreadSheet(data)

}

func BuildSpreadSheet(data SectorData) {

	//We need to figure out a way to make sure that we aren't repeating the day.

	//Each new dataset is going to be 10 rows long. So we should check iterations of rows.

	// 9 rows actually...

	// if sheet.Rows[9][0].Value ==

	newRows := len(sheet.Rows) + 9

	newColumns := len(sheet.Columns)

	err := service.ExpandSheet(sheet, uint(newRows), uint(newColumns)) // Expand the sheet to 20 rows and 10 columns

	row := len(sheet.Rows) + 1

	column := 0

	//Get Today's Date
	now := time.Now()

	date := fmt.Sprintf("%d/%d/%d", now.Day(), int(now.Month()), now.Year())

	sheet.Update(row, column, date)

	column++

	sheet.Update(row, column, "24 Hr Change")

	column++

	sheet.Update(row, column, "7 Day Change")

	column++

	sheet.Update(row, column, "1 Month Change")

	column++

	sheet.Update(row, column, "YTD Change")

	//Go one more row down.
	row++

	//Iterate through sectors and add the name to the rows
	for _, sector := range data.Sectors {

		//Reset column
		column = 0

		sheet.Update(row, column, sector.Name)

		column++

		//24 HR change
		growth24H := fmt.Sprintf("%.2f%%", sector.GrowthPercentage24H)
		sheet.Update(row, column, growth24H)

		column++

		growth7D := fmt.Sprintf("%.2f%%", sector.GrowthPercentage7D)
		sheet.Update(row, column, growth7D)

		column++

		growth1M := fmt.Sprintf("%.2f%%", sector.GrowthPercentage1M)
		sheet.Update(row, column, growth1M)

		column++

		growthYTD := fmt.Sprintf("%.2f%%", sector.GrowthPercentageYTD)
		sheet.Update(row, column, growthYTD)

		//End by iterating the row
		row++
	}

	// Make sure call Synchronize to reflect the changes.
	err = sheet.Synchronize()

	if err != nil {
		fmt.Println(err)
	}

}

func GetSectorData() SectorData {

	DefineSectors()
	//Define our sectors
	c := Sector{
		Name:     "Currencies",
		NameList: currency,
	}

	p := Sector{
		Name:     "Platforms",
		NameList: dappPlatforms,
	}

	s := Sector{
		Name:     "Settlements",
		NameList: settlements,
	}

	a := Sector{
		Name:     "Anonymous Currencies",
		NameList: anon,
	}

	da := Sector{
		Name:     "Digital Assets",
		NameList: digitalAssets,
	}

	cc := Sector{
		Name:     "Interchain Network",
		NameList: crosschain,
	}

	var sectors []Sector

	//Append everything here
	sectors = append(sectors, c)
	sectors = append(sectors, p)
	sectors = append(sectors, s)
	sectors = append(sectors, a)
	sectors = append(sectors, da)
	sectors = append(sectors, cc)

	//Basically for each currency in each sector we are going to get their current market cap and add it to a total.

	for i, sector := range sectors {
		for _, name := range sector.NameList {
			var mcYTD float64
			//rename these too confusing
			od := 1
			sd := 6
			om := 30
			ytd := time.Now().YearDay()

			start := time.Now().AddDate(0, 0, -(ytd))
			end := start.AddDate(0, 0, ytd)

			data := coinmarketcap.GetHistoricalDailyByDate(name, start, end)

			mc24h, _ := data[od].MarketCap.Float64()
			mc7d, _ := data[sd].MarketCap.Float64()
			mc1m, _ := data[om].MarketCap.Float64()

			if len(data) < ytd {
				mcYTD = 0
			} else {
				mcYTD, _ = data[ytd-1].MarketCap.Float64()
			}

			mc, _ := data[0].MarketCap.Float64()

			sectors[i].MarketCap += mc
			sectors[i].MarketCap24H += mc24h
			sectors[i].MarketCap7D += mc7d
			sectors[i].MarketCap1M += mc1m
			sectors[i].MarketCapYTD += mcYTD

		}

		mc := sectors[i].MarketCap
		mc24h := sectors[i].MarketCap24H
		mc7d := sectors[i].MarketCap7D
		mc1m := sectors[i].MarketCap1M
		mcYTD := sectors[i].MarketCapYTD

		sectors[i].GrowthPercentage24H = ((mc - mc24h) / mc24h) * 100
		sectors[i].GrowthPercentage7D = ((mc - mc7d) / mc7d) * 100
		sectors[i].GrowthPercentage1M = ((mc - mc1m) / mc1m) * 100
		sectors[i].GrowthPercentageYTD = ((mc - mcYTD) / mcYTD) * 100
	}

	//Now we need to do just 2 thing which are these:
	// 1. Order them in an array by greatest to least
	slice.Sort(sectors[:], func(i, j int) bool {
		return sectors[i].GrowthPercentage7D < sectors[j].GrowthPercentage7D
	})

	//sectors is now sorted

	// 2. Between the top and the lowest calculate the width of the bar we show on the bwr
	//Basically do this: the top one is going to be 50~ pixels. The bottom one will be 50 in the opposite direction
	//Then just do the math and round between those.

	//lowest
	var multiplier float64

	lowest := math.Abs(sectors[0].GrowthPercentage7D)

	highest := math.Abs(sectors[len(sectors)-1].GrowthPercentage7D)

	if lowest > highest {
		//Get this magic number out of here
		multiplier = 75.00 / lowest
	} else {
		multiplier = 75.00 / highest
	}

	for i, s := range sectors {

		sectors[i].BarWidth = int(math.Abs(s.GrowthPercentage7D) * multiplier)

		if s.GrowthPercentage7D > 0 {
			sectors[i].Negative = false
		} else {
			sectors[i].Negative = true
		}

	}

	data := SectorData{
		Sectors: sectors,
	}

	return data

}

//This function simply defines the global variables for our sectors.
func DefineSectors() {
	currency = []string{"bitcoin", "bitcoin-cash", "litecoin", "iota", "tether", "bitcoin-gold", "nano", "bitcoin-diamond", "maker", "dogecoin"}
	dappPlatforms = []string{"ethereum", "eos", "tron", "neo", "ethereum-classic", "qtum", "zilliqa", "lisk", "aeternity", "bytom", "rchain"}
	settlements = []string{"ripple", "stellar", "dash"}
	anon = []string{"monero", "bytecoin-bcn", "verge", "zcash", "bitcoin-private"}
	digitalAssets = []string{"omisego", "nem", "bitshares", "populous", "waves", "wanchain", "vechain"}
	crosschain = []string{"ark", "icon", "ontology", "decred"}

	return
}
