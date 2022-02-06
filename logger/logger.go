package logger

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/spf13/viper"
)

//Initial log
var (
	LogReq  *log.Logger
	LogRes  *log.Logger
	FileReq *os.File
	FileRes *os.File
)

func RefreshLog() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	osArgsCheck := len(os.Args)
	environment := ""
	if osArgsCheck <= 1 {
		environment = "develop"
	} else {
		environment = os.Args[1]
	}

	//Format
	currentDate := time.Now()
	dateString := currentDate.Format("20060102")
	// timenow = strings.Replace(timenow, ":", " ", -1)

	// set location of log file
	logpathReq := viper.GetString(environment+".log.pathLogReq") + "_" + dateString + ".log"
	logpathRes := viper.GetString(environment+".log.pathLogRes") + "_" + dateString + ".log"

	flag.Parse()
	var err1 error
	FileReq, err1 = os.OpenFile(logpathReq, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err1 != nil {
		log.Fatal(err1)
	}

	var err2 error
	FileRes, err2 = os.OpenFile(logpathRes, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err2 != nil {
		log.Fatal(err2)
	}

	LogReq = log.New(FileReq, "[Request INFO]: ", log.Ldate|log.Ltime)
	LogRes = log.New(FileRes, "[Response INFO]: ", log.Ldate|log.Ltime)

}
