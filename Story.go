package main

import (
	"fmt"
	"os"
	"encoding/json"
)


// Our struct is stored in a map
// This specific line had to be taken from the gophercises video on the question
// A map is used because each chapter name varies and is a string, thus it can be used as an index for a map
type Story map[string] Story_arc


// This struct holds each story arc and its elements
type Story_arc struct {
	Title string `json:"title"`
	Story []string `json:"story"`
	Options []Option `json:"options"`
}

// Each option has 2 elements
type Option struct {
	Text string `json:"text"`
	Arc string `json:"arc"`
}

func main(){

	// Generate a map that we will later dump the json into
	var story Story

	// Grab the json file and make sure it exists
	jsonFile, err := os.Open("gopher.json")
	checkErr(err)

	// Decode the json file and dump it into the map
	decoder := json.NewDecoder(jsonFile)
	err = decoder.Decode(&story)

	// Verify the decode resolved
	checkErr(err)

	// TODO: Remove this line, for testing purposes only
	fmt.Printf("%+v\n",story)
}


// Verify that an error did not occur, if it did... panic!
func checkErr(err error){
	if err !=nil{
		panic(err)
	}
}
