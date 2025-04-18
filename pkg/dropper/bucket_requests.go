package dropper

import (
	"log/slog"
)

func (c *Client) GetAllBuckets() ([]Bucket, error) {
	response, err := c.request("GET", "/api/bucket/all", nil)

	if err != nil {
		return nil, err
	}

	resource, err := unmarshallRequestResponse[ResourceResponse[[]Bucket]](response)

	if err != nil {
		return nil, err
	}

	buckets := resource.Data
	return buckets, nil
}

func (c *Client) HasBuckets() bool {
	buckets, err := c.GetAllBuckets()

	if err != nil {
		slog.Warn("failed fetching dropper buckets", "error", err)
		return false
	}

	return len(buckets) > 0
}
