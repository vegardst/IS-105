package restful

//////////////////////////////////
// 	ICA05, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

import (
	"os"
	"encoding/json"
	"fmt"
)

type API struct {
	Google   string `json:"Google"`
	Weather  string `json:"Weather"`
	WeatherW string `json:"WeatherW"`
}

// Simply collecting stored API`s
func GetAPI() *API {
	file := "./json/API.json"
	data, _ := os.Open(file)
	decoder := json.NewDecoder(data)
	api := new(API)
	err := decoder.Decode(&api)
	if err != nil {
		fmt.Println("error:", err)
	}
	return api
}
