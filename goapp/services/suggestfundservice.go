package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/thonp571/myfundgoapp/goapp/model"
)

func GetSuggestFund(timeRange string) model.ResFundArr {
	// str := ""
	// find file path
	filePath, err := filepath.Abs("../fundsdata.json")
	if err != nil {
		fmt.Println(err)
	}
	// Open file path
	jsonFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// we initialize our Users array
	fundArr := model.FundArr{}

	// we unmarshal our byteArray which contains our jsonFile
	json.Unmarshal(byteValue, &fundArr)

	sortedFund := sortFund(fundArr)

	rangeFund := splitFundAndChooseRange(sortedFund, timeRange)

	var respFund model.ResFundArr
	for i, v := range rangeFund.Data {
		respFund.Data = append(respFund.Data, model.ResFundData{
			Name:        v.Thailand_fund_code,
			RankOfFund:  i + 1,
			UpdatedDate: v.Nav_date,
			Performance: v.Nav_return,
			Price:       v.Nav,
		})
	}

	return respFund
}

// Append Go_date AND sort Slice by price Performance
func sortFund(fundArr model.FundArr) model.FundArr {
	for i := 0; i < len(fundArr.Data); i++ {
		// Parse string to time.Time
		tJson, err := time.Parse(time.RFC3339, fundArr.Data[i].Nav_date)
		if err != nil {
			fmt.Println("err:", err)
		}
		fundArr.Data[i].Go_date = tJson
	}

	// sort Fund by price Performance
	sort.Slice(fundArr.Data, func(i, j int) bool { return fundArr.Data[i].Nav_return > fundArr.Data[j].Nav_return })
	return fundArr
}

// Split Fund by day,week,month and year into Slice
// Then select Range that user request
func splitFundAndChooseRange(sortedFund model.FundArr, timeRange string) (chosenFund model.FundArr) {
	var dayFund model.FundArr
	var weekFund model.FundArr
	var monthFund model.FundArr
	var yearFund model.FundArr
	for i := 0; i < len(sortedFund.Data); i++ {
		// diff time.Now().Sub(tJson)
		// find diff from Data date and time.Now()
		diffHours := int(time.Since(sortedFund.Data[i].Go_date).Hours())
		if diffHours >= 0 && diffHours <= 23 { // 24 hrs = 1 day
			dayFund.Data = append(dayFund.Data, sortedFund.Data[i])
		} else if diffHours >= 0 && diffHours <= 167 { // 168 = 1 week
			weekFund.Data = append(weekFund.Data, sortedFund.Data[i])
		} else if diffHours >= 0 && diffHours <= 729 { // 730 = 1 month
			monthFund.Data = append(monthFund.Data, sortedFund.Data[i])
		} else if diffHours >= 0 && diffHours <= 8764 { // 8765 = 1 year
			yearFund.Data = append(yearFund.Data, sortedFund.Data[i])
		}
	}

	timeRange = strings.ToLower(timeRange)
	if timeRange == "1d" {
		chosenFund = dayFund
	} else if timeRange == "1w" {
		chosenFund = weekFund
	} else if timeRange == "1m" {
		chosenFund = monthFund
	} else if timeRange == "1y" {
		chosenFund = yearFund
	} else { // more than 1 year
		chosenFund = sortedFund
	}
	return
}
