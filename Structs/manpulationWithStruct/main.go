package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var Body []byte

type user struct {
	UserId    int    `json;"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {

	url := "http://jsonplaceholder.typicode.com/todos"

	resp, err := http.Get(url)

	if err != nil {
		log.Printf("Error while send the request %v", err)
	}

	if resp.StatusCode != 200 {
		log.Fatalf("Failed to fetch data. HTTP Status Code: %d", resp.StatusCode)
	}

	Body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error while reading the response %v", err)
		return
	}
	var userId int
	fmt.Print("Enter userIdP
	|: ")

	_, err = fmt.Scan(&userId)
	if err != nil {
		fmt.Println("Error: Invalid input")
		return
	}

	result := workPercentage(Body, userId)

	fmt.Printf("Percentage of completion of task by UserId 1 is %.2f%%\n", result)
}

func workPercentage(body []byte, userId int) float64 {

	var users []user

	err := json.Unmarshal(body, &users)
	if err != nil {
		log.Printf("Error while unmarsheling the data %v", err)
	}
	var totalTask int = 0
	for _, user := range users {
		if user.UserId == userId {
			totalTask++
		}
	}
	fmt.Println(totalTask)

	var completedTask int = 0
	for _, user := range users {
		if user.UserId == userId && user.Completed == true {
			completedTask++
		}
	}
	fmt.Println(completedTask)

	completedWorkPercentage := (float64(completedTask) / float64(totalTask)) * 100

	return completedWorkPercentage

}
