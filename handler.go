package function

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	handler "github.com/openfaas/templates-sdk/go-http"
)

// Handle handles the http request, sleeps a specified duration set as an
// environment variable, and returns a simple response.
// Accepts a Duration as a raw string as an override
func Handle(req handler.Request) (handler.Response, error) {
	var err error
	duration, err := time.ParseDuration(os.Getenv("SLEEP_DURATION"))
	if err != nil {
		msg := "Unable to parse SLEEP_DURATION"
		log.Println(msg)
		return handler.Response{
			StatusCode: http.StatusInternalServerError,
			Body:       []byte(msg),
		}, nil
	}
	if durationOverride, err := time.ParseDuration(string(req.Body)); err == nil {
		duration = durationOverride
	}

	log.Println(fmt.Sprintf("sleeping %s", duration))
	time.Sleep(duration)
	sleptMsg := fmt.Sprintf("slept %s", duration)
	log.Println(sleptMsg)

	return handler.Response{
		StatusCode: http.StatusOK,
		Body:       []byte(sleptMsg),
	}, nil
}
