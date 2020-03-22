package batch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/config"
	"github.com/Pickausernaame/sporttech_udp_server/UDPServer/models"
	"log"
	"math/rand"
	"net/http"
)

const (
	TOKEN = "179673ec2a5a8b34ef8ff14ec5f6b5064ec1f295"
)

type Batch struct {
	DataArray []models.Data `json:"batch"`
}

func InitBatch(capacity int) (d *Batch) {
	d = &Batch{
		DataArray: make([]models.Data, capacity),
	}
	return d
}

func (d *Batch) mock() {
	for i := 0; i < len(d.DataArray); i++ {
		d.DataArray[i].Fields.AccX = json.Number(fmt.Sprintf("%f", rand.Float64()))
		d.DataArray[i].Fields.AccY = json.Number(fmt.Sprintf("%f", rand.Float64()))
		d.DataArray[i].Fields.AccZ = json.Number(fmt.Sprintf("%f", rand.Float64()))
		d.DataArray[i].Fields.GyroX = json.Number(fmt.Sprintf("%f", rand.Float64()))
		d.DataArray[i].Fields.GyroY = json.Number(fmt.Sprintf("%f", rand.Float64()))
		d.DataArray[i].Fields.GyroZ = json.Number(fmt.Sprintf("%f", rand.Float64()))
	}
}

func (d *Batch) clear() {
	d.DataArray = nil
}

func (d *Batch) Send(conf *config.Config) {
	client := http.Client{}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(&d)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := json.Marshal(d)
	//fmt.Println(string(b))
	req, err := http.NewRequest("POST", conf.URL, bytes.NewBuffer(b))
	if err != nil {
		fmt.Println(conf.URL)
		log.Fatal("SENDING BATCH ERROR ", err)
	}
	fmt.Println("SENDING BATCH")
	req.Header.Set("Authorization", "Token "+TOKEN)
	req.Header.Add("Content-Type", "application/json")
	_, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(resp)
	d.clear()
	d.DataArray = make([]models.Data, conf.BATCH_CAPACITY)
	return
}
