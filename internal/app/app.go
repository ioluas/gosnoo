package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/ioluas/gosnoo/internal/reddit"
)

// Run starts the TUI application.
func Run(svc *reddit.Service) error {
	p := tea.NewProgram(newModel(svc), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
