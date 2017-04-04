package restful

//////////////////////////////////
// 	ICA05, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

import (
	"fmt"
	"encoding/json"

	"strings"
)

type Population struct {
	TotalPopulation []struct {
		Date       string `json:"date"`
		Population int `json:"population"`
	} `json:"total_population"`
}

func GetPopulation(country string) string {
	url := "http://api.population.io/1.0/population/" + country + "/today-and-tomorrow/?format=json"
	population := new(Population)
	slice := GetJSONByte(url)
	json.Unmarshal([]byte(slice), &population)

	result := fmt.Sprintf("%v", population.TotalPopulation)
	result = strings.Replace(result, "} {", " -> ", -1)
	result = strings.Replace(result, "[{", " ", -1)
	result = strings.Replace(result, "}]", " ", -1)
	return result

}
