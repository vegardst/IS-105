package restful

//////////////////////////////////
// 	ICA05, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

import (
	"fmt"
	"encoding/json"
	"reflect"
)

type Data struct {
	Countries []struct {
		Name          string `json:"name"`
		Alpha2        string `json:"alpha-2"`
		Alpha3        string `json:"alpha-3"`
		CountryCode   string `json:"country-code"`
		Iso31662      string `json:"iso_3166-2"`
		Region        string `json:"region"`
		SubRegion     string `json:"sub-region"`
		RegionCode    string `json:"region-code"`
		SubRegionCode string `json:"sub-region-code"`
	} `json:"countries"`
}

func GetCountry(code string) string {
	file := "./json/country.json"
	data := new(Data)
	bytes := GetJSONFileByte(file)
	json.Unmarshal(bytes, &data)
	string := "Unknown"
	for _, value := range data.Countries {
		s := reflect.ValueOf(&value).Elem()
		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			if f.Interface() == code {
				string = fmt.Sprintf("%v", s.FieldByName("Name"))
			}
		}
	}
	return string
}