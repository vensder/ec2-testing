package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func hostname() string {
	hostname, err := os.Hostname()
	if err == nil {
		return hostname
	}
	return "localhost"
}

func getResponceBody(url string) string {
	timeout := time.Duration(1 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println("Can't get URL: " + url)
		return "can't get URL" + url
	}
	defer resp.Body.Close() // Close body only if response non-nil
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			fmt.Println("Can't get responce Body for URL: " + url)
			return "can't get  responce Body for URL: " + url
		}
		bodyString := string(bodyBytes)
		return bodyString
	}
	return "Status is Not OK for URL: " + url
}

// func collectMetaURLs(start_url string) []string {
// 	timeout := time.Duration(1 * time.Second)
// 	client := http.Client{
// 		Timeout: timeout,
// 	}
// 	resp, err := client.Get(start_url + "/")
// }

func getMetaData(meta_data_item string) string {
	return getResponceBody("http://169.254.169.254/latest/meta-data/" + meta_data_item + "/")
}

func main() {
	meta_data_items := []string{
		"ami-id",
		"hostname",
		"instance-id",
		"instance-type",
		"local-hostname",
		"local-ipv4",
		"public-ipv4",
		"public-keys",
		"security-groups",
		"services",
		"error-handling-test",
	}

	meta_data_items_map := make(map[string]string)
	for _, item := range meta_data_items {
		meta_data_items_map[item] = getMetaData(item)
		fmt.Println(item + ": " + meta_data_items_map[item])
	}
	fmt.Println("Server starting...")
	hostname := hostname()
	fmt.Println("Hostname: " + hostname)
	env_color := os.Getenv("color")
	fmt.Println(env_color)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		file_content, err := ioutil.ReadFile("date.txt")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(file_content))

		fmt.Fprintf(w, ("<p>URL Path: " + r.URL.Path + "</p>"))
		fmt.Fprintf(w, ("<p>date.txt: " + string(file_content) + "</p>"))
		fmt.Fprintf(w, ("<h1 style=\"background-color:" + env_color + ";\">Host: " + hostname + "</h1>"))
		for _, item := range meta_data_items {
			fmt.Fprintf(w, ("<h2>" + item + ": " + meta_data_items_map[item] + "</h2>"))
		}

		for _, e := range os.Environ() {
			fmt.Println(string(e))
			fmt.Fprintf(w, ("<p><small>" + string(e) + "</small></p>"))
		}

		fmt.Println(hostname)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
