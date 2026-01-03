package app

import (
	"fmt"
	"strings"

	"github.com/ioluas/gosnoo/internal/state"
	"github.com/ioluas/gosnoo/internal/ui"
)

func (m Model) View() string {
	if m.err != nil {
		return ui.Styles.Error.Render(fmt.Sprintf("Error: %v\n\nPress 'r' to retry, 'q' to quit", m.err))
	}

	var content string
	switch m.screen {
	case state.ScreenPostList:
		content = m.viewPostList()
	case state.ScreenPostDetail:
		content = m.viewPostDetail()
	case state.ScreenComments:
		content = m.viewComments()
	}

	return content
}

func (m Model) viewPostList() string {
	var b strings.Builder

	header := ui.Styles.Header.Render(fmt.Sprintf(" r/%s ", m.subreddit))
	b.WriteString(header)
	b.WriteString("\n\n")

	if m.loading && len(m.posts) == 0 {
		b.WriteString(ui.Styles.Muted.Render("Loading..."))
		return b.String()
	}

	for i, post := range m.posts {
		cursor := "  "
		if i == m.cursor {
			cursor = ui.Styles.Cursor.Render("> ")
		}

		score := ui.Styles.Score.Render(fmt.Sprintf("[%d]", post.Score))
		title := post.Title
		if i == m.cursor {
			title = ui.Styles.Selected.Render(title)
		}
		comments := ui.Styles.Muted.Render(fmt.Sprintf("(%d comments)", post.NumComments))

		b.WriteString(fmt.Sprintf("%s%s %s %s\n", cursor, score, title, comments))
	}

	if m.loading {
		b.WriteString(ui.Styles.Muted.Render("\nLoading more..."))
	}

	b.WriteString("\n")
	b.WriteString(ui.Styles.Help.Render("↑/↓: navigate • enter: view • pgdn: more • r: refresh • q: quit"))

	return b.String()
}

func (m Model) viewPostDetail() string {
	if len(m.posts) == 0 {
		return "No post selected"
	}

	post := m.posts[m.cursor]
	var b strings.Builder

	b.WriteString(ui.Styles.Header.Render(fmt.Sprintf(" r/%s ", post.Subreddit)))
	b.WriteString("\n\n")

	b.WriteString(ui.Styles.Title.Render(post.Title))
	b.WriteString("\n")

	meta := fmt.Sprintf("by %s • %d points • %d comments", post.Author, post.Score, post.NumComments)
	b.WriteString(ui.Styles.Muted.Render(meta))
	b.WriteString("\n\n")

	if post.IsSelf && post.SelfText != "" {
		b.WriteString(post.SelfText)
		b.WriteString("\n")
	} else if post.URL != "" {
		b.WriteString(ui.Styles.Link.Render(post.URL))
		b.WriteString("\n")
	}

	b.WriteString("\n")
	b.WriteString(ui.Styles.Help.Render("c: comments • esc: back • q: quit"))

	return b.String()
}

func (m Model) viewComments() string {
	if len(m.posts) == 0 {
		return "No post selected"
	}

	post := m.posts[m.cursor]
	var b strings.Builder

	b.WriteString(ui.Styles.Header.Render(fmt.Sprintf(" Comments: %s ", truncate(post.Title, 40))))
	b.WriteString("\n\n")

	if m.loading {
		b.WriteString(ui.Styles.Muted.Render("Loading comments..."))
		return b.String()
	}

	if len(m.comments) == 0 {
		b.WriteString(ui.Styles.Muted.Render("No comments"))
		return b.String()
	}

	for _, comment := range m.comments {
		indent := strings.Repeat("  ", comment.Depth)
		author := ui.Styles.Author.Render(comment.Author)
		score := ui.Styles.Muted.Render(fmt.Sprintf("[%d]", comment.Score))

		b.WriteString(fmt.Sprintf("%s%s %s\n", indent, author, score))

		// Wrap comment body with indent
		body := wrapText(comment.Body, 80-len(indent))
		for _, line := range strings.Split(body, "\n") {
			b.WriteString(fmt.Sprintf("%s%s\n", indent, line))
		}
		b.WriteString("\n")
	}

	b.WriteString(ui.Styles.Help.Render("esc: back • q: quit"))

	return b.String()
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

func wrapText(s string, width int) string {
	if width <= 0 {
		width = 80
	}
	var result strings.Builder
	for _, line := range strings.Split(s, "\n") {
		if len(line) <= width {
			result.WriteString(line)
			result.WriteString("\n")
			continue
		}
		for len(line) > width {
			result.WriteString(line[:width])
			result.WriteString("\n")
			line = line[width:]
		}
		if len(line) > 0 {
			result.WriteString(line)
			result.WriteString("\n")
		}
	}
	return strings.TrimSuffix(result.String(), "\n")
}
