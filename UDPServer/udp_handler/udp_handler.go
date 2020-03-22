package udp_handler

import (
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/batch"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/config"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/tags"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

//import "../models"

const (
	PAGE_SIZE = 4
	ACC_X     = 0
	ACC_Y     = 1
	ACC_Z     = 2
	GYRO_X    = 3
	GYRO_Y    = 4
	GYRO_Z    = 5
	MS        = 6
)

type Udp_handler struct {
	Conn           *net.UDPConn
	GLOBAL_BATCH   []batch.Batch
	GLOBAL_COUNTER int
	LOCAL_COUNTER  int
	Tags           *tags.Tags
	Conf           *config.Config
}

//var GLOBAL_BATCH = make([]batch.Batch, PAGE_SIZE)
//var GLOBAL_COUNTER int
//var GLOBAL_TAGS models.Tags

func New(c *config.Config) (u *Udp_handler) {

	u = &Udp_handler{}

	// Генерируем конфиг
	u.Conf = c
	fmt.Println("CONFIG READY")
	// Генерируем Тэги
	u.Tags = tags.TagsIni()
	fmt.Println("TAGS READY")

	fmt.Println("YOUR NAME: ", u.Tags.User)
	fmt.Println("EXERCISE: ", u.Tags.Exercise)
	fmt.Println("REPEATS: ", u.Tags.Repeats)
	fmt.Println("")

	// Выделяем место под батчи
	u.GLOBAL_BATCH = make([]batch.Batch, PAGE_SIZE)
	for i := 0; i < len(u.GLOBAL_BATCH); i++ {
		u.GLOBAL_BATCH[i] = *batch.InitBatch(u.Conf.BATCH_CAPACITY)
	}

	// Открываем UDP сокет
	udpAddr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%s", u.Conf.HOST_IP, u.Conf.PORT))
	if err != nil {
		log.Fatal(err)
	}
	u.Conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("UDP SERVER IS READY " + fmt.Sprintf("%s:%s", u.Conf.HOST_IP, u.Conf.PORT))

	return
}

// В бесконечном цикле слушаем что приходит, накапливаем пакеты, формируем батчи, отправляем батчи по URL,
// меняем страницы.
func (u *Udp_handler) Handle() {
	//return

	buffer := make([]byte, 512)
	_, _, err := u.Conn.ReadFromUDP(buffer)

	if err != nil {
		fmt.Println(err)
		return
	} else {
		data := strings.Split(string(buffer), ";")

		// TAKING PACKET FROM JSON
		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Fields.AccX = json.Number(data[ACC_X])

		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Fields.AccY = json.Number(data[ACC_Y])

		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Fields.AccZ = json.Number(data[ACC_Z])

		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Fields.GyroX = json.Number(data[GYRO_X])

		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Fields.GyroY = json.Number(data[GYRO_Y])

		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Fields.GyroZ = json.Number(data[GYRO_Z])

		// TAGS TAKING
		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Tags.DeviceId = u.Tags.DeviceId
		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Tags.User = u.Tags.User
		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Tags.Ms, err = strconv.Atoi(data[MS])
		if err != nil {
			fmt.Println("ERROR OF GETTING MS FROM CIRCUIT")
			log.Fatal(err)
			return
		}

		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Measurement = u.Tags.Exercise

		u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Time = time.Now()
		defTime := time.Time{}
		if u.Conf.TIME_OF_START == defTime {
			u.Conf.TIME_OF_START = u.GLOBAL_BATCH[u.GLOBAL_COUNTER].DataArray[u.LOCAL_COUNTER].Time
		}
		u.LOCAL_COUNTER++
		if u.LOCAL_COUNTER%100 == 0 {
			//jso, _ := json.Marshal(GLOBAL_BATCH[GLOBAL_COUNTER])
			//fmt.Println(string(jso))
			fmt.Println("LOCAL COUNTER ", u.LOCAL_COUNTER)
			fmt.Println("BATCH CAPACITY ", u.Conf.BATCH_CAPACITY)
		}

		if u.LOCAL_COUNTER == u.Conf.BATCH_CAPACITY {
			fmt.Println("SETTING BATCH TO DB")
			old := u.GLOBAL_COUNTER
			go u.GLOBAL_BATCH[old].Send(u.Conf)
			if err != nil {
				log.Fatal("BATCH SENDING ERROR: ", err)
			}
			u.LOCAL_COUNTER = 0
			u.GLOBAL_COUNTER++
			if u.GLOBAL_COUNTER == PAGE_SIZE {
				u.GLOBAL_COUNTER = 0
			}
		}
	}
}
