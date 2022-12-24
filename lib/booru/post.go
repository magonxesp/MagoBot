package booru

import (
	"fmt"
	"github.com/MagonxESP/MagoBot/utils"
	"net/http"
	"strconv"
	"strings"
)

type Post struct {
	PreviewUrl   string `json:"preview_url"`
	SampleUrl    string `json:"sample_url"`
	FileUrl      string `json:"file_url"`
	Directory    int64  `json:"directory"`
	Hash         string `json:"hash"`
	Height       int    `json:"height"`
	Id           int64  `json:"id"`
	Image        string `json:"image"`
	Change       int    `json:"change"`
	Owner        string `json:"owner"`
	ParentId     int64  `json:"parentId"`
	Rating       string `json:"rating"`
	Sample       int    `json:"sample"`
	SampleHeight int    `json:"sample_height"`
	SampleWidth  int    `json:"sample_width"`
	Score        int    `json:"score"`
	Tags         string `json:"tags"`
	Width        int    `json:"width"`
}

type PostListRequest struct {
	// Limit how many posts you want to retrieve. There is a default limit of 100 posts per request.
	Limit int
	// The Page number
	Page int
	// The Tags to search for. Any tag combination that works on the website will work here.
	Tags []string
	// Cid is the Change ID of the post. This is in Unix time so there are likely others with the same value if updated at the same time.
	Cid int
	// The post Id.
	Id int
	// Set to true for Json formatted response.
	Json bool
	// The Booru on will perform the request
	Booru string
}

func NewPostListRequest(booru string, tags []string) *PostListRequest {
	return &PostListRequest{
		Limit: 100,
		Page:  1,
		Tags:  tags,
		Json:  true,
		Booru: booru,
	}
}

func (p *PostListRequest) ToQueryString() string {
	params := map[string]string{}

	if p.Limit != 0 {
		params["limit"] = strconv.Itoa(p.Limit)
	}

	if p.Page != 0 {
		params["pid"] = strconv.Itoa(p.Page)
	}

	if len(p.Tags) > 0 {
		params["tags"] = strings.Join(p.Tags, " ")
	}

	if p.Cid != 0 {
		params["cid"] = strconv.Itoa(p.Cid)
	}

	if p.Id != 0 {
		params["cid"] = strconv.Itoa(p.Id)
	}

	if p.Json != false {
		params["json"] = "1"
	}

	return strings.Join(utils.MapToKeyValueList(params, "="), "&")
}

func GetPostList(request *PostListRequest) ([]Post, error) {
	url := fmt.Sprintf(
		"%s/index.php?page=dapi&s=post&q=index&%s",
		request.Booru,
		request.ToQueryString(),
	)

	response, err := http.Get(url)

	if response.StatusCode != 200 {
		return nil, NewPostRequestError(fmt.Sprintf("An error ocurred fetching the post list from %s", url))
	}

	var posts []Post
	err = UnmarshalResponseBody(response.Body, &posts)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
