package server

//assignment struct
type Assignment struct {
	ID     int32  `json:"id"`
	Title  string `json:"title"`
	Class  string `json:"class"`
	Prompt string `json:"prompt"`
	Link   string `json:"link"`
}

// seed data as an Assignment slice
var Assignments = []Assignment{
	{ID: 1, Title: "Stack Implementation", Class: "Data Structures and Algorithms", Prompt: "Implement a stack data structure in a coding language of your choice", Link: "www.google.com"},
	{ID: 2, Title: "Queue Implementation", Class: "Data Structures and Algorithms", Prompt: "Implement a queue data structure in a coding language of your choice", Link: "www.google.com"},
	{ID: 3, Title: "Tree Implementation", Class: "Data Structures and Algorithms", Prompt: "Implement a tree data structure in a coding language of your choice", Link: "www.google.com"},
}
