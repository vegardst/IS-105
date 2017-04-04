package main

//////////////////////////////////
// 	ICA05, IS-105	 	//
// 	2017, 	Zwirc		//
//////////////////////////////////
import (
	"./restful"
	"html/template"
	"net/http"
	"strings"
	"fmt"
	"path"
	"log"
	"net"
	"os"
	"time"
)

// GOOS=linux GOARCH=amd64 build -> Upload
// Template idea from http://www.alexedwards.net/blog/golang-response-snippets#html
func main() {
	fmt.Println("Starting server:")
	// Include css
	http.Handle("/web/", http.StripPrefix("/web/", http.FileServer(http.Dir("web"))))
	// Get that index
	http.HandleFunc("/", index)
	// ICON handler (So it does not run twice)
	http.HandleFunc("/favicon.ico", faviconHandler)
	// Listen on port
	err := http.ListenAndServe(":8080", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	// Get users IP
	ip, port, er := net.SplitHostPort(r.RemoteAddr)
	if er != nil {
		fmt.Print("Error")
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		fmt.Print("Could not fetch IP")
		return
	}
	fmt.Println("Visitor: " + ip + ":" + port)
	// For local IP testing:
	if len(ip) < 5 {
		ip = "158.37.63.125" // Server ip
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	city := "Kristiansand"
	if strings.Contains(r.URL.RawQuery, "city=") {
		city = strings.SplitN(r.URL.RawQuery, "city=", 2)[1]
	}
	data := rest.Loaddata(city, ip)

	// Save visits
	saveVisit(ip + " // " + data.UserData.Org + " // " + data.UserData.City + " // " + data.UserData.Country)

	// Generate page and include web folder and index.html
	index := path.Join("web", "index.html")
	tmpl, start := template.ParseFiles(index)

	// Start the server with loaded data
	if start := tmpl.Execute(w, data); start != nil {
		http.Error(w, start.Error(), http.StatusInternalServerError)
	}
	// Handle error
	err := r.ParseForm()
	if err != nil {
		http.Error(w, start.Error(), http.StatusInternalServerError)
	}
}
func saveVisit(data string) error {
	/*
	Open log.txt, and append user data with time of visit.
	Also format the text with spacings between each visit.
	*/
	t := time.Now()
	f, err := os.OpenFile("./log.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(t.String() + ":\n " + data + "\n \n")
	if err != nil {
		return err
	}
	return nil
}
func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "web/favicon.ico")
}
