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

// A basic html template to display each page.
// Had to look up the gophercies video on this, have no experience with html templates!
// Extra comments included for future reference
var handlrTemplate = `
<!DOCTYPE html>
<html>
	<head>
	<meta charset="utf-8">
	<title>Story</title>
	</head>
	<style>
		li{
	  width: 350px;
		text align: center;
		}
		p {
			text align: center;
		}
	</style>
	<body bgcolor ="#ADD8E6">
		<h1>{{.Title}}</h1>
		{{range .Story}}
		<!-- Creates a paragraph with the story. Able to use just the dot since it was used earlier -->
		<p>{{.}}</p>
		{{end}}
		<ul>
			<!--range works as in go, in other words as many options will appear as exist -->
			{{range .Options}}
				<!--The respective text will link to the respective arc -->
				<li> <a href="/{{.Arc}}">{{.Text}}</a></li>
			{{end}}
		</ul>
	</body>
</html>
`
// Print the title of each story_arc followed by the actual story for that chapter
func printArcCmd(Title string,storyptr *Story){
	title := (*storyptr)[Title].Title
	story := (*storyptr)[Title].Story
	fmt.Println(title + ":")
	for i := 0; i < len(story); i++{
		fmt.Printf("%s\n",story[i])
	}
}

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
