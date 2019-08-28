package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

func DoGET(url, headers string) (string, int) {
	// 	resp, err := http.Get(url, headers)
	// 	if err != nil {
	// 		log.Fatal("Failed GET", err)
	// 	}
	// 	defer resp.Body.Close()
	// 	body, err := ioutil.ReadAll(resp.Body)
	// 	if err != nil {
	// 		log.Fatal("Failed GET", err)
	// 	}
	// 	return body, resp.StatusCode

	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	// req.Header.Set(headers)
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed GET", err)
	}
	return string(body), resp.StatusCode
}

func DoPOST(url, headers, data string) {

}

func DoPUT(url, headers, data string) {

}

func DoPATCH(url, headers, data string) {

}
func DoDELETE(url, headers string) {

}
