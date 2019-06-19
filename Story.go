package main

import (
	"fmt"
	"os"
	"encoding/json"
	"flag"
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
	decodeJson(grabJson(), &story)
	printArc("intro", &story)
}

func grabJson() (*os.File){
	fileName := flag.String("json_File", "gopher.json", "The json file's name")
	// Grab the json file and make sure it exists
	jsonFile, err := os.Open(*fileName)
	checkErr(err)

	return jsonFile
}

func decodeJson(jsonFile *os.File, storyptr *Story){

	// Decode the json file and dump it into the map
	decoder := json.NewDecoder(jsonFile)
	err := decoder.Decode(storyptr)

	// Verify the decode resolved
	checkErr(err)

	/* TODO: Remove this line, for testing purposes only
	fmt.Println((*storyptr)["intro"].Title)
	fmt.Println((*storyptr)["intro"].Story)
	fmt.Println((*storyptr)["intro"].Options[0].Text)
	fmt.Println((*storyptr)["intro"].Options[1].Text)*/
}


// Print the title of each story_arc followed by the actual story for that chapter
func printArc(Title string,storyptr *Story){
	fmt.Println((*storyptr)[Title].Title)
	for i := 0; i < len((*storyptr)[Title].Story); i++{
		fmt.Println((*storyptr)[Title].Story[i])
	}
}

// Verify that an error did not occur, if it did... panic!
func checkErr(err error){
	if err !=nil{
		panic(err)
	}
}
