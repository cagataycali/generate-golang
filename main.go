package main

import "fmt"

const SPREADSHEET_ID = "1x_W7Z2o_TGmEjL5cLTFbjO1R3KzQOqIhQKu9RQ4a_P4"
const API_KEY = "AIzaSyA5el9Fo8rMSYkcMjUqLfJi4tDB5_n0bzY"

func main() {
	url := fmt.Sprint("https://sheets.googleapis.com/v4/spreadsheets/", SPREADSHEET_ID, "/values:batchGet?key=", API_KEY, "&fields=valueRanges(range,values)&ranges=Mentees&ranges=Aktif%20Mentorluklar&ranges=Jobs&ranges=Interns")
	// we need more mappers for other pages
	// we have to call gh api as well, then map it.
	// Merge two contributions arrays within..
	_, err := MapPerson(url)
	if err != nil {
		panic(err)
	}
}
