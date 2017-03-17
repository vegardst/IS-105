package restful

type weatherData struct {
	Name string `json:"name"`
	Main map[string]int `json:"main"`
	Wind map[string]float64 `json:"wind"`
	Weather []status `json:"weather"`
}
type status struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

func GetWeather(city string) *weatherData {
	url := "http://api.openweathermap.org/data/2.5/weather?APPID=19026fe429b4298c6dad54698d4559b8&units=metric&q=" + city
	weather := new(weatherData)
	GetJSON(url, weather)
	return weather
}

// Query internal
//func Query(city string) WeatherData {
//	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=19026fe429b4298c6dad54698d4559b8&units=metric&q=" + city)
//	if err != nil {
//		return WeatherData{}
//	}
//	defer resp.Body.Close()
//	var data WeatherData
//	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
//		return WeatherData{}
//	}
//	return data
//
//}

