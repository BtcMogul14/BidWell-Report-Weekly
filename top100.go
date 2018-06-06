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

//Should just make these into structs, and then put them in arrays - so 2 arrays
//with structs - next iteration
type Top100Data struct {
	Coin1Name                   string
	Coin1Slug                   string
	Coin1Price                  string
	Coin1PriceGrowthPercentage  string
	Coin1WebsiteSlug            string
	Coin2Name                   string
	Coin2Slug                   string
	Coin2Price                  string
	Coin2PriceGrowthPercentage  string
	Coin2WebsiteSlug            string
	Coin3Name                   string
	Coin3Slug                   string
	Coin3Price                  string
	Coin3PriceGrowthPercentage  string
	Coin3WebsiteSlug            string
	Loser1Name                  string
	Loser1Slug                  string
	Loser1Price                 string
	Loser1PriceGrowthPercentage string
	Loser1WebsiteSlug           string
	Loser2Name                  string
	Loser2Slug                  string
	Loser2Price                 string
	Loser2PriceGrowthPercentage string
	Loser2WebsiteSlug           string
	Loser3Name                  string
	Loser3Slug                  string
	Loser3Price                 string
	Loser3PriceGrowthPercentage string
	Loser3WebsiteSlug           string
}

type Top100ChartData struct {
	Day1           string
	Day2           string
	Day3           string
	Day4           string
	Day5           string
	Day6           string
	Day7           string
	Coin1Name      string
	Coin2Name      string
	Coin3Name      string
	Loser1Name     string
	Loser2Name     string
	Loser3Name     string
	Coin1Day1Gain  string
	Coin1Day2Gain  string
	Coin1Day3Gain  string
	Coin1Day4Gain  string
	Coin1Day5Gain  string
	Coin1Day6Gain  string
	Coin1Day7Gain  string
	Coin2Day1Gain  string
	Coin2Day2Gain  string
	Coin2Day3Gain  string
	Coin2Day4Gain  string
	Coin2Day5Gain  string
	Coin2Day6Gain  string
	Coin2Day7Gain  string
	Coin3Day1Gain  string
	Coin3Day2Gain  string
	Coin3Day3Gain  string
	Coin3Day4Gain  string
	Coin3Day5Gain  string
	Coin3Day6Gain  string
	Coin3Day7Gain  string
	Loser1Day1Gain string
	Loser1Day2Gain string
	Loser1Day3Gain string
	Loser1Day4Gain string
	Loser1Day5Gain string
	Loser1Day6Gain string
	Loser1Day7Gain string
	Loser2Day1Gain string
	Loser2Day2Gain string
	Loser2Day3Gain string
	Loser2Day4Gain string
	Loser2Day5Gain string
	Loser2Day6Gain string
	Loser2Day7Gain string
	Loser3Day1Gain string
	Loser3Day2Gain string
	Loser3Day3Gain string
	Loser3Day4Gain string
	Loser3Day5Gain string
	Loser3Day6Gain string
	Loser3Day7Gain string
}

func CreateTop100() {
	t := template.New("top100.tmpl")
	t, err := t.ParseFiles("./templates/top100.tmpl") // Parse template file.

	if err != nil {
		log.Fatal(err)
	}

	data := GetTop100Data()

	spew.Dump(data)

	f, err := os.Create("./pdf/sections/top100/top100.pug")
	if err != nil {
		log.Println("create file: ", err)
		return
	}

	err = t.Execute(f, data)

	if err != nil {
		fmt.Println(err)
	}

	f.Close()

	//Now Create the Chart
	CreateTop100Chart(data)
}

