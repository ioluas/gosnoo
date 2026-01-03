package reddit

import (
	"context"

	"github.com/vartanbeno/go-reddit/v2/reddit"
)

// Service wraps the go-reddit client to provide a clean interface for the TUI.
type Service struct {
	client *reddit.Client
}

// NewReadonlyClient creates a read-only Reddit client (no auth required).
func NewReadonlyClient() (*reddit.Client, error) {
	return reddit.NewReadonlyClient()
}

// NewService creates a new Reddit service wrapper.
func NewService(client *reddit.Client) *Service {
	return &Service{client: client}
}

// Posts fetches posts from a subreddit.
// sort can be: "hot", "new", "top", "rising"
// Returns posts, the "after" cursor for pagination, and any error.
func (s *Service) Posts(subreddit, sort, after string, limit int) ([]Post, string, error) {
	ctx := context.Background()
	opts := &reddit.ListOptions{
		Limit: limit,
		After: after,
	}

	var posts []*reddit.Post
	var resp *reddit.Response
	var err error

	switch sort {
	case "new":
		posts, resp, err = s.client.Subreddit.NewPosts(ctx, subreddit, opts)
	case "top":
		posts, resp, err = s.client.Subreddit.TopPosts(ctx, subreddit, &reddit.ListPostOptions{
			ListOptions: *opts,
			Time:        "day",
		})
	case "rising":
		posts, resp, err = s.client.Subreddit.RisingPosts(ctx, subreddit, opts)
	default: // "hot"
		posts, resp, err = s.client.Subreddit.HotPosts(ctx, subreddit, opts)
	}

	if err != nil {
		return nil, "", err
	}

	afterCursor := ""
	if resp != nil {
		afterCursor = resp.After
	}

	return mapPosts(posts), afterCursor, nil
}

// Comments fetches comments for a post.
func (s *Service) Comments(subreddit, postID string) ([]Comment, error) {
	ctx := context.Background()

	// Fetch post with comments
	postAndComments, _, err := s.client.Post.Get(ctx, postID)
	if err != nil {
		return nil, err
	}

	return flattenComments(postAndComments.Comments, 0), nil
}

// Subreddit fetches subreddit information.
func (s *Service) Subreddit(name string) (*Subreddit, error) {
	ctx := context.Background()

	sr, _, err := s.client.Subreddit.Get(ctx, name)
	if err != nil {
		return nil, err
	}

	return &Subreddit{
		Name:        sr.Name,
		Title:       sr.Title,
		Description: sr.Description,
		Subscribers: sr.Subscribers,
		IsNSFW:      sr.NSFW,
	}, nil
}
