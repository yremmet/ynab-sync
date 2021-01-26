package client

import "net/http"

type ClientError struct {
	reponse *http.Response
	errorFunc func(clientError *ClientError) string
	error error

}

func (m *ClientError) Error() string {
	if (m.error != nil) {
		return m.error.Error()
	}
	return m.errorFunc(m)
}

func defaultError(response *http.Response, error error) *ClientError {

	return &ClientError{
		reponse: response,
		errorFunc: func(clientError *ClientError) string {
			return clientError.reponse.Status
		},
		error: error,
	}
}


func httpError(response *http.Response) *ClientError {

	return &ClientError{
		reponse: response,
		errorFunc: func(clientError *ClientError) string {
			return clientError.reponse.Status
		},
	}
}

