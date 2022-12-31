package dropper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Client struct {
	BaseUrl      string
	ClientId     string
	ClientSecret string
	BearerToken  string
}

func NewClient(baseUrl string, clientId string, clientSecret string) *Client {
	return &Client{
		BaseUrl:      strings.TrimSuffix(baseUrl, "/"),
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
}

func (c *Client) request(method string, endpoint string, body io.Reader) (*http.Response, error) {
	client := &http.Client{}
	request, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.BaseUrl, endpoint), body)

	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	if c.BearerToken != "" {
		request.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.BearerToken))
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func unmarshallRequestResponse[T interface{}](response *http.Response) (*T, error) {
	content, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	var decoded T
	stringContent := string(content)
	log.Println(stringContent)
	err = json.Unmarshal(content, &decoded)

	if err != nil {
		return nil, err
	}

	return &decoded, nil
}

func createRequestBodyReader(data interface{}) (io.Reader, error) {
	encoded, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	return bytes.NewReader(encoded), nil
}

func (c *Client) Authenticate() error {
	response, err := c.request(
		"GET",
		fmt.Sprintf(
			"/oauth2/token?grant_type=client_credentials&client_id=%s&client_secret=%s&scope=read",
			c.ClientId,
			c.ClientSecret,
		),
		nil,
	)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf("an error occurred during authentication, status code %d", response.StatusCode)
	}

	authorization, err := unmarshallRequestResponse[AuthorizationResponse](response)

	if err != nil {
		return err
	}

	c.BearerToken = authorization.AccessToken
	return nil
}
