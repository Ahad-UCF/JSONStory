// This file holds all the information about the overall story structures
// and related functions
package main

import(
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
func printArc(Title string,storyptr *Story){
	title := (*storyptr)[Title].Title
	story := (*storyptr)[Title].Story
	fmt.Println(title)
	for i := 0; i < len(story); i++{
		fmt.Println(story[i])
	}
}
