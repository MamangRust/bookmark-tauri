package response

type ServiceError struct {
	Err         error
	Description string
}

func (se ServiceError) Error() string {
	return se.Description + ": " + se.Err.Error()
}

type ServiceResponse struct {
	Data  interface{}
	Error error
}
