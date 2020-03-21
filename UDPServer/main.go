package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/tg_logger"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/udp_handler"
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

	udp := udp_handler.New()

	tglog := tg_logger.New(udp.Conf)
	fmt.Println("TESTING PROXY MODE")
	err := tglog.TestProxy()
	if err != nil {
		return
	}

	//udp.Handle()
	err = keyboard.Open()
	if err != nil {
		panic(err)
	}
	defer keyboard.Close()

	in := make(chan keyboard.Key)
	go EscHandler(in)

loop:
	for {
		select {
		case char := <-in:
			if char == keyboard.KeyEsc {
				fmt.Println("TRAIN COMPLETE\n")
				l := tg_logger.NewLog(udp.Tags, udp.Conf)
				// Отправка лога в тг
				tglog.SendLogInChannel(l)
				break loop
			}
		default:
			udp.Handle()
		}
	}
	fmt.Println("END OF BATCHING")

}

func EscHandler(input chan int) {
	for {
		_, char, _ := keyboard.GetKey()
		input <- char
		if char == keyboard.KeyEsc {
			keyboard.Close()
			break
		}
	}
}
