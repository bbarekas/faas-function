package function

import (
	"fmt"
	"net/http"

	handler "github.com/openfaas/templates-sdk/go-http"
)

// Handle a function invocation
func Handle(req handler.Request) (handler.Response, error) {
	var err error

	message := fmt.Sprintf("Method:%s\n", string(req.Method))
	message = message + fmt.Sprintf("Body:\n%s", string(req.Body))
	message = message + fmt.Sprintf("\nHeaders:\n")
	//message := message + fmt.Sprintf(":\n%s", string(req.Header.))

	// Loop over header names
	for name, values := range req.Header {
		// Loop over all values for the name.
		for _, value := range values {
			message = message + fmt.Sprintf("> %s : %s\n", name, value)
		}
	}

	return handler.Response{
		Body:       []byte(message),
		StatusCode: http.StatusOK,
	}, err
}
