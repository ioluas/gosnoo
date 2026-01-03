package ui

import "github.com/charmbracelet/lipgloss"

// Styles contains all the lipgloss styles used in the TUI.
var Styles = struct {
	Header   lipgloss.Style
	Title    lipgloss.Style
	Selected lipgloss.Style
	Cursor   lipgloss.Style
	Score    lipgloss.Style
	Author   lipgloss.Style
	Muted    lipgloss.Style
	Error    lipgloss.Style
	Help     lipgloss.Style
	Link     lipgloss.Style
}{
	Header: lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")).
		Background(lipgloss.Color("62")).
		Padding(0, 1),

	Title: lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("15")),

	Selected: lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("212")),

	Cursor: lipgloss.NewStyle().
		Foreground(lipgloss.Color("212")),

	Score: lipgloss.NewStyle().
		Foreground(lipgloss.Color("208")),

	Author: lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Bold(true),

	Muted: lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")),

	Error: lipgloss.NewStyle().
		Foreground(lipgloss.Color("196")).
		Bold(true),

	Help: lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")),

	Link: lipgloss.NewStyle().
		Foreground(lipgloss.Color("39")).
		Underline(true),
}
