package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/thilob97/hubporter/internal/models"
)

type Model struct {
	Table     table.Model
	Cursor    int
	Workflows []models.Workflow
	Err       error
	LastLog   string
}

func NewModel() Model {
	return Model{
		Cursor:    0,
		Workflows: make([]models.Workflow, 0),
	}
}
