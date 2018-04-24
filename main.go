package main

import (
	"log"

	"github.com/essyding/ajk-elasticsearch/curls"
)

var (
	cityID = 11
	propID = 1199795382
)

func main() {
	rets := curls.AjkRec(cityID, propID).Get("results")
	n := len(rets.MustArray())
	for i := 0; i < n; i++ {
		p := rets.GetIndex(i)
		pid, err := p.Get("PROID").Int()
		if err != nil {
			log.Printf("cannot get PROID: %v\n", err)
			continue
		}
		payload, err := p.MarshalJSON()
		if err != nil {
			log.Printf("cannot marshal json: %v\n", err)
			continue
		}
		if !curls.AjkEsPut(pid, payload) {
			log.Printf("failed to put %v to elastic search.", pid)
		}
	}
}
