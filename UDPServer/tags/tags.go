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

var Exercises = [7]string{"squats", "push_ups", "chin_ups", "abs", "diamond_push_ups", "wide_push_ups", "clap_push_ups"}

func TagsIni() (t *Tags) {
	fmt.Println("starting of tags ini")
	msg := `
		Supporting exercises:
		1) squats (приседания)
		2) push_ups (отжимания)
		3) chin_ups (подтягивания)
		4) abs (пресс)
		5) diamond_push_ups (алмазные отжимания)
		6) wide_push_ups (широкие отжимания)
		7) clap_push_ups (отжимания с хлопками)
		8) test (тестовый режим)
		`
	t = &Tags{}

	exist := false
	fmt.Println("Insert your name: ")
	reader := bufio.NewReader(os.Stdin)
	name := ""
	fmt.Scanln(&name)
	repeats := ""
	for {
		fmt.Println("Insert your repeats(int): ")
		fmt.Scanln(&repeats)
		num, err := strconv.Atoi(repeats)
		if err != nil {
			fmt.Println("BAD INPUT PLEASE TRY AGAIN")
			continue
		}
		repeats = strconv.Itoa(num)
		break
	}

	t.User = name + "_" + time.Now().Format("2006-01-02T15:04:05") + "_" + repeats
	t.Repeats = repeats
	fmt.Println("Your name tag: ", t.User)
	fmt.Println(msg)
	for {
		fmt.Println("Insert number of exercise: ")
		num, _ := reader.ReadString('\n')
		fmt.Println(num)
		switch num[0] {
		case '1':
			t.Exercise = Exercises[0]
			break
		case '2':
			t.Exercise = Exercises[1]
			break
		case '3':
			t.Exercise = Exercises[2]
			break
		case '4':
			t.Exercise = Exercises[3]
			break
		case '5':
			t.Exercise = Exercises[4]
			break
		case '6':
			t.Exercise = Exercises[5]
			break
		case '7':
			t.Exercise = Exercises[6]
			break
		case '8':
			t.Exercise = "test"
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
