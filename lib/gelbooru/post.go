package gelbooru

import (
	"github.com/MagonxESP/MagoBot/utils"
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
}

func NewPostListRequest(tags []string) *PostListRequest {
	return &PostListRequest{
		Limit: 100,
		Page:  1,
		Tags:  tags,
		Json:  true,
	}
}

func (p *PostListRequest) ToQueryString() string {
	params := map[string]string{}

	if p.Limit != 0 {
		params["limit"] = strconv.Itoa(p.Limit)
	}

	if p.Page != 0 {
		params["page"] = strconv.Itoa(p.Page)
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

func FetchPostList(limit int, page int, tags []string) {
	// TODO poner constantes o algo para cambiar de origen de gelbooru y hacer el fetch de posts
}
