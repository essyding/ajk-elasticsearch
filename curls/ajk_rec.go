package curls

import (
	"fmt"
	"log"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
)

const recTmpl = "https://shanghai.anjuke.com/v3/ajax/rec/profile/?cityid=%v&proids=%v&resulttype=3&page=1&pagesize=200"

// AjkRec simply call the HTTP GET and returns the parsed JSON body on success. Otherwise, return nil.
func AjkRec(cityId, propId int) *simplejson.Json {
	url := fmt.Sprintf(recTmpl, cityId, propId)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("cannot fetch url = %v, err = %v\n", url, err)
	}
	defer resp.Body.Close()

	body, err := simplejson.NewFromReader(resp.Body)
	if err != nil {
		log.Printf("cannot parse JSON object, err = %v\n", err)
		return nil
	}
	return body
}
