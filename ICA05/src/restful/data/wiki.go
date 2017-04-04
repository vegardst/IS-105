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

func GetWiki(city string) string {
	url := "http://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&titles=" + city
	slice := GetJSONByte(url)
	data := make(map[string]interface{})
	json.Unmarshal(slice, &data)
	////// Handle missing Wiki

	error := func() string {
		return "Could not load Wiki for " + city + " \n ,There may be a error with the search!"
	}

	string := fmt.Sprintf("%v", data)
	if strings.Contains(string, "missing:") {
		error()
	} else if strings.Contains(string, "invalidreason:") {
		error()
	} else {
		string := fmt.Sprintf("%v", data)
		p := strings.SplitAfterN(string, "extract:", 2)[1]
		s := strings.SplitAfterN(p, "pageid", 5)[0]
		string = strings.Replace(s, "]]]]", " ", -1)
		return string
	}
	return error()
}
