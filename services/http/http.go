package services

import (
	"crypto/tls"
	"log"
	"net/http"
)

type LogHttpResponse struct {
	Path       string
	StatusCode int
	Message    string
}

type HttpResponse struct {
	Path       string
	StatusCode int
	Message    string
}

func Request(url string) (*http.Response, error) {
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	res, err := http.Get(url)

	return res, err
}

func LogResponse(res LogHttpResponse) {
	log.Printf(
		"%d %d: GET %s: Status: %d - %s\n",
		log.Ldate,
		log.Ltime,
		res.Path,
		res.StatusCode,
		res.Message,
	)
}

func Response(res HttpResponse) {
	log.Printf(
		"%d %d: %s (Status: %d)\n",
		log.Ldate,
		log.Ltime,
		res.Message,
		res.StatusCode,
	)
}
