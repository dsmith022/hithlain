package library

import "time"

/*
	Nodes are equivalent to a passage within a story.

Self explanitorily, each node will have 'root' which will be the
only node with no Parents and mark the starting point for the story.
'Body' denotes the unparsed raw text body entered via nano.
*/
type Node struct {
	ID        int64
	Title     string
	StoryID   int64
	Body      string
	ParentIDs []int
	ChildIDs  []int
	CreatedAt time.Time
}

type Story struct {
	ID        int64
	Title     string
	RootNode  *Node
	Nodes     []*Node
	CreatedAt time.Time
}
