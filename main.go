package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const (
	CONTENT_TYPE    string = "application/json"
	ACCEPT          string = "application/json"
	REFERER         string = "https://d.easytrader.emofid.com/"
	CONNECTION      string = "keep-alive"
	ACCEPT_ENCODING string = "gzip, deflate, br"
	USER_AGENT      string = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36"
)

type ReqBody struct {
	CautionAgreementSelected bool   `json:"cautionAgreementSelected`
	EasySource               int    `json:"easySource`
	FinanceID                int    `json:"financeId`
	Isin                     string `json:"isin`
	Price                    int    `json:"price`
	Quantity                 int    `json:"quantity`
	ReferenceKey             string `json:"referenceKey`
	Side                     int    `json:"side`
	ValidityDateJalali       string `json:"validityDateJalali`
	ValidityType             int    `json:"validityType`
}

func sendRequest(i *int, auth, id string, quantity, price int) {
	fmt.Println("Sending Request ", *i, "...")
	*i++
	// time.Sleep(time.Second * 1)
	url := "https://d11.emofid.com/easy/api/OmsOrder"
	_ = url
	data := ReqBody{
		CautionAgreementSelected: false,
		EasySource:               1,
		FinanceID:                1,
		Isin:                     id,
		Price:                    price,
		Quantity:                 quantity,
		// ReferenceKey:             "cbca6842-a7b7-41cf-8b3d-a64ea9ca917b",
		ReferenceKey:       "89c047ee-a1d2-4f58-8714-9c5601b35166",
		Side:               0,
		ValidityDateJalali: "1399/5/17",
		ValidityType:       74,
	}
	jsonData, _ := json.Marshal(&data)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", CONTENT_TYPE)
	req.Header.Set("Accept", ACCEPT)
	req.Header.Set("Authorization", auth)
	req.Header.Set("Referer", REFERER)
	req.Header.Set("User-Agent", USER_AGENT)
	req.Header.Set("Connection", CONNECTION)
	req.Header.Set("Accept-Encoding", ACCEPT_ENCODING)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := make([]string, 4)
	var ii uint8 = 0
	for scanner.Scan() {
		data[ii] = scanner.Text()
		ii++
	}

	price, _ := strconv.ParseInt(data[2], 10, 32)
	quantity, _ := strconv.ParseInt(data[3], 10, 32)
	i := 1
	for {
		sendRequest(&i, data[0], data[1], int(quantity), int(price))
	}
}
