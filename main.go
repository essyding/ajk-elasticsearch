package main

import (
	"log"
	"strconv"
	"time"

	"github.com/essyding/ajk-elasticsearch/curls"
)

const maxChanLen = 100000

var (
	cityID     = 11
	initPropID = 1204922769
)

func main() {
	c := make(chan int, maxChanLen)
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	c <- initPropID
	for propID := range c {
		log.Printf("fetch %v started.", propID)
		go func(prodID int) {
			rets := curls.AjkRec(cityID, propID).Get("results")
			n := len(rets.MustArray())
			for i := 0; i < n; i++ {
				p := rets.GetIndex(i)
				pid, err := p.Get("PROID").Int()
				if err != nil {
					log.Printf("cannot get PROID: %v\n", err)
					continue
				}
				proprice, err := p.Get("PROPRICE").String()
				if err != nil {
					log.Printf("cannot get PROPRICE: %v", err)
					continue
				}
				// potential property to consider. add to the queue.
				price, err := strconv.Atoi(proprice)
				if err == nil && price > 650 && price < 810 {
					c <- pid
				}
				// form the payload and add to the database.
				payload, err := p.MarshalJSON()
				if err != nil {
					log.Printf("cannot marshal json: %v\n", err)
					continue
				}
				if !curls.AjkEsPut(pid, payload) {
					log.Printf("failed to put %v to elastic search.", pid)
				}
			}
		}(propID)
		<-ticker.C
	}
}
