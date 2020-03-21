package models

import (
	"encoding/json"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/tags"
	"time"
)

type Packet struct {
	AccX  json.Number `json:"accX"`
	AccY  json.Number `json:"accY"`
	AccZ  json.Number `json:"accZ"`
	GyroX json.Number `json:"gyroX"`
	GyroY json.Number `json:"gyroY"`
	GyroZ json.Number `json:"gyroZ"`
}

type Data struct {
	Measurement string    `json:"measurement"`
	Tags        tags.Tags `json:"tags"`
	Time        time.Time `json:"time"`
	Fields      Packet    `json:"fields"`
}
