package controller

import (
	"encoding/json"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"net/http"
)

type TrialController struct {
	AbstractController
}

func (controller TrialController) TrialGetIPController(writer http.ResponseWriter, request *http.Request) {
	response := make(map[string]interface{})
	ip := request.RemoteAddr
	xforward := request.Header.Get("X-Forwarded-For")

	response["IP"] = ip
	response["X-Forwarded-For"] = xforward

	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	helper.PanicIfError(err)
}
