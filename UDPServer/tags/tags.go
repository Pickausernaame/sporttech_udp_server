package tags

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	NAME_OF_USER_ENV = "NAME_OF_USER_ENV"
	DEVICE_ID_ENV    = "DEVICE_ID_ENV"
)

type Tags struct {
	User     string `json:"user"`
	DeviceId int    `json:"deviceid"`
	Ms       int    `json:"millis_time"`
	Exercise string
	Repeats  string
}

var Exercises = [4]string{"squats", "push_ups", "chin_ups", "abs"}

func TagsIni() (t *Tags) {

	msg := `
Supporting exercises:
1) squats (приседания)
2) push_ups (отжимания)
3) chin_ups (подтягивания)
4) abs (пресс)
`

	t = &Tags{}
	exist := false
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Insert your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Print("Insert your repeats: ")
	repeats, _ := reader.ReadString('\n')
	t.User = name + "_" + time.Now().Format("2006-01-02T15:04:05") + "_" + repeats
	t.Repeats = repeats
	fmt.Println("Your name tag: ", t.User)
	fmt.Println(msg)
	for {
		fmt.Print("Insert number of exercise: ")
		num, _ := reader.ReadString('\n')
		switch num[:len(num)-2] {
		case "1":
			t.Exercise = Exercises[0]
			break
		case "2":
			t.Exercise = Exercises[1]
			break
		case "3":
			t.Exercise = Exercises[2]
			break
		case "4":
			t.Exercise = Exercises[3]
			break
		default:
			fmt.Println("Wrong numb, please insert new numb")
		}
		if len(t.Exercise) != 0 {
			break
		}
	}

	DeviceId, exist := os.LookupEnv(DEVICE_ID_ENV)
	if !exist {
		log.Fatal("NOT FOUND DEVICE ID ENV")
	}
	t.DeviceId, _ = strconv.Atoi(DeviceId)

	return
}
