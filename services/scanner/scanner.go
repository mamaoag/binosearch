package services

import (
	"log"

	proxy "github.com/mamaoag/binosearch/services/http"
	url "github.com/mamaoag/binosearch/services/url"
)

func ScanEndpoint(resource url.Url) {
	var message string

	res, err := proxy.Request(resource.Full)

	if err != nil {
		log.Fatalln(err)
	}

	code := res.StatusCode

	if code == 404 {
		message = "Not found. ❌"
	} else {
		message = "Found. ✅"
	}

	response := proxy.LogHttpResponse{
		Path:       resource.Path,
		StatusCode: res.StatusCode,
		Message:    message,
	}

	proxy.LogResponse(response)
}
