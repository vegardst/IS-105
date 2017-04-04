package restful

//////////////////////////////////
// 	ICA05, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

type WeatherData struct {
	Name    string `json:"name"`
	Main    map[string]float64 `json:"main"`
	Wind    map[string]float64 `json:"wind"`
	Weather []status `json:"weather"`
	Sys     sys `json:"sys"`
}
type status struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}
type sys struct {
	Country string `json:"country"`
	Sunrise int `json:"sunrise"`
	Sunset  int `json:"sunset"`
}

func GetWeather(city string) *WeatherData {
	api := GetAPI().Weather
	url := "http://api.openweathermap.org/data/2.5/weather?APPID=" + api  + "&units=metric&q=" + city
	weather := new(WeatherData)
	GetJSON(url, weather)
	return weather
}
