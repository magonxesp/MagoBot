package dropper

func (c *Client) GetAllBuckets() ([]Bucket, error) {
	response, err := c.request("GET", "/api/bucket/all", nil)

	if err != nil {
		return nil, err
	}

	buckets, err := unmarshallRequestResponse[[]Bucket](response)
	return *buckets, err
}
