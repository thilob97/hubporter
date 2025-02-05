package ui

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	BaseStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("240"))

	LogStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("yellow"))
)
