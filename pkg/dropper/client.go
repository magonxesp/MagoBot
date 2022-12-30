package dropper

type Client struct {
	BaseUrl      string
	ClientKey    string
	ClientSecret string
	BearerToken  string
}

func NewClient(baseUrl string, clientKey string, clientSecret string) *Client {
	return &Client{
		BaseUrl:      baseUrl,
		ClientKey:    clientKey,
		ClientSecret: clientSecret,
	}
}
