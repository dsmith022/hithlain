-- +goose Up
CREATE TABLE nodes(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    story_id uuid NOT NULL,
    body TEXT NOT NULL,
    parent_nodes
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE stories(
    id UUID PRIMARY KEY default gen_random_uuid(),
    title TEXT NOT NULL,
    root_node_id INTEGER 
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
);

CREATE TABLE node_edges(
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    parent_node UUID NOT NULL references nodes(id) ON DELETE CASCADE,
    child_node UUID NOT NULL references nodes(id) ON DELETE CASCADE,
);

-- +goose Down
DROP TABLE IF EXISTS nodes;
DROP TABLE IF EXISTS node_edges;
DROP TABLE IF EXISTS stories;

/*FINISH SETTING UP GOOSE WHEN YOU RESUME*/