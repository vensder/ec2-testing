package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

var timeout = time.Duration(1000 * time.Millisecond)

var meta_data_items = []string{
	"ami-id",
	"hostname",
	"instance-id",
	"instance-type",
	"local-hostname",
	"local-ipv4",
	"public-ipv4",
	"public-keys",
	"security-groups",
}

var meta_items_struct = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

func hostname() string {
	hostname, err := os.Hostname()
	if err == nil {
		return hostname
	}
	return "localhost"
}

func getMetaData(meta_data_item string) string {
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get("http://169.254.169.254/latest/meta-data/" + meta_data_item + "/")
	if err != nil {
		fmt.Println("Can't get aws meta data item: " + meta_data_item)
		return "can't get " + meta_data_item
	}
	defer resp.Body.Close() // Close body only if response non-nil
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println("Can't get responce Body for meta data item: " + meta_data_item)
			return "can't get " + meta_data_item
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return "Status is Not OK for " + meta_data_item
}

func putMetaData(meta_data_item string) {
	item := getMetaData(meta_data_item)
	meta_items_struct.Lock()
	meta_items_struct.m[meta_data_item] = item
	meta_items_struct.Unlock()
}

func main() {

	for _, item := range meta_data_items {
		go putMetaData(item)
	}

	fmt.Println("Server starting...")
	hostname := hostname()
	env_color := os.Getenv("color")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		file_content, err := ioutil.ReadFile("date.txt")
		if err != nil {
			fmt.Println(err)
		}

		fmt.Fprintf(w, ("<p>URL Path: " + r.URL.Path + "</p>"))
		fmt.Fprintf(w, ("<p>date.txt: " + string(file_content) + "</p>"))
		fmt.Fprintf(w, ("<h1 style=\"background-color:" + env_color + ";\">Host: " + hostname + "</h1>"))
		for _, item := range meta_data_items {
			meta_items_struct.RLock()
			fmt.Fprintf(w, ("<h2>" + item + ": " + meta_items_struct.m[item] + "</h2>"))
			meta_items_struct.RUnlock()
		}

		for _, e := range os.Environ() {
			fmt.Fprintf(w, ("<p>" + string(e) + "</p>"))
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
