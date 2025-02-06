package ui

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/table"
	"github.com/thilob97/hubporter/internal/models"
)

type Tab int

const (
	TabWorkflows Tab = iota
	TabSettings
)

type Model struct {
	Staging    table.Model
	Production table.Model
	Settings   list.Model
	Cursor     int
	Workflows  []models.Workflow
	Err        error
	LastLog    string
	ActiveTab  Tab
}

func NewModel() Model {
	return Model{
		Cursor:    0,
		Workflows: make([]models.Workflow, 0),
		ActiveTab: TabWorkflows,
	}
}

type SettingsItem struct {
	title       string
	description string
}

func (i SettingsItem) Title() string       { return i.title }
func (i SettingsItem) Description() string { return i.description }
func (i SettingsItem) FilterValue() string { return i.title }

// Constructor für ein neues SettingsItem
func NewSettingsItem(title, description string) SettingsItem {
	return SettingsItem{
		title:       title,
		description: description,
	}
}
