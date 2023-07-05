package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://twitter.com/"+path, nil)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error creating request", http.StatusInternalServerError)
		return
	}
	req.Header.Set("User-Agent", "Googlebot")
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error making request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		http.Error(w, "Error reading response", http.StatusInternalServerError)
		return
	}

	w.Write(body)
}
