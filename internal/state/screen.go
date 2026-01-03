package state

// Screen represents the current screen/view in the TUI.
type Screen int

const (
	ScreenPostList Screen = iota
	ScreenPostDetail
	ScreenComments
	ScreenSubredditList
	ScreenHelp
)

// String returns the string representation of the screen.
func (s Screen) String() string {
	switch s {
	case ScreenPostList:
		return "posts"
	case ScreenPostDetail:
		return "post"
	case ScreenComments:
		return "comments"
	case ScreenSubredditList:
		return "subreddits"
	case ScreenHelp:
		return "help"
	default:
		return "unknown"
	}
}
