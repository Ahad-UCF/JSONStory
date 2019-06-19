// This file simmply runs the main function of the program
package main

import ()

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


}
