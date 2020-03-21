package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	IP_ENV             = "IP_ENV"
	PORT_ENV           = "PORT_ENV"
	URL_ENV            = "URL_ENV"
	BATCH_CAP_ENV      = "BATCH_CAP_ENV"
	NAME_OF_ACTION_ENV = "NAME_OF_ACTION_ENV"

	//TAGS

	MOCK_MODE_ENV = "MOCK_MODE_ENV"
	BOT_TOKEN     = "BOT_TOKEN"
)

type Config struct {
	HOST_IP        string
	PORT           string
	URL            string
	BATCH_CAPACITY int
	NAME_OF_ACTION string
	MOCK_MODE      bool
	TOKEN          string
	TIME_OF_START  time.Time
}

func ConfigIni() (conf *Config) {

	conf = &Config{}
	exist := false

	conf.HOST_IP, exist = os.LookupEnv(IP_ENV)
	if !exist {
		log.Fatal("NOT FOUND IP")
	}

	conf.PORT, exist = os.LookupEnv(PORT_ENV)
	if !exist {
		log.Fatal("NOT FOUND PORT")
	}

	BATCH_CAPACITY, exist := os.LookupEnv(BATCH_CAP_ENV)
	if !exist {
		log.Fatal("NOT FOUND BATCH CAPACITY")
	}
	conf.BATCH_CAPACITY, _ = strconv.Atoi(BATCH_CAPACITY)

	conf.URL, exist = os.LookupEnv(URL_ENV)
	if !exist {
		log.Fatal("NOT FOUND URL")
	}

	mock_mode, exist := os.LookupEnv(MOCK_MODE_ENV)
	if !exist {
		log.Fatal("NOT FOUND MOCK_MODE_ENV")
	}

	fmt.Println("########################################")
	if strings.ToLower(mock_mode) == "true" {
		conf.MOCK_MODE = true
		fmt.Println("######### WORKING IN MOCK MODE #########")
	} else {
		conf.MOCK_MODE = false
		fmt.Println("######### WORKING IN REAL MODE #########")
	}
	fmt.Println("########################################")

	conf.TOKEN, exist = os.LookupEnv(BOT_TOKEN)
	if !exist {
		log.Fatal("NOT FOUND TELEGRAM token ENV")
	}
	return
}
