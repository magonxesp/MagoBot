package booru

import (
	"encoding/json"
	"io"
)

func UnmarshalResponseBody(body io.ReadCloser, v any) error {
	content, err := io.ReadAll(body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &v)

	if err != nil {
		return err
	}

	return nil
}
