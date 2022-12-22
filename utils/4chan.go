package utils

import (
	"fmt"
	"github.com/moshee/go-4chan-api/api"
)

func ThreadUrl(thread *api.Thread) string {
	return fmt.Sprintf(
		"https://boards.4channel.org/%s/thread/%d",
		thread.Board,
		thread.Id(),
	)
}

func RandomThreadFromBoard(board string) (*api.Thread, error) {
	threads, err := api.GetThreads(board)

	if err != nil {
		return nil, err
	}

	page := threads[RandomInt(0, len(threads)-1)]
	threadId := page[RandomInt(0, len(page)-1)]

	thread, err := api.GetThread(board, threadId)

	if err != nil {
		return nil, err
	}

	return thread, nil
}

func RandomPostFromThread(thread *api.Thread) (*api.Post, error) {
	_, _, err := thread.Update()

	if err != nil {
		return nil, err
	}
	
	return thread.Posts[RandomInt(0, len(thread.Posts)-1)], nil
}
