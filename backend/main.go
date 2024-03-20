package main

import (
	"encoding/json"
	"errors"

	"github.com/JamesWilliamPage/cortex-helper-backend/core"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct {
	Action string          `json:"action"`
	Data   json.RawMessage `json:"data"`
}

type Handler func(data json.RawMessage) (interface{}, error)

var handlers = map[string]Handler{
	"getCharacters": core.GetCharactersHandler,
	// Add more handlers here.
}

func HandleRequest(req Request) (interface{}, error) {
	handler, ok := handlers[req.Action]
	if !ok {
		return nil, errors.New("invalid action")
	}
	return handler(req.Data)
}

func main() {
	lambda.Start(HandleRequest)
}
