package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

var (
	date       = "2022-07-13"
	time       = "18:30"
	p    int64 = 4
	// 台中新馬辣
	company  = "-LjL6vW09dVGOC0tGamg"
	branchID = "-Mw-S9FSZby-GN_0tiZe"
	// 台中中山燒肉
	// company  = -LzoSPyWXzTNSaE - I4QJ
	// branchID = -MdytTkuohNf5wnBz1vZ
)

func main() {
	client := &http.Client{}
	url := "https://inline.app/api/booking-capacitiesV3?companyId=" + company + "%3Ainline-live-1&branchId=" + branchID
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	val := gjson.Get(string(body), "default."+date+".times."+time).Array()
	for _, v := range val {
		if v.Int() == p {
			url := "https://inline.app/api/reservations/booking"
			jsonData := []byte(`{"language":"zh-tw","company":"` + company + `:inline-live-1","branch":"` + branchID + `","groupSize":"` + strconv.FormatInt(p, 10) + `","kids":0,"gender":0,"purposes":[],"email":"","name":"李亞諦","phone":"+886937550247","note":"","date":"` + date + `","time":"` + time + `","numberOfKidChairs":0,"numberOfKidSets":0,"skipPhoneValidation":false,"referer":"www.google.com"}`)
			req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
			req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36")
			req.Header.Set("Content-Type", "application/json")
			if err != nil {
				log.Fatal(err)
			}

			resp, err := client.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatal(err)
			}
			log.Println(string(body))
		}
	}
}
