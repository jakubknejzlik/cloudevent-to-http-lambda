package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	handler "github.com/jakubknejzlik/cloudevents-lambda-handler"

	cloudevents "github.com/cloudevents/sdk-go"
)

func Receive(url string) func(event cloudevents.Event) error {

	return func(event cloudevents.Event) (err error) {
		fmt.Println("new event", event.ID(), event.Data, "; content type:", event.DataContentType(), "; media type:", event.DataMediaType())
		data := event.Data.([]byte)

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
		if err != nil {
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		fmt.Println("received response", resp.Status, resp.StatusCode)
		if err != nil {
			return
		}

		return
	}
}

func main() {
	url := os.Getenv("URL")
	h := handler.NewCloudEventsLambdaHandler(Receive(url))
	h.Start()
}
