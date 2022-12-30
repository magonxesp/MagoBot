package dropper

import "log"

func (c *Client) GetAllBuckets() ([]Bucket, error) {
	response, err := c.request("GET", "/api/bucket/all", nil)

	if err != nil {
		return nil, err
	}

	buckets, err := unmarshallRequestResponse[[]Bucket](response)
	return *buckets, err
}

func (c *Client) HasBuckets() bool {
	buckets, err := c.GetAllBuckets()

	if err != nil {
		log.Println(err)
		return false
	}

	return len(buckets) > 0
}
