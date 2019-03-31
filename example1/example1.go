package example1

import (
	"encoding/json"
	"fmt"
	m "go-concurrency-examples/models"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://github.com",
	"https://gitlab.com",
	"https://oreilly.com",
	"https://cloud.google.com",
	"https://aws.amazon.com",
	"https://heroku.com",
	"https://digitalocean.com",
	"https://linkedin.com",
	"https://facebook.com",
	"https://facebookewew.com",
	"https://twitter.com",
}

var data []m.Web

func fetchURL(url string, wg *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
		wg.Done()
		webResp := m.Web{URL: url, Status: err.Error()}
		data = append(data, webResp)
		return "", nil
	}
	wg.Done()
	fmt.Println(resp.Status)
	webResp := m.Web{URL: url, Status: resp.Status}
	data = append(data, webResp)
	return resp.Status, nil
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit")
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetchURL(url, &wg)
	}
	wg.Wait()
	fmt.Println("Returning Response")
	w.Header().Set("Content-Type", "application/json")
	resData := m.Resp{Status: "success", Data: data}
	json.NewEncoder(w).Encode(resData)
	data = nil
}

func HandleRequests() {
	log.Println("Application started")
	http.HandleFunc("/", getStatus)
	http.ListenAndServe(":8090", nil)
}
