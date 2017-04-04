package restful

//////////////////////////////////
// 	ICA05, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////

type UserData struct {
	As          string `json:"as"`
	City        string `json:"city"`
	Country     string `json:"country"`
	CountryCode string `json:"countryCode"`
	Isp         string `json:"isp"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
	Org         string `json:"org"`
	Query       string `json:"query"`
	Region      string `json:"region"`
	RegionName  string `json:"regionName"`
	Status      string `json:"status"`
	Timezone    string `json:"timezone"`
	Zip         string `json:"zip"`
}

func GetUserData(ip string) *UserData {
	url := "http://ip-api.com/json/" + ip
	userdata := new(UserData)
	GetJSON(url, userdata)
	return userdata
}
