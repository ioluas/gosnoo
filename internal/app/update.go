package app

import (
	"github.com/charmbracelet/bubbletea"
	"github.com/ioluas/gosnoo/internal/reddit"
	"github.com/ioluas/gosnoo/internal/state"
	"github.com/ioluas/gosnoo/internal/ui"
)

// Message types for async operations.
type postsLoadedMsg struct {
	posts []reddit.Post
	page  PageInfo
	err   error
}

type commentsLoadedMsg struct {
	comments []reddit.Comment
	err      error
}

func (m Model) Init() tea.Cmd {
	return m.fetchPostsCmd("hot")
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case postsLoadedMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		m.posts = msg.posts
		m.page = msg.page
		m.cursor = 0
		return m, nil

	case commentsLoadedMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
			return m, nil
		}
		m.comments = msg.comments
		return m, nil
	}

	return m, nil
}

func (m Model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.String()

	// Global keys
	if ui.Keys.Quit.Contains(key) {
		return m, tea.Quit
	}

	switch m.screen {
	case state.ScreenPostList:
		return m.handlePostListKeys(key)
	case state.ScreenPostDetail:
		return m.handlePostDetailKeys(key)
	case state.ScreenComments:
		return m.handleCommentsKeys(key)
	}

	return m, nil
}

func (m Model) handlePostListKeys(key string) (tea.Model, tea.Cmd) {
	switch {
	case ui.Keys.Up.Contains(key):
		if m.cursor > 0 {
			m.cursor--
		}
	case ui.Keys.Down.Contains(key):
		if m.cursor < len(m.posts)-1 {
			m.cursor++
		}
	case ui.Keys.Enter.Contains(key):
		if len(m.posts) > 0 {
			m.screen = state.ScreenPostDetail
		}
	case ui.Keys.NextPage.Contains(key):
		if m.page.After != "" && !m.loading {
			m.loading = true
			return m, m.fetchMorePostsCmd("hot")
		}
	case ui.Keys.Refresh.Contains(key):
		m.loading = true
		m.page.After = ""
		return m, m.fetchPostsCmd("hot")
	}
	return m, nil
}

func (m Model) handlePostDetailKeys(key string) (tea.Model, tea.Cmd) {
	switch {
	case ui.Keys.Back.Contains(key):
		m.screen = state.ScreenPostList
	case ui.Keys.Comments.Contains(key):
		if len(m.posts) > 0 {
			m.screen = state.ScreenComments
			m.loading = true
			return m, m.fetchCommentsCmd(m.posts[m.cursor].ID)
		}
	}
	return m, nil
}

func (m Model) handleCommentsKeys(key string) (tea.Model, tea.Cmd) {
	switch {
	case ui.Keys.Back.Contains(key):
		m.screen = state.ScreenPostDetail
		m.comments = nil
	}
	return m, nil
}

// Commands for async fetching.
func (m Model) fetchPostsCmd(sort string) tea.Cmd {
	return func() tea.Msg {
		posts, after, err := m.reddit.Posts(m.subreddit, sort, "", 25)
		return postsLoadedMsg{
			posts: posts,
			page:  PageInfo{After: after},
			err:   err,
		}
	}
}

func (m Model) fetchMorePostsCmd(sort string) tea.Cmd {
	return func() tea.Msg {
		posts, after, err := m.reddit.Posts(m.subreddit, sort, m.page.After, 25)
		return postsLoadedMsg{
			posts: append(m.posts, posts...),
			page:  PageInfo{After: after},
			err:   err,
		}
	}
}

func (m Model) fetchCommentsCmd(postID string) tea.Cmd {
	return func() tea.Msg {
		comments, err := m.reddit.Comments(m.subreddit, postID)
		return commentsLoadedMsg{
			comments: comments,
			err:      err,
		}
	}
}
