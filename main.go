package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Chanlenger aproaching!")

	body, error := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if error != nil {
		println("Errror")
	}

	fmt.Println(string(body))
	go writeData(string(body))
}

func writeData(data string) {
	os.WriteFile("test.txt", []byte(data), 0644)
}
