package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

var app = flag.String("app", "app","name of demo app")
var project = flag.String("project","demo project", "sample project to set a project tag")
var posts = flag.Int("posts",1, "how many posts")
var entries = flag.Int("entries", 1, "number of entries per post")
var host = flag.String("host", "localhost","host")

// Simple script to add some data to the frisk database for demo purposes
func main() {
	flag.Parse()
	friskUrl := "http://" + *host + "/v1/frisk/record"
	log.Println("Starting to put data in to frisk for app: ", *app)
	for postCount := 1; postCount <= *posts; postCount++ {
		var postRecords []map[string]interface{}
		for entryCount := 1; entryCount <= *entries; entryCount++ {
			// Randomly select pass or fail
			result := "pass"
			r := rand.Intn(2)
			log.Println("RESULT: ", r)
			if r == 1 {
				result = "fail"
			}
			e := map[string]interface{}{
				"title": "test-" + strconv.Itoa(postCount) + "-" + strconv.Itoa(entryCount),
				"tags": map[string]string{
					"app": *app,
					"project": *project,
					"result": result,
				},
			}
			postRecords = append(postRecords, e)
		}
		// Send payload to api
		err := postFrisk(friskUrl,postRecords)
		if err != nil {
			log.Println("PostErr: ", err)
		}
	}
}

func postFrisk(friskUrl string, records []map[string]interface{}) error {
	d,err := json.Marshal(records)
	if err != nil {
		return err
	}
	client := http.Client{}
	req, err := http.NewRequest("POST", friskUrl, bytes.NewBuffer(d))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	log.Println("Code: ", resp.StatusCode)
	return nil
}

