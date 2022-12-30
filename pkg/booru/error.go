package booru

type PostRequestError struct {
	message string
}

func NewPostRequestError(message string) *PostRequestError {
	return &PostRequestError{
		message: message,
	}
}

func (e *PostRequestError) Error() string {
	return e.message
}
