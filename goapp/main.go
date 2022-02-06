package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/thonp571/myfundgoapp/goapp/handler"
	"github.com/thonp571/myfundgoapp/middleware"
)

type reviews_data struct {
	review_id string
	date      time.Time
}

func main() {

	fmt.Println("Sort Example")
	var listOfReviews = make([]reviews_data, 0)
	listOfReviews = append(listOfReviews, reviews_data{review_id: "1", date: time.Now()})
	listOfReviews = append(listOfReviews, reviews_data{review_id: "2", date: time.Now().AddDate(0, 0, 7*1)})
	listOfReviews = append(listOfReviews, reviews_data{review_id: "3", date: time.Now().AddDate(0, 0, 7*-1)})
	fmt.Println(listOfReviews)

	sort.Slice(listOfReviews, func(i, j int) bool { return listOfReviews[i].date.Before(listOfReviews[j].date) })
	fmt.Println(listOfReviews)

	osArgsCheck := len(os.Args)
	argsWithoutProg := ""
	if osArgsCheck <= 1 {
		argsWithoutProg = "develop"
	} else {
		argsWithoutProg = os.Args[1]
	}

	env := "GO_ENV"
	os.Setenv(env, argsWithoutProg)
	environment := os.Getenv(env)
	if environment != "develop" {
		err := "Please choose environment develop, staging or production"
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}
	log.Print("\"Info\": Environment is " + environment)
	if environment == "develop" {
		gin.SetMode(gin.DebugMode)
	}

	server := gin.Default()
	server.Use(middleware.GetMiddlewareOption())
	server.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"POST"},
		AllowHeaders:  []string{"Origin"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s ", err))
	}

	handler.InitializeRoutes(server)
	port := viper.GetString(environment + ".port")

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: server,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}

}
