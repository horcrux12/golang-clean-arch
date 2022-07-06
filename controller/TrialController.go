package controller

import (
	"encoding/json"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/dto/out"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/service/CronJob"
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

func (controller TrialController) CobaReminder(writer http.ResponseWriter, request *http.Request) {
	var inputRequest in.ToDoListRequest
	helper.ReadFromRequestBody(request, &inputRequest)
	inputRequest.UUIDKey = helper.GetUUID()

	dateFormat := "2006-01-02 15:04:05"
	inputRequest.DueDate = helper.TimeStrToTimeWithTimeFormat(inputRequest.DueDateStr, dateFormat)
	inputRequest.ReminderDate = helper.TimeStrToTimeWithTimeFormat(inputRequest.ReminderDateStr, dateFormat)
	inputRequest.RepeatFromDate = helper.TimeStrToTimeWithTimeFormat(inputRequest.RepeatFromDateStr, dateFormat)
	inputRequest.RepeatUntilDate = helper.TimeStrToTimeWithTimeFormat(inputRequest.RepeatUntilDateStr, dateFormat)

	//fmt.Println(helper.StructToJSON(in.ToDoListRequest{}))

	switch request.Method {
	case "POST":
		CronJob.ReminderService.AddScheduler(inputRequest)
		break
	}

	helper.WriteToResponseBody(writer, out.WebResponse{Payload: out.PayloadWebResponse{Data: "SUCCESS"}})
}
