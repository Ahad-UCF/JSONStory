// This file holds all the information about the overall story structures
// and related functions
package main

import(
	"os"
	"fmt"
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

// Print the title of each story_arc followed by the actual story for that chapter
func printArcCmd(Title string,storyptr *Story){
	title := (*storyptr)[Title].Title
	story := (*storyptr)[Title].Story
	fmt.Println(title + ":")
	for i := 0; i < len(story); i++{
		fmt.Printf("%s\n",story[i])
	}
}

// TODO: Develop webapp version

// The below functions are all for the prelimenary command line story

// Print the options that each story arc has
func printOptionsCmd(Title string,storyptr *Story){
	options := (*storyptr)[Title].Options

	// Terminate the program when we reach the part of the json with no options
	if len(options) == 0{
		os.Exit(0)
	}
	fmt.Println("Choices")

	// Cycle through and the print the Text of each option
	for i:=0; i < len(options); i++{
		fmt.Printf("Choice %d:	%s\n", i,options[i].Text)
	}
}

// Grab which option the user wants to follow
func getOptionCmd(Title string,storyptr *Story)(int){
	choice := -1
	fmt.Scan(&choice)
	return choice
}

// Reads the entire story
// Calls each function until the end case is detected by the print options command
func readStoryCmd(Title string,storyptr *Story){
	title := Title
	for{
		options := (*storyptr)[title].Options
		printArcCmd(title, storyptr)
		printOptionsCmd(title, storyptr)
		title = options[getOptionCmd(title, storyptr)].Arc
		fmt.Println(title)
	}
}
