package models

// TODO: Define the Workflow and Action structs
type Workflow struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	ID         int    `json:"id"`
	InsertedAt int64  `json:"insertedAt"`
	UpdatedAt  int64  `json:"updatedAt"`
	Enabled    bool   `json:"enabled"`
}

type Action struct {
	Properties map[string]interface{} `json:"properties"`
	Type       string                 `json:"type"`
}
