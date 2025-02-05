package ui

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/thilob97/hubporter/internal/models"
)

type Tab int

const (
	TabWorkflows Tab = iota
	TabSettings
)

type Model struct {
	Table     table.Model
	Cursor    int
	Workflows []models.Workflow
	Err       error
	LastLog   string
	ActiveTab Tab
}

func NewModel() Model {
	return Model{
		Cursor:    0,
		Workflows: make([]models.Workflow, 0),
		ActiveTab: TabWorkflows,
	}
}
