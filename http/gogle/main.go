package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    // Define the URL you want to get
    url := "https://www.google.com"

    // Make the GET request
    response, err := http.Get(url)
    if err != nil {
        log.Fatalf("Error fetching URL: %v", err)
    }
    defer response.Body.Close()

    // Read the response body
    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalf("Error reading response body: %v", err)
    }

    // Print the response body
    fmt.Println(string(body))
}
