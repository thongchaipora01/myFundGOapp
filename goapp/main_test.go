package main

import (
	"strconv"
	"testing"

	"github.com/gen2brain/dlgs"
	serv "github.com/thonp571/myfundgoapp/goapp/services"
)

func TestTimeRangeList(t *testing.T) {
	out, ret, err := dlgs.List("Choose interesting Time Range of Fund", "Select item from list:", []string{"1D", "1W", "1M", "1Y"})
	if err != nil {
		t.Error(err)
	}

	if out != "" {
		t.Log("out:", out, "ret:", ret)
	}
}

func TestSortedFundList(t *testing.T) {

	fund := serv.GetSuggestFund("1Y")
	if fund.Data != nil {
		t.Log(fund.Data[0])
	} else {
		t.Error("Null data")
	}

	var respFund []string
	// append Fund list to slice
	for i, v := range fund.Data {
		respFund = append(respFund, strconv.Itoa(i+1)+". "+v.Name)
	}
	if fund.Data != nil {
		t.Log(respFund[0])
	} else {
		t.Error("Null data")
	}

}

func TestSelectedFund(t *testing.T) {
	// make GUI of sorted Fund by nav_return
	out, ret, err := dlgs.List("Click for more Infomation", "Select item from list:", []string{"1. KT-OIL2", "2. ASP-VIETRMF2", "3. PRINCIPAL VNEQ-A", "4. PRINCIPAL VNEQ-I"})
	if err != nil {
		t.Error(err)
	}

	if out != "" {
		t.Log("out:", out, "ret:", ret)
	}
}

func TestShowFundInfo(t *testing.T) {
	//{KT-OIL2 1 2022-02-03T00:00:00.000Z 79.33466 3.4868}
	chosenFundInfo := "Name: " + "KT-OIL2" + "\n" + "Rank of fund: " + "1" + "\n" + "Updated date: " + "2022-02-03T00:00:00.000Z" + "\n" + "Performance: " + "79.33466" + "\n" + "Price: " + "3.4868"

	// Display a chosen Fund infomation
	ret, err := dlgs.Info("The Fund Infomation", chosenFundInfo)
	if err != nil {
		t.Error(err)
	}

	if ret != false {
		t.Log("out:", chosenFundInfo, "ret:", ret)
	}
}
