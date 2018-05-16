package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Decarium/go-coinmarketcap/coinmarketcap"
	"github.com/alecthomas/template"
	"github.com/davecgh/go-spew/spew"
)

type MarketCapChartData struct {
	Day1  int
	Day2  int
	Day3  int
	Day4  int
	Day5  int
	Day6  int
	Day7  int
	DayV1 int
	DayV2 int
	DayV3 int
	DayV4 int
	DayV5 int
	DayV6 int
	DayV7 int
}

//Main function to create the marketcap
func CreateMarketCapChart() {
	t := template.New("marketcapchart.tmpl")
	t, err := t.ParseFiles("./templates/marketcapchart.tmpl")

	if err != nil {
		log.Fatal(err)
	}

	data := GetMarketCapChartData()

	f, err := os.Create("./pdf/sections/charts/weeklyMarketCap.chart.js")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, data)

	if err != nil {
		fmt.Println(err)
	}

	f.Close()

}

func GetMarketCapChartData() MarketCapChartData {

	var data MarketCapChartData

	//Get Weekly Growth in ints

	data = GetDailyGrowthPastWeek(data)

	spew.Dump(data)

	return data

}

func GetDailyGrowthPastWeek(mccd MarketCapChartData) MarketCapChartData {

	t := time.Now().AddDate(0, 0, -8)

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	//So we do last week + 7 days so if this report is released on a Tuesday, we get Monday to Monday
	end := start.AddDate(0, 0, 7)

	global, err := coinmarketcap.GetGlobalHistoricalTicksDailyByDate(start, end)

	if err != nil {
		fmt.Println(err)
	}

	mccd.Day1 = int(global.MarketCapByAvailableSupply[0].Amount / BILLION)
	mccd.Day2 = int(global.MarketCapByAvailableSupply[1].Amount / BILLION)
	mccd.Day3 = int(global.MarketCapByAvailableSupply[2].Amount / BILLION)
	mccd.Day4 = int(global.MarketCapByAvailableSupply[3].Amount / BILLION)
	mccd.Day5 = int(global.MarketCapByAvailableSupply[4].Amount / BILLION)
	mccd.Day6 = int(global.MarketCapByAvailableSupply[5].Amount / BILLION)
	mccd.Day7 = int(global.MarketCapByAvailableSupply[6].Amount / BILLION)

	mccd.DayV1 = int(global.VolumeUsd[0].Amount / BILLION)
	mccd.DayV2 = int(global.VolumeUsd[1].Amount / BILLION)
	mccd.DayV3 = int(global.VolumeUsd[2].Amount / BILLION)
	mccd.DayV4 = int(global.VolumeUsd[3].Amount / BILLION)
	mccd.DayV5 = int(global.VolumeUsd[4].Amount / BILLION)
	mccd.DayV6 = int(global.VolumeUsd[5].Amount / BILLION)
	mccd.DayV7 = int(global.VolumeUsd[6].Amount / BILLION)

	return mccd

}
