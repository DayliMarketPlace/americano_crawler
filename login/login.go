package main

import (
	"encoding/json"
	"log"

	"github.com/gocolly/colly"
)

//unmarshal need exported expression
type RealResponse struct {
	Data PaidData `json:"data"`
}

//separated dataset
type PaidData []struct {
	ID           string `json:"id"`
	UserID       string `json:"partnerUserId"`
	CardNumber   string `json:"partnerUserCardNumber"`
	Name         string `json:"partnerUserName"`
	PaidCoin     string `json:"paidCoin"`
	PaidItem     string `json:"paidItem"`
	ItemQuantity string `json:"paidItemQuantity"`
	PaidStatus   string `json:"paidStatus"`
	PaidTime     string `json:"paidTimestampDateTime"`
}

func main() {
	// create a new collector
	c := colly.NewCollector()

	// authenticate
	err := c.Post("http://partner.yellocoin.com/sign/ajax/doCheckPartner", map[string]string{"email": "dmp@daylifg.com", "password": "daily@broccoli!2"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)

		if r.Request.URL.Path == "/orders/data/paidListData" {
			msg := &RealResponse{}
			//msgs := []interface{}
			err = json.Unmarshal(r.Body, msg)
			if err != nil {
				log.Println("response raw : ", r.Body)
				log.Fatal(err.Error())
				return
			}
			log.Println("response data : ", msg)
		}
	})

	result := c.Post("http://partner.yellocoin.com/orders/data/paidListData", map[string]string{
		"sEcho": "9", "iColumns": "6", "sColumns": " ,,,,,", "iDisplayStart": "0", "iDisplayLength": "25",
		"mDataProp_0": "paidDate", "mDataProp_1": "partnerGroupCorpName", "mDataProp_2": "partnerUserName", "mDataProp_3": "paidItem", "mDataProp_4": "function", "mDataProp_5": "function",
		"iSortCol_0": "0", "sSortDir_0": "asc",
		"paidStartDate": "2019/01/01", "paidEndDate": "2019/01/08", "userName": "장현민"})

	if result != nil {
		log.Println("result : ", result)
		log.Fatal(result)
	}

	c.Visit("http://partner.yellocoin.com/")
}
