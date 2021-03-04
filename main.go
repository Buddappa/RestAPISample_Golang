package main

import (
	"fmt"
	"io/ioutil"
    "log"
    "net/http"
)

type server struct{}

func (s *server) ServeHTTP(w http.ResponseWriter, request  *http.Request) {
	startDate := request.URL.Query().Get("startDate")
	endDate := request.URL.Query().Get("endDate")
	w.Header().Set("Content-Type", "application/json")
	switch request.Method{
	case "GET":
		w.WriteHeader(http.StatusOK)
		response, err := http.Get("https://api.fda.gov/food/enforcement.json?search=report_date:["+startDate+"+TO+"+endDate+"]&limit=1")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
		  fmt.Println(string(data))
		  w.Write([]byte( data))
    }
	}
}

func main() {
    s := &server{}
    http.Handle("/GetEnforcementData", s)
    log.Fatal(http.ListenAndServe(":8080", nil))
}