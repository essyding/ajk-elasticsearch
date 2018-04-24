package curls

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

type Payload struct {
	PROID     int    `json:"PROID"`
	CITYID    int    `json:"CITYID"`
	TITLE     string `json:"TITLE"`
	LINK      string `json:"LINK"`
	PROPRICE  string `json:"PROPRICE"`
	ROOMNUM   int    `json:"ROOMNUM"`
	HALLNUM   int    `json:"HALLNUM"`
	TOILETNUM int    `json:"TOILETNUM"`
	AREANUM   string `json:"AREANUM"`
	COMMID    int    `json:"COMMID"`
	COMMNAME  string `json:"COMMNAME"`
	AREACODE  string `json:"AREACODE"`
	BROKERID  int    `json:"BROKERID"`
	IMAGESRC  string `json:"IMAGESRC"`
	Pricing   int    `json:"Pricing"`
	SOJ       string `json:"SOJ"`
	PROPTYPE  int    `json:"PROPTYPE"`
}

func AjkEsPut(propId int, payload []byte) bool {
	body := bytes.NewReader(payload)

	url := fmt.Sprintf("http://localhost:9200/ajk/prop/%v?pretty", propId)

	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		log.Printf("error creating new PUT request: %v\n", err)
		return false
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("charset", "utf-8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("cannot put data in es: %v\n", err)
		return false
	}
	defer resp.Body.Close()
	return true
}
