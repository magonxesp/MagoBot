package dropper

import "errors"

type DropRequest struct {
	Source     string `json:"source"`
	BucketName string `json:"bucket_name"`
}

func (c *Client) Drop(source string, bucket *Bucket) error {
	reader, err := createRequestBodyReader(&DropRequest{
		Source:     source,
		BucketName: bucket.Name,
	})

	if err != nil {
		return err
	}

	response, err := c.request("POST", "/api/bucket/all", reader)

	if err != nil {
		return err
	}

	status, err := unmarshallRequestResponse[Status](response)

	if err != nil {
		return err
	}

	switch response.StatusCode {
	case 400:
	case 500:
		return errors.New(status.Error)
	}

	return nil
}
