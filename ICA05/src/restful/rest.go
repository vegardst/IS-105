package restful

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"time"
)
var httpClient = &http.Client{Timeout: 10 * time.Second}
func RunAll(){

}

// Return JSON as byteslice
func GetJSONByte(url string) []byte{
	res, err := httpClient.Get(url)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}
	return body
}

// Return JSON + Interface
func GetJSON(url string, target interface{}) error {
	r, err := httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}