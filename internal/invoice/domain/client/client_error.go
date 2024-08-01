package client

type ClientNotFoundError struct {
}

func (e ClientNotFoundError) Error() string {
	return "client not found"
}
