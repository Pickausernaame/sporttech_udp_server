package tg_logger

import (
	"fmt"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/config"
	proxy "github.com/Pickausernaame/sporttech_udp_server/UDPServer/http_proxy_client"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
)

const (
	CHAT_ID = -1001191905185
)

type TgLogger struct {
	Conf   *config.Config
	Client *http.Client
}

func New(c *config.Config) (l *TgLogger) {
	l = &TgLogger{
		Conf:   c,
		Client: proxy.NewClientWithProxy(),
	}
	return
}

func (l *TgLogger) TestProxy() error {
	_, err := tgbotapi.NewBotAPIWithClient(l.Conf.TOKEN, l.Client)
	if err != nil {
		fmt.Println("BAD PROXY")
		return err
	}
	fmt.Println("OK")
	return nil
}

func (logger *TgLogger) SendLogInChannel(l *Log) {

	msgTemplate := `
	NAME: %s
	ACTION: %s
	REPEATS: %s
	TIME OF START: %s
	TIME OF END: %s
    `
	msg := fmt.Sprintf(msgTemplate, l.Username, l.Exercise, l.Repeats, l.Time_of_start.Format("2006-01-02T15:04:05"), l.Time_of_end.Format("2006-01-02T15:04:05"))
	bot, err := tgbotapi.NewBotAPIWithClient(logger.Conf.TOKEN, logger.Client)
	if err != nil {
		fmt.Println("SOMETHING FALIED")
		fmt.Println("DONT WORRY. ALL YOUR BATCHES ON SERVER")
		fmt.Println("PLEASE SCREEN YOUR RESULTS AND POST IT ON TELEGRAM CHANEL\n")
		fmt.Println("YOUR RESULTS:")
		fmt.Println(msg)

		log.Panic(err)
	}
	_, err = bot.Send(tgbotapi.NewMessage(CHAT_ID, msg))
	if err != nil {
		fmt.Println(err)
	}
	return
}
