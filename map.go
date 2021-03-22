package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type GoogleSheetResponse struct {
	ValueRanges []struct {
		Range  string     `json:"range"`
		Values [][]string `json:"values"`
	} `json:"valueRanges"`
}

type Person struct {
	RegisteredAt  string `json:"registeredAt"`
	Name          string `json:"name"`
	TwitterHandle string `json:"TwitterHandle"`
	GitHub        string `json:"GitHub"`
	LinkedIn      string `json:"LinkedIn"`
	Interests     string `json:"Interests"`
	Goals         string `json:"Goals"`
	Mentor        string `json:"Mentor"`
	Stackoverflow string `json:"Stackoverflow"`
}

func Parser(person []string) (p Person) {
	p.RegisteredAt = person[0]
	p.Name = person[1]
	p.TwitterHandle = person[2]
	p.GitHub = person[3]
	p.LinkedIn = person[4]
	p.Interests = person[5]
	p.Goals = person[6]
	p.Mentor = person[7]
	// Sorry for shitty google spreadsheet api, I'll be fix when I join google as developer.
	if len(person) > 8 {
		p.Stackoverflow = person[8]
	}
	return p
}

func GetData(url string) (gs GoogleSheetResponse, err error) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(body, &gs)
	return
}

func MapPerson(url string) ([]Person, error) {
	result, err := GetData(url)
	if err != nil {
		return nil, err
	}
	mentees := result.ValueRanges[0].Values[7:]
	menteesArr := make([]Person, len(mentees))
	for i, person := range mentees {
		p := Parser(person)
		fmt.Println(p)
		menteesArr[i] = p
	}

	return menteesArr, nil
}
