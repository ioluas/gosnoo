package reddit

import (
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

// mapPost converts a go-reddit Post to our app-level Post.
func mapPost(p *reddit.Post) Post {
	return Post{
		ID:          p.ID,
		Title:       p.Title,
		Author:      p.Author,
		Subreddit:   p.SubredditName,
		URL:         p.URL,
		SelfText:    p.Body,
		Score:       p.Score,
		NumComments: p.NumberOfComments,
		CreatedUTC:  p.Created.UTC(),
		IsNSFW:      p.NSFW,
		IsSelf:      p.IsSelfPost,
	}
}

// mapPosts converts a slice of go-reddit Posts.
func mapPosts(posts []*reddit.Post) []Post {
	result := make([]Post, len(posts))
	for i, p := range posts {
		result[i] = mapPost(p)
	}
	return result
}

// mapComment converts a go-reddit Comment to our app-level Comment.
func mapComment(c *reddit.Comment, depth int) Comment {
	return Comment{
		ID:         c.ID,
		Author:     c.Author,
		Body:       c.Body,
		Score:      c.Score,
		CreatedUTC: c.Created.UTC(),
		Depth:      depth,
	}
}

// flattenComments recursively flattens a comment tree with depth information.
func flattenComments(comments []*reddit.Comment, depth int) []Comment {
	var result []Comment
	for _, c := range comments {
		result = append(result, mapComment(c, depth))
		if len(c.Replies.Comments) > 0 {
			result = append(result, flattenComments(c.Replies.Comments, depth+1)...)
		}
	}
	return result
}
