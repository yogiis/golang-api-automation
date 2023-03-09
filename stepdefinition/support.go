package stepdefinition

import (
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/yogiis/golang-api-automation/helper"
)

func NewHTTPClient() *http.Client {
	return &http.Client{}
}

func SendHTTPRequest(request *http.Request) (*http.Response, error) {
	client := NewHTTPClient()
	response, _ := client.Do(request)
	return response, nil
}

func GetAPIKey() string {
	env := godotenv.Load()
	helper.LogPanicln(env)
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		panic("API_KEY environment variable not set")
	}
	return apiKey
}

func CloseResponseBody(response *http.Response) {
	if response != nil && response.Body != nil {
		response.Body.Close()
	}
}

func AddAPIKeyHeaderIfHas(hitEndpoint *http.Request) {
	apiKey := GetAPIKey()
	if apiKey != "" {
		hitEndpoint.Header.Set("api-key", apiKey)
	}
}

func ReaderResponseBody(response *http.Response) ([]byte, error) {
	return io.ReadAll(response.Body)
}