func CreateTop100Chart(dataOld Top100Data) {

	t := template.New("top100Chart.tmpl")
	t, err := t.ParseFiles("./templates/top100Chart.tmpl") // Parse template file.

	if err != nil {
		log.Fatal(err)
	}

	data := GetTop100ChartData(dataOld)

	spew.Dump(data)

	f, err := os.Create("./pdf/sections/charts/gainersAndLosers.chart.js")
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

func GetTop100Data() Top100Data {

	var data Top100Data

	tickers, err := coinmarketcap.GetTickers()

	if err != nil {
		fmt.Println(err)
	}

	//Should put these all into an array...

	var highest float64
	var secondHighest float64
	var thirdHighest float64

	var lowest float64
	var secondLowest float64
	var thirdLowest float64

	var highestSymbol string
	var secondHighestSymbol string
	var thirdHighestSymbol string

	var lowestSymbol string
	var secondLowestSymbol string
	var thirdLowestSymbol string

	for k, v := range tickers {
		change := v.Quotes["USD"].PercentChange7D

		if change > highest {
			thirdHighest = secondHighest
			thirdHighestSymbol = secondHighestSymbol

			secondHighest = highest
			secondHighestSymbol = highestSymbol

			highest = change
			highestSymbol = k

		} else if change > secondHighest {
			thirdHighest = secondHighest
			thirdHighestSymbol = secondHighestSymbol

			secondHighest = change
			secondHighestSymbol = k

		} else if change > thirdHighest {
			thirdHighest = change
			thirdHighestSymbol = k
		}

		if change < lowest {
			thirdLowest = secondLowest
			thirdLowestSymbol = secondLowestSymbol

			secondLowest = lowest
			secondLowestSymbol = lowestSymbol

			lowest = change
			lowestSymbol = k

		} else if change < secondLowest {
			thirdLowest = secondLowest
			thirdLowestSymbol = secondLowestSymbol

			secondLowest = change
			secondLowestSymbol = k

		} else if change < thirdLowest {
			thirdLowest = change
			thirdLowestSymbol = k
		}

	}

	data.Coin1Name = tickers[highestSymbol].Name
	data.Coin1Slug = highestSymbol
	//Do 2 decimals for now, but we will have to figure out how to show this better.
	data.Coin1Price = fmt.Sprintf("$%.2f", tickers[highestSymbol].Quotes["USD"].Price)
	data.Coin1PriceGrowthPercentage = fmt.Sprintf("%.2f%%", highest)
	data.Coin1WebsiteSlug = tickers[highestSymbol].Slug

	data.Coin2Name = tickers[secondHighestSymbol].Name
	data.Coin2Slug = secondHighestSymbol
	data.Coin2Price = fmt.Sprintf("$%.2f", tickers[secondHighestSymbol].Quotes["USD"].Price)
	data.Coin2PriceGrowthPercentage = fmt.Sprintf("%.2f%%", secondHighest)
	data.Coin2WebsiteSlug = tickers[secondHighestSymbol].Slug

	data.Coin3Name = tickers[thirdHighestSymbol].Name
	data.Coin3Slug = thirdHighestSymbol
	data.Coin3Price = fmt.Sprintf("$%.2f", tickers[thirdHighestSymbol].Quotes["USD"].Price)
	data.Coin3PriceGrowthPercentage = fmt.Sprintf("%.2f%%", thirdHighest)
	data.Coin3WebsiteSlug = tickers[thirdHighestSymbol].Slug

	data.Loser1Name = tickers[lowestSymbol].Name
	data.Loser1Slug = lowestSymbol
	//Do 2 decimals for now, but we will have to figure out how to show this better.
	data.Loser1Price = fmt.Sprintf("$%.2f", tickers[lowestSymbol].Quotes["USD"].Price)
	data.Loser1PriceGrowthPercentage = fmt.Sprintf("%.2f%%", lowest)
	data.Loser1WebsiteSlug = tickers[lowestSymbol].Slug

	data.Loser2Name = tickers[secondLowestSymbol].Name
	data.Loser2Slug = secondLowestSymbol
	data.Loser2Price = fmt.Sprintf("$%.2f", tickers[secondLowestSymbol].Quotes["USD"].Price)
	data.Loser2PriceGrowthPercentage = fmt.Sprintf("%.2f%%", secondLowest)
	data.Loser2WebsiteSlug = tickers[secondLowestSymbol].Slug

	data.Loser3Name = tickers[thirdLowestSymbol].Name
	data.Loser3Slug = thirdLowestSymbol
	data.Loser3Price = fmt.Sprintf("$%.2f", tickers[thirdLowestSymbol].Quotes["USD"].Price)
	data.Loser3PriceGrowthPercentage = fmt.Sprintf("%.2f%%", thirdLowest)
	data.Loser3WebsiteSlug = tickers[thirdLowestSymbol].Slug

	return data

}

func GetTop100ChartData(data Top100Data) Top100ChartData {

	var newData Top100ChartData

	newData.Coin1Name = data.Coin1Name
	newData.Coin2Name = data.Coin2Name
	newData.Coin3Name = data.Coin3Name
	newData.Loser1Name = data.Loser1Name
	newData.Loser2Name = data.Loser2Name
	newData.Loser3Name = data.Loser3Name

	t := time.Now().AddDate(0, 0, -8)

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	end := start.AddDate(0, 0, 8)

	spew.Dump(data.Coin1PriceGrowthPercentage)

	//We may need website slug here so may need a revamp.
	days := coinmarketcap.GetHistoricalDailyByDate(data.Coin1WebsiteSlug, start, end)

	//days[0] is just for calculating % gained on day 1.
	newData.Day7 = fmt.Sprintf("%d/%d", days[1].Date.Month(), days[1].Date.Day())
	newData.Day6 = fmt.Sprintf("%d/%d", days[2].Date.Month(), days[2].Date.Day())
	newData.Day5 = fmt.Sprintf("%d/%d", days[3].Date.Month(), days[3].Date.Day())
	newData.Day4 = fmt.Sprintf("%d/%d", days[4].Date.Month(), days[4].Date.Day())
	newData.Day3 = fmt.Sprintf("%d/%d", days[5].Date.Month(), days[5].Date.Day())
	newData.Day2 = fmt.Sprintf("%d/%d", days[6].Date.Month(), days[6].Date.Day())
	newData.Day1 = fmt.Sprintf("%d/%d", days[7].Date.Month(), days[7].Date.Day())

	day0Close, _ := days[7].Close.Float64()
	day1Close, _ := days[6].Close.Float64()
	day2Close, _ := days[5].Close.Float64()
	day3Close, _ := days[4].Close.Float64()
	day4Close, _ := days[3].Close.Float64()
	day5Close, _ := days[2].Close.Float64()
	day6Close, _ := days[1].Close.Float64()
	day7Close, _ := days[0].Close.Float64()

	//All this data should be at the close of Market Monday 7PM so 00:00 UTC

	newData.Coin1Day1Gain = fmt.Sprintf("%.2f", (100 * (day1Close - day0Close) / day0Close))
	newData.Coin1Day2Gain = fmt.Sprintf("%.2f", (100 * (day2Close - day0Close) / day0Close))
	newData.Coin1Day3Gain = fmt.Sprintf("%.2f", (100 * (day3Close - day0Close) / day0Close))
	newData.Coin1Day4Gain = fmt.Sprintf("%.2f", (100 * (day4Close - day0Close) / day0Close))
	newData.Coin1Day5Gain = fmt.Sprintf("%.2f", (100 * (day5Close - day0Close) / day0Close))
	newData.Coin1Day6Gain = fmt.Sprintf("%.2f", (100 * (day6Close - day0Close) / day0Close))
	newData.Coin1Day7Gain = fmt.Sprintf("%.2f", (100 * (day7Close - day0Close) / day0Close))

	days = coinmarketcap.GetHistoricalDailyByDate(data.Coin2WebsiteSlug, start, end)

	day0Close, _ = days[7].Close.Float64()
	day1Close, _ = days[6].Close.Float64()
	day2Close, _ = days[5].Close.Float64()
	day3Close, _ = days[4].Close.Float64()
	day4Close, _ = days[3].Close.Float64()
	day5Close, _ = days[2].Close.Float64()
	day6Close, _ = days[1].Close.Float64()
	day7Close, _ = days[0].Close.Float64()

	newData.Coin2Day1Gain = fmt.Sprintf("%.2f", (100 * (day1Close - day0Close) / day0Close))
	newData.Coin2Day2Gain = fmt.Sprintf("%.2f", (100 * (day2Close - day0Close) / day0Close))
	newData.Coin2Day3Gain = fmt.Sprintf("%.2f", (100 * (day3Close - day0Close) / day0Close))
	newData.Coin2Day4Gain = fmt.Sprintf("%.2f", (100 * (day4Close - day0Close) / day0Close))
	newData.Coin2Day5Gain = fmt.Sprintf("%.2f", (100 * (day5Close - day0Close) / day0Close))
	newData.Coin2Day6Gain = fmt.Sprintf("%.2f", (100 * (day6Close - day0Close) / day0Close))
	newData.Coin2Day7Gain = fmt.Sprintf("%.2f", (100 * (day7Close - day0Close) / day0Close))

	days = coinmarketcap.GetHistoricalDailyByDate(data.Coin3WebsiteSlug, start, end)

	day0Close, _ = days[7].Close.Float64()
	day1Close, _ = days[6].Close.Float64()
	day2Close, _ = days[5].Close.Float64()
	day3Close, _ = days[4].Close.Float64()
	day4Close, _ = days[3].Close.Float64()
	day5Close, _ = days[2].Close.Float64()
	day6Close, _ = days[1].Close.Float64()
	day7Close, _ = days[0].Close.Float64()

	newData.Coin3Day1Gain = fmt.Sprintf("%.2f", (100 * (day1Close - day0Close) / day0Close))
	newData.Coin3Day2Gain = fmt.Sprintf("%.2f", (100 * (day2Close - day0Close) / day0Close))
	newData.Coin3Day3Gain = fmt.Sprintf("%.2f", (100 * (day3Close - day0Close) / day0Close))
	newData.Coin3Day4Gain = fmt.Sprintf("%.2f", (100 * (day4Close - day0Close) / day0Close))
	newData.Coin3Day5Gain = fmt.Sprintf("%.2f", (100 * (day5Close - day0Close) / day0Close))
	newData.Coin3Day6Gain = fmt.Sprintf("%.2f", (100 * (day6Close - day0Close) / day0Close))
	newData.Coin3Day7Gain = fmt.Sprintf("%.2f", (100 * (day7Close - day0Close) / day0Close))

	days = coinmarketcap.GetHistoricalDailyByDate(data.Loser1WebsiteSlug, start, end)

	day0Close, _ = days[7].Close.Float64()
	day1Close, _ = days[6].Close.Float64()
	day2Close, _ = days[5].Close.Float64()
	day3Close, _ = days[4].Close.Float64()
	day4Close, _ = days[3].Close.Float64()
	day5Close, _ = days[2].Close.Float64()
	day6Close, _ = days[1].Close.Float64()
	day7Close, _ = days[0].Close.Float64()

	newData.Loser1Day1Gain = fmt.Sprintf("%.2f", (100 * (day1Close - day0Close) / day0Close))
	newData.Loser1Day2Gain = fmt.Sprintf("%.2f", (100 * (day2Close - day0Close) / day0Close))
	newData.Loser1Day3Gain = fmt.Sprintf("%.2f", (100 * (day3Close - day0Close) / day0Close))
	newData.Loser1Day4Gain = fmt.Sprintf("%.2f", (100 * (day4Close - day0Close) / day0Close))
	newData.Loser1Day5Gain = fmt.Sprintf("%.2f", (100 * (day5Close - day0Close) / day0Close))
	newData.Loser1Day6Gain = fmt.Sprintf("%.2f", (100 * (day6Close - day0Close) / day0Close))
	newData.Loser1Day7Gain = fmt.Sprintf("%.2f", (100 * (day7Close - day0Close) / day0Close))

	days = coinmarketcap.GetHistoricalDailyByDate(data.Loser2WebsiteSlug, start, end)

	day0Close, _ = days[7].Close.Float64()
	day1Close, _ = days[6].Close.Float64()
	day2Close, _ = days[5].Close.Float64()
	day3Close, _ = days[4].Close.Float64()
	day4Close, _ = days[3].Close.Float64()
	day5Close, _ = days[2].Close.Float64()
	day6Close, _ = days[1].Close.Float64()
	day7Close, _ = days[0].Close.Float64()

	newData.Loser2Day1Gain = fmt.Sprintf("%.2f", (100 * (day1Close - day0Close) / day0Close))
	newData.Loser2Day2Gain = fmt.Sprintf("%.2f", (100 * (day2Close - day0Close) / day0Close))
	newData.Loser2Day3Gain = fmt.Sprintf("%.2f", (100 * (day3Close - day0Close) / day0Close))
	newData.Loser2Day4Gain = fmt.Sprintf("%.2f", (100 * (day4Close - day0Close) / day0Close))
	newData.Loser2Day5Gain = fmt.Sprintf("%.2f", (100 * (day5Close - day0Close) / day0Close))
	newData.Loser2Day6Gain = fmt.Sprintf("%.2f", (100 * (day6Close - day0Close) / day0Close))
	newData.Loser2Day7Gain = fmt.Sprintf("%.2f", (100 * (day7Close - day0Close) / day0Close))

	days = coinmarketcap.GetHistoricalDailyByDate(data.Loser3WebsiteSlug, start, end)

	day0Close, _ = days[7].Close.Float64()
	day1Close, _ = days[6].Close.Float64()
	day2Close, _ = days[5].Close.Float64()
	day3Close, _ = days[4].Close.Float64()
	day4Close, _ = days[3].Close.Float64()
	day5Close, _ = days[2].Close.Float64()
	day6Close, _ = days[1].Close.Float64()
	day7Close, _ = days[0].Close.Float64()

	newData.Loser3Day1Gain = fmt.Sprintf("%.2f", (100 * (day1Close - day0Close) / day0Close))
	newData.Loser3Day2Gain = fmt.Sprintf("%.2f", (100 * (day2Close - day0Close) / day0Close))
	newData.Loser3Day3Gain = fmt.Sprintf("%.2f", (100 * (day3Close - day0Close) / day0Close))
	newData.Loser3Day4Gain = fmt.Sprintf("%.2f", (100 * (day4Close - day0Close) / day0Close))
	newData.Loser3Day5Gain = fmt.Sprintf("%.2f", (100 * (day5Close - day0Close) / day0Close))
	newData.Loser3Day6Gain = fmt.Sprintf("%.2f", (100 * (day6Close - day0Close) / day0Close))
	newData.Loser3Day7Gain = fmt.Sprintf("%.2f", (100 * (day7Close - day0Close) / day0Close))

	return newData
}
