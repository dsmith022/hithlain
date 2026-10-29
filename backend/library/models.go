package library

import "time"

/*
	Nodes are equivalent to a passage within a story.

Self explanitorily, each node will have 'root' which will be the
only node with no Parents and mark the starting point for the story.
'Body' denotes the unparsed raw text body entered via nano.
*/
type Node struct {
	ID        int       `json:"id"`
	StoryID   int       `json:"story_id"`
	Body      string    `json:"body"`
	Parents   []Node    `json:"parents"`
	Children  []Node    `json:"children"`
	CreatedAt time.Time `json:"created_at"`
}

type Story struct {
	ID        int       `json:"id"`
	RootNode  Node      `json:"first_node"`
	Nodes     []Node    `json:"nodes"`
	CreatedAt time.Time `json:"created_at"`
}

// Librarian is the account-linked object which acts as the topmost point of a savefile's filetree.
type Librarian struct {
	Library  []Story `json:"library"`
	FilePath string  `json:"-"`
}
