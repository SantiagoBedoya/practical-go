package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

// BodyRequest define data structure for request
type BodyRequest struct {
	Sides int `json:"sides"`
}

// BodyResponse define data structure for response
type BodyResponse struct {
	Result int `json:"result"`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var response events.APIGatewayProxyResponse
	body := BodyRequest{}

	if err := json.Unmarshal([]byte(request.Body), &body); err != nil {
		res, _ := json.Marshal(map[string]string{
			"message": "we cannot unmarshal JSON body",
		})
		response = events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       string(res),
		}
		return response, nil
	}
	bodyResponse := BodyResponse{
		Result: rand.Intn((body.Sides+1)-1) + 1,
	}
	res, _ := json.Marshal(bodyResponse)
	response = events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       string(res),
	}

	return response, nil
}
