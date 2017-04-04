package rest

//////////////////////////////////
// 	ICA05, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

import (
	"./data"
	"fmt"
	"time"
)

type Data struct {
	Weather     *restful.WeatherData
	WeatherW    *restful.Wunderground
	Wiki        string
	Countrycode string
	Country     string
	Population  string
	UserData    *restful.UserData
	VisitorIP   string
	API         *restful.API
	Temperature float64
}

// Laster all JSON data inn i en struct som brukes i html siden
func Loaddata(city string, ip string) Data {
	var data Data
	data.Weather = restful.GetWeather(city)
	if data.Weather.Name == "" {
		data.Weather = restful.GetWeather("Asdas")
	}
	// Set country code
	data.Countrycode = data.Weather.Sys.Country

	//For the sake of concurrency
	var load bool = true
	loaded := 0
	itemstoload := 6

	// GoRoutines
	go func() {
		data.API = restful.GetAPI()
		loaded++
	}()

	go func() {
		// Not realy in use, since it get multiple results for each location, so need a way to handle that.
		// To be added //
		data.WeatherW = restful.GetWunderground(city, data.Countrycode)
		loaded++
	}()

	go func() {
		data.VisitorIP = ip
		loaded++
	}()

	go func() {
		data.Country = restful.GetCountry(data.Countrycode)
		data.Population = restful.GetPopulation(data.Country)
		loaded++
	}()
	go func() {
		data.Wiki = restful.GetWiki(city)
		loaded++
	}()

	go func() {
		data.UserData = restful.GetUserData(ip)
		loaded++
	}()

	for load {
		fmt.Println("Loading....")
		time.Sleep(1000 * time.Millisecond)
		if loaded == itemstoload {
			load = false
		}
	}

	return data

}
