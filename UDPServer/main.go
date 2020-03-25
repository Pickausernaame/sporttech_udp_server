package main

import (
	"fmt"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/batch"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/config"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/models"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/tg_logger"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/udp_handler"
	"github.com/eiannone/keyboard"
)

func main() {

	//// Если тестируем без IMU
	//if conf.MOCK_MODE {
	//	// Создаем свои батчи
	//	for i := 0; i < len(GLOBAL_BATCH); i++ {
	//		GLOBAL_BATCH[i] = *initBatch(conf.BATCH_CAPACITY)
	//		GLOBAL_BATCH[i].mock()
	//	}
	//
	//
	//	// Отпраляем батчи в по URL
	//	for i := 0; i < PAGE_SIZE; i++ {
	//		SendBatch(&GLOBAL_BATCH[i])
	//		time.Sleep(time.Second * 10)
	//	}
	//} else { // Работаем с IMU
	ch := make(chan error)
	c := config.ConfigIni()
	tglog := tg_logger.New(c)
	fmt.Println("TESTING PROXY MODE")
	err := tglog.TestProxy()
	if err != nil {
		return
	}

	udp := udp_handler.New(c, ch)
	fmt.Println("UDP CREATED")

	//udp.Handle()
	err = keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	in := make(chan keyboard.Key)
	go EscHandler(in)
	badReqs := 0
loop:
	for {
		select {
		case char := <-in:
			if char == keyboard.KeyEsc {
				fmt.Println("TRAIN COMPLETE\n")
				l := tg_logger.NewLog(udp.Tags, udp.Conf)
				// Отправка лога в тг
				tglog.SendLogInChannel(l)
				b := batch.Batch{}
				b.DataArray = make([]models.Data, udp.LOCAL_COUNTER)
				for i := 0; i < udp.LOCAL_COUNTER; i++ {
					b.DataArray = append(b.DataArray, udp.GLOBAL_BATCH[udp.GLOBAL_COUNTER].DataArray[i])
				}
				b.Send(udp.Conf, udp.ErrorChan)
				break loop
			}
			break
		case err := <- ch:
			fmt.Println(err)
			badReqs = badReqs + 1
			if badReqs > 2 {
				tglog.SendBadLogInChannel()
				break loop
			}
			break
		default:
			udp.Handle()
		}
	}
	fmt.Println("END OF BATCHING")
}

func EscHandler(input chan keyboard.Key) {
	for {
		fmt.Println("PRESS ESC TO STOP TRAIN")
		_, char, _ := keyboard.GetKey()
		input <- char
		if char == keyboard.KeyEsc {
			keyboard.Close()
			break
		}

	}
}
