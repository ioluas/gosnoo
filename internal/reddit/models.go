package reddit

import "time"

// Post represents a Reddit post/submission.
type Post struct {
	ID          string
	Title       string
	Author      string
	Subreddit   string
	URL         string
	SelfText    string
	Score       int
	NumComments int
	CreatedUTC  time.Time
	IsNSFW      bool
	IsSelf      bool
	Thumbnail   string
}

// Comment represents a Reddit comment.
type Comment struct {
	ID         string
	Author     string
	Body       string
	Score      int
	CreatedUTC time.Time
	Depth      int
}

// Subreddit represents a subreddit.
type Subreddit struct {
	Name        string
	Title       string
	Description string
	Subscribers int
	IsNSFW      bool
}
