package library

import "time"

type Node struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	Parents   []Node    `json:"parents"`
	Children  []Node    `json:"children"`
	CreatedAt time.Time `json:"created_at"`
}

type Story struct {
	ID        int       `json:"id"`
	Nodes     []Node    `json:"nodes"`
	CreatedAt time.Time `json:"created_at"`
}

type Librarian struct {
	Library  []Story `json:"library"`
	FilePath string  `json:"-"`
}
