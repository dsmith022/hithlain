-- +goose Up
CREATE TABLE nodes(
    id INTEGER PRIMARY KEY ,
    title TEXT NOT NULL,
    body TEXT NOT NULL
);

CREATE TABLE stories(
    id INTEGER PRIMARY KEY,
    title TEXT NOT NULL,
    root_node_id INTEGER NOT NULL references nodes(id) ON DELETE CASCADE
);

CREATE TABLE node_edges(
    id INTEGER PRIMARY KEY,
    parent_node INTEGER NOT NULL references nodes(id) ON DELETE CASCADE,
    child_node INTEGER NOT NULL references nodes(id) ON DELETE CASCADE
);

CREATE TABLE story_nodes(
    story_id INTEGER NOT NULL references stories(id) ON DELETE CASCADE,
    node_id INTEGER NOT NULL references nodes(id) ON DELETE CASCADE,
    PRIMARY KEY(story_id, node_id)
);

-- +goose Down
DROP TABLE IF EXISTS nodes;
DROP TABLE IF EXISTS node_edges;
DROP TABLE IF EXISTS stories;
DROP TABLE IF EXISTS story_id;

