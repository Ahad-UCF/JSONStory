// This file simmply runs the main function of the program
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	// Generate a map that we will later dump the json into
	var story Story
	// Generate a bool value to determine whether we are using command line to print
	CMD := grabCMD()
	decodeJson(grabJson(), &story)

	if (CMD){
		// Command line story
		readStoryCmd("intro", &story)
	}
	// Create a handler with our Story variable
	h := newHandler(story)
	// For testing purposes, check if an http page on port 3000 works
	fmt.Println("STARTING SERVER ON 3000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", 3000),h))
}
