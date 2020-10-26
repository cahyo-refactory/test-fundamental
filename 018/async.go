package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

// func main() {
// 	// A slice of sample websites
// 	urls := []string{
// 		"https://www.easyjet.com/",
// 		"https://www.skyscanner.de/",
// 		"https://www.ryanair.com",
// 		"https://wizzair.com/",
// 		"https://www.swiss.com/",
// 	}
// 	for _, url := range urls {
// 		checkUrl(url)
// 	}
// }

// //checks and prints a message if a website is up or down
// func checkUrl(url string) {
// 	_, err := http.Get(url)
// 	if err != nil {
// 		fmt.Println(url, "is down !!!")
// 		return
// 	}
// 	fmt.Println(url, "is up and running.")
// }

func dataGet() {

	resp, err1 := http.Get("https://httpbin.org/get")
	if err1 != nil {
		fmt.Println(err1)
	}
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(body))
}

func dataPost() {
	reQuestBody, err1 := json.Marshal(map[string]string{
		"name":  "bada",
		"email": "bada@gmail.com",
	})
	if err1 != nil {
		fmt.Println(err1)
	}
	resp, err2 := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(reQuestBody))
	if err2 != nil {
		fmt.Println(err2)
	}

	body, err3 := ioutil.ReadAll(resp.Body)
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(string(body))
}

func dataPostFormData() {

	formData := url.Values{
		"name": {"lili"},
	}

	resp, err2 := http.PostForm("https://httpbin.org/post", formData)
	if err2 != nil {
		fmt.Println(err2)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	fmt.Println(result["form"])
}

func main() {
	dataGet()
	dataPost()
	dataPostFormData()
	// A slice of sample websites
	urls := []string{
		"https://www.sekolahdesain.id/",
		"https://soizee.com/",
		"https://refactory.id/",
		"https://enroll.refactory.id/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.google.co.id/",
		"https://www.bing.com/",
		"https://data.go.id/",
	}

	// for _, url := range urls {
	// 	// Call the function check
	// 	go checkUrl(url)
	// }
	// time.Sleep(1 * time.Second)

	var wg sync.WaitGroup

	for _, u := range urls {
		// Increment the wait group counter
		wg.Add(1)
		go func(url string) {
			// Decrement the counter when the go routine completes
			defer wg.Done()
			// Call the function check
			checkUrl(url)
		}(u)
	}
	// Wait for all the checkWebsite calls to finish
	wg.Wait()

}

//checks and prints a message if a website is up or down
func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down !!!")
		return
	}
	fmt.Println(url, "is up and running.")
}
