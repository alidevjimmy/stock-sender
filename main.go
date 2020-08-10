package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	CONTENT_TYPE    string = "application/json"
	ACCEPT          string = "application/json"
	REFERER         string = "https://d.easytrader.emofid.com/"
	CONNECTION      string = "keep-alive"
	ACCEPT_ENCODING string = "gzip, deflate, br"
	USER_AGENT      string = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36"
)

var (
	authorization string = "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImI3MmYyMjczZTE4YTQ0YjQ5OTFmMDg3ODIzNzQyYmI1IiwidHlwIjoiYXQrand0In0.eyJuYmYiOjE1OTcwMzcwNjksImV4cCI6MTU5NzA0Nzg3OSwiaXNzIjoiaHR0cHM6Ly9hY2NvdW50LmVtb2ZpZC5jb20iLCJhdWQiOlsiZWFzeTJfYXBpIiwiaHR0cHM6Ly9hY2NvdW50LmVtb2ZpZC5jb20vcmVzb3VyY2VzIl0sImNsaWVudF9pZCI6ImVhc3kyX2NsaWVudF9wa2NlIiwic3ViIjoiNGE3NTFkZDctYjgzYS00NDg5LTk0NWMtODgwNjMxMzlmNjFjIiwiYXV0aF90aW1lIjoxNTk2OTUyNDk1LCJpZHAiOiJsb2NhbCIsInBrIjoiNGE3NTFkZDctYjgzYS00NDg5LTk0NWMtODgwNjMxMzlmNjFjIiwidHdvX2ZhY3Rvcl9lbmFibGVkIjoiZmFsc2UiLCJwcmVmZXJyZWRfdXNlcm5hbWUiOiI0YTc1MWRkNy1iODNhLTQ0ODktOTQ1Yy04ODA2MzEzOWY2MWMiLCJuYW1lIjoiNGE3NTFkZDctYjgzYS00NDg5LTk0NWMtODgwNjMxMzlmNjFjIiwicGhvbmVfbnVtYmVyIjoiMDkzNzI1MTY2MDgiLCJwaG9uZV9udW1iZXJfdmVyaWZpZWQiOnRydWUsIm5hdGlvbmFsX2lkIjoiMTgxMDYyMDY5NCIsIm5hdGlvbmFsX2lkX3ZlcmlmaWVkIjoidHJ1ZSIsImN1c3RvbWVyX2lzaW4iOiIxMTI5MTgxMDYyMDY5NCIsInNjb3BlIjpbIm9wZW5pZCIsImVhc3kyX2FwaSJdLCJhbXIiOlsicHdkIl19.XoMeozhikJKz8zFsxSWuSLMbAU5iHc4XnBQAt1r669LlA3sSLsg98z-2W5YIYEagkFjpJgmIbWHKXsMo5ZYv0o2Pzl5YgPeW1zAhSMBqTTMmi_5tPNJA5Xg-udrG1JDNgLagXNlWm3RoIlyBlh-YBwaf6X0azB6u05bXnSBu7GZpRd2JQx7_ykHelTHeoyGDNUSflyuDNSLhvg8cwbPJObHxgG40OREPl2kibF_JjfV08BOo8iUtT8kTs26l3291ioL8vaxp6K-QEMVvGMJ5CKvVKq19XZBwse_k_jl1YeqDgmfqEHk1SJ0lwTnRBMcLcYYwwA-eBEFyxqBVSXLSvA"
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

func sendRequest(i *int) {
	fmt.Println("Sending Request ", *i, "...")
	*i++
	// time.Sleep(time.Second * 1)
	url := "https://d11.emofid.com/easy/api/OmsOrder"
	_ = url
	data := ReqBody{
		CautionAgreementSelected: false,
		EasySource:               1,
		FinanceID:                1,
		Isin:                     "IRO1SITA0001", // sita
		Price:                    24710,          // sita
		Quantity:                 250,            // sita
		ReferenceKey:             "89c047ee-a1d2-4f58-8714-9c5601b35166",
		Side:                     0,
		ValidityDateJalali:       "1399/5/17",
		ValidityType:             74,
	}
	jsonData, _ := json.Marshal(&data)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", CONTENT_TYPE)
	req.Header.Set("Accept", ACCEPT)
	req.Header.Set("Authorization", authorization)
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

	// fmt.Printf("Status : %v\n", resp.Status)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Printf("Body : %v\n", string(body))
	// wg.Done()
}

func main() {
	// var wg sync.WaitGroup
	i := 1
	for {
		// wg.Add(1)
		// time.Sleep(time.Microsecond * 150000)
		sendRequest(&i)

	}

	// wg.Wait()
}
