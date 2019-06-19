// This file simmply runs the main function of the program
package main

import ()

func main(){
	// Generate a map that we will later dump the json into
	var story Story
	decodeJson(grabJson(), &story)
	// Command line story
	readStoryCmd("intro", &story)
}
