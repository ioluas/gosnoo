package app

import (
	"github.com/ioluas/gosnoo/internal/reddit"
	"github.com/ioluas/gosnoo/internal/state"
)

// Model represents the application state.
type Model struct {
	screen    state.Screen
	subreddit string
	posts     []reddit.Post
	comments  []reddit.Comment
	cursor    int
	page      PageInfo
	loading   bool
	err       error
	reddit    *reddit.Service
	width     int
	height    int
}

// PageInfo holds pagination state.
type PageInfo struct {
	After string
}

func newModel(svc *reddit.Service) Model {
	return Model{
		screen:    state.ScreenPostList,
		subreddit: "reddevils",
		reddit:    svc,
	}
}
