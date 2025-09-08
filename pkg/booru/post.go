package booru

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/MagonxESP/MagoBot/internal/infraestructure/helpers"
)

type Posts struct {
	XMLName xml.Name `xml:"posts"`
	Count   int      `xml:"count,attr"`
	Offset  int      `xml:"offset,attr"`
	Items   []Post   `xml:"post"`
}

type Post struct {
	Height        int    `xml:"height,attr"`
	Score         int    `xml:"score,attr"`
	FileURL       string `xml:"file_url,attr"`
	ParentID      string `xml:"parent_id,attr"`
	SampleURL     string `xml:"sample_url,attr"`
	SampleWidth   int    `xml:"sample_width,attr"`
	SampleHeight  int    `xml:"sample_height,attr"`
	PreviewURL    string `xml:"preview_url,attr"`
	Rating        string `xml:"rating,attr"`
	Tags          string `xml:"tags,attr"`
	ID            int    `xml:"id,attr"`
	Width         int    `xml:"width,attr"`
	Change        int64  `xml:"change,attr"`
	MD5           string `xml:"md5,attr"`
	CreatorID     int    `xml:"creator_id,attr"`
	HasChildren   bool   `xml:"has_children,attr"`
	CreatedAt     string `xml:"created_at,attr"`
	Status        string `xml:"status,attr"`
	Source        string `xml:"source,attr"`
	HasNotes      bool   `xml:"has_notes,attr"`
	HasComments   bool   `xml:"has_comments,attr"`
	PreviewWidth  int    `xml:"preview_width,attr"`
	PreviewHeight int    `xml:"preview_height,attr"`
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
	// The Booru on will perform the request
	Booru string
	// The API Key if it is required for example for Rule34
	ApiKey string
	// The UserId that is the owner of the api key if it is required for example for Rule34
	UserId string
}

func NewPostListRequest(booru string, tags []string) *PostListRequest {
	return &PostListRequest{
		Limit: 100,
		Page:  0,
		Tags:  tags,
		Booru: booru,
	}
}

func NewRule34PostListRequest(apiKey string, userId string, tags []string) *PostListRequest {
	return &PostListRequest{
		Limit: 100,
		Page:  0,
		Tags:  tags,
		Booru: Rule34,
		ApiKey: apiKey,
		UserId: userId,
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

	if p.ApiKey != "" {
		params["api_key"] = p.ApiKey
	}

	if p.UserId != "" {
		params["user_id"] = p.UserId
	}

	return strings.Join(helpers.MapToKeyValueList(params, "="), "&")
}

func GetPostList(request *PostListRequest) (*Posts, error) {
	url := fmt.Sprintf(
		"%s/index.php?page=dapi&s=post&q=index&%s",
		request.Booru,
		request.ToQueryString(),
	)

	slog.Debug("fetching booru posts", "url", url, "booru", request.Booru)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != 200 {
		return nil, NewPostRequestError(fmt.Sprintf("An error ocurred fetching the post list from %s", url))
	}

	content, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Warn("failed reading booru response body", "error", err, "url", request.Booru)
		return nil, err
	}

	slog.Debug("booru response body", "body", string(content), "booru", request.Booru)
	posts, err := deserializeResponse(content)

	if err != nil {
		slog.Warn("failed deserializing booru response body", "error", err, "booru", request.Booru)
		return nil, err
	}

	return posts, nil
}

func deserializeResponse(content []byte) (*Posts, error) {
	var posts Posts

	err := xml.Unmarshal(content, &posts)
	if err != nil {
		return nil, err
	}

	return &posts, nil
}
