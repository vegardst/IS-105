package restful

//////////////////////////////////
// 	ICA05, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)

var httpClient = &http.Client{Timeout: 10 * time.Second}

// JSON from url to byteslice
func GetJSONByte(url string) []byte {
	r, err := httpClient.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	return body
}

// JSON url to interface
func GetJSON(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}

// JSON from file to byteslice
func GetJSONFileByte(file string) []byte {
	r, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err.Error())
	}

	return r
}
