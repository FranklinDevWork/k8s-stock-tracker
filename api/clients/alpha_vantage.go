package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/FranklinDevWork/k8s-stock-tracker/api/models"
)

type AlphaVantageClient struct {
	Url    string
	ApiKey string
}

func (client *AlphaVantageClient) MakeQueryRequest(symbol, function string) (models.AlphaVantageResponse, error) {
	req, _ := http.NewRequest("GET", client.Url, nil)
	query := req.URL.Query()
	query.Add("function", function)
	query.Add("symbol", symbol)
	query.Add("apikey", client.ApiKey)
	req.URL.RawQuery = query.Encode()

	httpClient := &http.Client{}
	response, err := httpClient.Do(req)
	if err != nil {
		return models.AlphaVantageResponse{}, errors.New("issues getting response")
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.AlphaVantageResponse{}, errors.New("issues reading response body")
	}

	// the api seems to return 200s even for an error :explodinghead:
	if ok := checkError(body); ok != nil {
		return models.AlphaVantageResponse{}, errors.New("issues getting valid data")
	}

	var responseData models.AlphaVantageResponse
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return models.AlphaVantageResponse{}, err
	}

	return responseData, nil
}

func checkError(body []byte) error {
	var errorResponse map[string]interface{} // check if the response is an error/information
	if err := json.Unmarshal(body, &errorResponse); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return err
	}

	if _, ok := errorResponse["Error Message"]; ok {
		return errors.New("issues getting valid response")
	}
	// likely hit a rate limit?
	if _, ok := errorResponse["Information"]; ok {
		return errors.New("issues getting valid response")
	}

	return nil
}
