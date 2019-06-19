// This utilities file stores functions that will be used throughout the program
package main

import (
	"os"
	"flag"
	"encoding/json"
)

// TODO: Create a flag function to grab whether a command line or webapp story will be used
func grabCMD() (bool) {
	CMD := flag.Bool("Command_Line", true, "true or false, will command line be used to tell the story?")
	flag.Parse()
	return *CMD
}

// function to grab and return a jsonFile to decode later
func grabJson() (*os.File){
	fileName := flag.String("json_File", "gopher.json", "The json file's name")
	flag.Parse()
	// Grab the json file and make sure it exists
	jsonFile, err := os.Open(*fileName)
	checkErr(err)

	return jsonFile
}

// Decodes a json and edits the story pointer to fill it with the decoded info
func decodeJson(jsonFile *os.File, storyptr *Story){

	// Decode the json file and dump it into the map
	decoder := json.NewDecoder(jsonFile)
	err := decoder.Decode(storyptr)

	// Verify the decode resolved
	checkErr(err)
}



// Verify that an error did not occur, if it did... panic!
func checkErr(err error){
	if err !=nil{
		panic(err)
	}
}
