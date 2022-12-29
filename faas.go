package utils

import "errors"

// GetFunctionListeningPort will return the port on which the function is
// listening to and the cloud it believes the function is running in, or an
// error. Currently only supports Azure Functions.
func GetFunctionListeningPort() (port, cloud string, err error) {
	if port := EnvOrDefault("FUNCTIONS_CUSTOMHANDLER_PORT", ""); port != "" {
		return port, "azure", nil
	}

	return "", "", errors.New("cloud environment not determinable")
}
