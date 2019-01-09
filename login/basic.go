package main

import (
	"encoding/json"
	"log"

	"github.com/gocolly/colly"
)

func getData(c *colly.Collector) {

	result := c.Post("http://partner.yellocoin.com/orders/data/paidListData", map[string]string{
		"sEcho": "9", "iColumns": "6", "sColumns": " ,,,,,", "iDisplayStart": "0", "iDisplayLength": "25",
		"mDataProp_0": "paidDate", "mDataProp_1": "partnerGroupCorpName", "mDataProp_2": "partnerUserName", "mDataProp_3": "paidItem", "mDataProp_4": "function", "mDataProp_5": "function",
		"iSortCol_0": "0", "sSortDir_0": "asc",
		"paidStartDate": "2019/01/01", "paidEndDate": "2019/01/08", "userName": "장현민"})

	if result != nil {
		log.Println("result : ", result)
		log.Fatal(result)
	}

	c.OnResponse(func(response *colly.Response) {
		log.Println("response status code : ", response.StatusCode)

		temp := []byte(`{"data":[{"id":"880370","partnerUserId":"1560","partnerBranchId":"6","partnerGroupCorpId":"37","partnerGroupCorpName":"\ub370\uc77c\ub9ac\ub9c8\ucf13\ud50c\ub808\uc774\uc2a4","partnerTid":"201901080156","partnerRid":"","partnerUserName":"\uc7a5\ud604\ubbfc","partnerUserCardNumber":"1016347","partnerBranchName":"\uc610\ub85c\uae08\uc735\uadf8\ub8f9","paidCoin":"2500","paidCoinNumberFormat":"2,500","paidDate":"20190108","paidItem":"\ube14\ub8e8\ubca0\ub9ac\uc5d0\uc774\ub4dc","paidItemQuantity":"1","paidIp":"210.220.77.22","paidStatus":"confirmed","paidTimestamp":"1546929509","paidTimestampDate":"2019-01-08","paidTimestampDateTime":"2019-01-08 15:38:29","regTimestamp":"1546929514","DT_RowId":"1560"},{"id":"879883","partnerUserId":"1560","partnerBranchId":"6","partnerGroupCorpId":"37","partnerGroupCorpName":"\ub370\uc77c\ub9ac\ub9c8\ucf13\ud50c\ub808\uc774\uc2a4","partnerTid":"201901080059","partnerRid":"","partnerUserName":"\uc7a5\ud604\ubbfc","partnerUserCardNumber":"1016347","partnerBranchName":"\uc610\ub85c\uae08\uc735\uadf8\ub8f9","paidCoin":"1500","paidCoinNumberFormat":"1,500","paidDate":"20190108","paidItem":"ICE \uc544\uba54\ub9ac\uce74\ub178","paidItemQuantity":"1","paidIp":"210.220.77.22","paidStatus":"confirmed","paidTimestamp":"1546909228","paidTimestampDate":"2019-01-08","paidTimestampDateTime":"2019-01-08 10:00:28","regTimestamp":"1546909244","DT_RowId":"1560"},{"id":"879029","partnerUserId":"1560","partnerBranchId":"6","partnerGroupCorpId":"37","partnerGroupCorpName":"\ub370\uc77c\ub9ac\ub9c8\ucf13\ud50c\ub808\uc774\uc2a4","partnerTid":"201901070064","partnerRid":"","partnerUserName":"\uc7a5\ud604\ubbfc","partnerUserCardNumber":"1016347","partnerBranchName":"\uc610\ub85c\uae08\uc735\uadf8\ub8f9","paidCoin":"2500","paidCoinNumberFormat":"2,500","paidDate":"20190107","paidItem":"\ubaa8\uacfc\uc5d0\uc774\ub4dc","paidItemQuantity":"1","paidIp":"210.220.77.22","paidStatus":"confirmed","paidTimestamp":"1546823509","paidTimestampDate":"2019-01-07","paidTimestampDateTime":"2019-01-07 10:11:49","regTimestamp":"1546823514","DT_RowId":"1560"},{"id":"878402","partnerUserId":"1560","partnerBranchId":"6","partnerGroupCorpId":"37","partnerGroupCorpName":"\ub370\uc77c\ub9ac\ub9c8\ucf13\ud50c\ub808\uc774\uc2a4","partnerTid":"201901040073","partnerRid":"","partnerUserName":"\uc7a5\ud604\ubbfc","partnerUserCardNumber":"1016347","partnerBranchName":"\uc610\ub85c\uae08\uc735\uadf8\ub8f9","paidCoin":"2500","paidCoinNumberFormat":"2,500","paidDate":"20190104","paidItem":"\ubaa8\uacfc\uc5d0\uc774\ub4dc","paidItemQuantity":"1","paidIp":"210.220.77.22","paidStatus":"confirmed","paidTimestamp":"1546565264","paidTimestampDate":"2019-01-04","paidTimestampDateTime":"2019-01-04 10:27:44","regTimestamp":"1546565277","DT_RowId":"1560"},{"id":"876855","partnerUserId":"1560","partnerBranchId":"6","partnerGroupCorpId":"37","partnerGroupCorpName":"\ub370\uc77c\ub9ac\ub9c8\ucf13\ud50c\ub808\uc774\uc2a4","partnerTid":"201901020096","partnerRid":"","partnerUserName":"\uc7a5\ud604\ubbfc","partnerUserCardNumber":"1016347","partnerBranchName":"\uc610\ub85c\uae08\uc735\uadf8\ub8f9","paidCoin":"1800","paidCoinNumberFormat":"1,800","paidDate":"20190102","paidItem":"ICE \ubaa8\uacfc\ucc28","paidItemQuantity":"1","paidIp":"210.220.77.22","paidStatus":"confirmed","paidTimestamp":"1546395673","paidTimestampDate":"2019-01-02","paidTimestampDateTime":"2019-01-02 11:21:13","regTimestamp":"1546395677","DT_RowId":"1560"}],"recordsSum":"10800","recordsTotal":"5","recordsFiltered":"5"}`)

		msg := &RealResponse{}
		err := json.Unmarshal(temp, msg)
		if err != nil {
			log.Println("temp data : ", temp)
			log.Fatal(err.Error())
			return
		}
		log.Println("temp data : ", msg)

		//msgs := []interface{}
		err = json.Unmarshal(response.Body, msg)
		if err != nil {
			log.Println("response raw : ", response.Body)
			log.Fatal(err.Error())
			return
		}
		log.Println("response data : ", msg)
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("http://partner.yellocoin.com/")
}
