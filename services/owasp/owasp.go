package services

import (
	"io/ioutil"
	"strings"

	proxy "github.com/mamaoag/binosearch/services/http"
	url "github.com/mamaoag/binosearch/services/url"
)

func unAuthCodes(statusCode int) bool {
	switch statusCode {
	case
		200,
		400,
		429,
		500,
		503:
		return true
	}

	return false
}

// API1:2019 - Broken Object Level Auth. Checks for GUIDs
func BrokenObjectLevelAuth(endpoints []url.Url) (bool, error) {

	var endpointIssue []url.Url
	var message string

	for i := 0; i < len(endpoints); i++ {
		res, err := proxy.Request(endpoints[i].Full)

		if err != nil {
			return false, err
		}

		message = "Endpoint shows no data."

		if unAuthCodes(res.StatusCode) {
			defer res.Body.Close()

			body, err := ioutil.ReadAll(res.Body)

			if err != nil {
				return false, err
			}

			bodyString := string(body)

			if strings.Contains(bodyString, "[") {
				endpointIssue = append(endpointIssue, endpoints[i])
				message = "Endpoint shows data."
			} else if strings.Contains(bodyString, "id") {
				endpointIssue = append(endpointIssue, endpoints[i])
				message = "Endpoint shows data."
			}
		}

		response := proxy.LogHttpResponse{
			Path:       endpoints[i].Path,
			StatusCode: res.StatusCode,
			Message:    message,
		}

		proxy.LogResponse(response)
	}

	if len(endpointIssue) == 0 {
		return false, nil
	}

	return true, nil
}
