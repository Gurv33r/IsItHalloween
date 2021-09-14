package server

type Assignment struct {
	ID     int32  `json:"id"`
	Title  string `json:"title"`
	Prompt string `json:"prompt"`
}

var Assignments = []Assignment{
	{1, "Implement Stack", "Implement a stack data structure in a programming language of your choice"},
	{2, "Implement Queue", "Implement a queue data structure in a programming language of your choice"},
	{3, "Implement Tree", "Implement a tree data structure in a programming language of your choice"},
}
