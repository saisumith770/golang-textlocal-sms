package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"net/http"
	"net/url"

	"github.com/joho/godotenv"
)

//the api endpoint for sending sms is same for all types of sms
//delay must be provided in seconds
func SendScheduledMessage(sender string, numbers []string, message string, delay int32) {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("ENV:", "failed to load environment variables")
	}

	var parsedDelimitedNumbers string = numbers[0]
	for i := 1; i < len(numbers); i++ {
		parsedDelimitedNumbers += "," + numbers[i]
	}

	params := url.Values{}
	params.Add("apiKey", os.Getenv("TEXT_LOCAL_API_KEY"))
	params.Add("sender", sender)
	params.Add("numbers", parsedDelimitedNumbers)
	params.Add("message", message)

	if delay != 0 {
		unix_time_stamp := time.Now().Add(time.Second * time.Duration(delay)).Unix()
		params.Add("schedule_time", fmt.Sprintf("%v", unix_time_stamp))
	}

	resp, err := http.PostForm("https://api.textlocal.in/send", params)
	if err != nil {
		log.Fatal("HTTPPOST:", err)
		return
	}

	var data any
	json.NewDecoder(resp.Body).Decode(&data)
	log.Print(data, resp.Status)
}

//calls the scheduled message func without delay
func SendOneToOneMessage(sender string, number string, message string) {
	SendScheduledMessage(sender, []string{number}, message, 0)
}

//calls the scheduled message func without delay
func SendBulkMessage(sender string, number []string, message string) {
	SendScheduledMessage(sender, number, message, 0)
}
