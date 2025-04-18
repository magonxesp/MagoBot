package booru

import (
	"testing"
)

func TestGetPostList(t *testing.T) {
	request := NewPostListRequest(Rule34, []string{"d.va"})
	posts, err := GetPostList(request)

	if err != nil {
		t.Error("failed fetching rule34 posts", err)
	}

	if posts == nil {
		t.Error("posts is nil")
	}

	t.Logf("result %+v\n", posts)
}
