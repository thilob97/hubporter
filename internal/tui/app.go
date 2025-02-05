package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/thilob97/hubporter/internal/api"
	"github.com/thilob97/hubporter/internal/models"
	"github.com/thilob97/hubporter/internal/ui"
)

// Message types
type (
	workflowsMsg []models.Workflow
	errMsg       error
)

// App represents the main TUI application
type App struct {
	model  ui.Model
	client *api.Client
}

// New creates a new TUI application
func New(client *api.Client) *App {
	columns := []table.Column{
		{Title: "ID", Width: 10},
		{Title: "Name", Width: 40},
		{Title: "Status", Width: 10},
		{Title: "Typ", Width: 20},
	}

	t := ui.TableStyle(columns)

	// Create model and set table
	model := ui.NewModel()
	model.Table = t

	return &App{
		model:  model,
		client: client,
	}
}

// Init implements tea.Model
func (a *App) Init() tea.Cmd {
	return func() tea.Msg {
		workflows, err := a.client.GetWorkflows()
		if err != nil {
			return errMsg(err)
		}
		return workflowsMsg(workflows)
	}
}

// updateRows updates the table rows with current workflow data
func (a *App) updateRows() {
	rows := make([]table.Row, len(a.model.Workflows))
	for i, w := range a.model.Workflows {
		status := "✓"
		if !w.Enabled {
			status = "✗"
		}
		rows[i] = table.Row{
			fmt.Sprintf("%d", w.ID),
			w.Name,
			status,
			w.Type,
		}
	}
	a.model.Table.SetRows(rows)
}

// Update implements tea.Model
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case workflowsMsg:
		a.model.Workflows = msg
		a.updateRows()

	case errMsg:
		a.model.Err = msg
		return a, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return a, tea.Quit

		case "up", "k":
			if a.model.Cursor > 0 {
				a.model.Cursor--
			}

		case "down", "j":
			if a.model.Cursor < len(a.model.Workflows)-1 {
				a.model.Cursor++
			}

		case "enter":
			if a.model.Cursor < len(a.model.Workflows) {
				a.model.LastLog = fmt.Sprintf("Selected workflow: %s",
					a.model.Workflows[a.model.Cursor].Name)
			}

		case "r":
			a.model.LastLog = fmt.Sprintf("Refreshing workflows...")
			return a, a.Init()
		case "tab", "1", "2":
			if a.model.ActiveTab == ui.TabWorkflows {
				a.model.ActiveTab = ui.TabSettings
			} else {
				a.model.ActiveTab = ui.TabWorkflows
			}
			a.model.LastLog = fmt.Sprintf("Switched to %s", a.getTabName())
		}
	}

	a.model.Table, cmd = a.model.Table.Update(msg)
	return a, cmd
}

func (a *App) getTabName() string {
	switch a.model.ActiveTab {
	case ui.TabWorkflows:
		return "Workflows"
	case ui.TabSettings:
		return "Settings"
	default:
		return "Unknown"
	}
}

// View implements tea.Model
func (a *App) View() string {
	if a.model.Err != nil {
		return fmt.Sprintf("Error: %v\n", a.model.Err)
	}

	var sb strings.Builder

	// Render tabs
	tabs := []string{"Workflows", "Settings"}
	for i, tab := range tabs {
		style := ui.TabStyle
		if a.model.ActiveTab == ui.Tab(i) {
			style = ui.ActiveTabStyle
		}
		sb.WriteString(style.Render(tab))
	}
	sb.WriteString("\n\n")

	// Render content based on active tab
	switch a.model.ActiveTab {
	case ui.TabWorkflows:
		sb.WriteString(ui.BaseStyle.Render(a.model.Table.View()))
	case ui.TabSettings:
		sb.WriteString("Settings view coming soon...")
	}

	// Help text
	sb.WriteString("\nPress q to quit • ↑/k and ↓/j to navigate • tab to switch views • enter to select\n")

	// Log line
	logText := "No action yet"
	if a.model.LastLog != "" {
		logText = a.model.LastLog
	}
	sb.WriteString(ui.LogStyle.Render(logText))
	sb.WriteString("\n")

	return sb.String()
}
