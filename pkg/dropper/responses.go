package dropper

type AuthorizationResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type Bucket struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
}

type Status struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}
