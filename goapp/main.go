package main

import (
	"strconv"
	"strings"

	"github.com/gen2brain/dlgs"
	serv "github.com/thonp571/myfundgoapp/goapp/services"
)

func main() {

	// make GUI of time range list for User choose 
	rangeSelected, _, err := dlgs.List("Choose interesting Time Range of Fund", "Select item from list:", []string{"1D", "1W", "1M", "1Y"})
	if err != nil {
		panic(err)
	}

	// query Sorted Fund
	fund := serv.GetSuggestFund(rangeSelected)
	var respFund []string
	// append Fund list to slice
	for i, v := range fund.Data {
		respFund = append(respFund, strconv.Itoa(i+1)+". "+v.Name)
	}
	// make GUI of sorted Fund by nav_return
	item, _, err := dlgs.List("Click for more Infomation", "Select item from list:", respFund)
	if err != nil {
		panic(err)
	}
	// Get Rank Number to string
	strChosenNo := item[0:3]
	strChosenNo = strings.ReplaceAll(strChosenNo, " ", "")
	strChosenNo = strings.ReplaceAll(strChosenNo, ".", "")

	// Convert string to int
	intChosenNo, _ := strconv.Atoi(strChosenNo)
	// Diff 1 for slice index of Fund data
	intChosenNo = intChosenNo - 1

	// Chosen Fund
	name := fund.Data[intChosenNo].Name
	rangeOfFund := strconv.Itoa(fund.Data[intChosenNo].RankOfFund)
	updatedDate := fund.Data[intChosenNo].UpdatedDate
	performance := strconv.FormatFloat(fund.Data[intChosenNo].Performance, 'f', -1, 64)
	price := strconv.FormatFloat(fund.Data[intChosenNo].Price, 'f', -1, 64)
	chosenFundInfo := "Name: " + name + "\n" + "Rank of fund: " + rangeOfFund + "\n" + "Updated date: " + updatedDate + "\n" + "Performance: " + performance + "\n" + "Price: " + price
	// Name (thailand_fund_code)
	// Rank of fund (already ranked by nav_return)
	// Updated date (nav_date)
	// Performance (nav_return)
	// Price (nav)

	// Display a chosen Fund infomation
	_, err = dlgs.Info("Choosen Fund Infomation", chosenFundInfo)
	if err != nil {
		panic(err)
	}

	// println("out:", item)

	// osArgsCheck := len(os.Args)
	// argsWithoutProg := ""
	// if osArgsCheck <= 1 {
	// 	argsWithoutProg = "develop"
	// } else {
	// 	argsWithoutProg = os.Args[1]
	// }

	// env := "GO_ENV"
	// os.Setenv(env, argsWithoutProg)
	// environment := os.Getenv(env)
	// if environment != "develop" {
	// 	err := "Please choose environment develop, staging or production"
	// 	panic(fmt.Errorf("fatal error config file: %s ", err))
	// }
	// log.Print("\"Info\": Environment is " + environment)
	// if environment == "develop" {
	// 	gin.SetMode(gin.DebugMode)
	// }

	// server := gin.Default()
	// server.Use(middleware.GetMiddlewareOption())
	// server.Use(cors.New(cors.Config{
	// 	AllowOrigins:  []string{"*"},
	// 	AllowMethods:  []string{"POST"},
	// 	AllowHeaders:  []string{"Origin"},
	// 	ExposeHeaders: []string{"Content-Length"},
	// }))

	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath("./config")
	// viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// err = viper.ReadInConfig()

	// if err != nil {
	// 	panic(fmt.Errorf("fatal error config file: %s ", err))
	// }

	// handler.InitializeRoutes(server)
	// port := viper.GetString(environment + ".port")

	// srv := &http.Server{
	// 	Addr:    ":" + port,
	// 	Handler: server,
	// }

	// if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
	// 	log.Fatalf("listen: %s\n", err)
	// }

}
