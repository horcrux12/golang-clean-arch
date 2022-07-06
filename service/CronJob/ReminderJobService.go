package CronJob

import (
	"fmt"
	"github.com/horcrux12/clean-rest-api-template/app"
	"github.com/horcrux12/clean-rest-api-template/dto/in"
	"github.com/horcrux12/clean-rest-api-template/helper"
	"github.com/horcrux12/clean-rest-api-template/service"
	"time"
)

type reminderService struct {
	service.AbstractService
}

var ReminderService = reminderService{}.New()

func (input reminderService) New() reminderService {
	return reminderService{
		AbstractService: service.AbstractService{
			FileName:    "Test",
			ServiceName: "Test",
		},
	}
}

func (input reminderService) AddScheduler(inputRequest in.ToDoListRequest) {
	fmt.Println(inputRequest.UUIDKey)
	if !inputRequest.ReminderDate.IsZero() {
		if !inputRequest.ReminderDate.Equal(inputRequest.RepeatFromDate) {
			_, err := app.ApplicationAttribute.CronScheduler.Every(1).
				Second().
				StartAt(inputRequest.ReminderDate).
				LimitRunsTo(1).
				Tag(inputRequest.UUIDKey).
				Do(input.reminderJob, inputRequest.UUIDKey)

			if err != nil {
				helper.PanicIfError(err)
			}
		}

		if !inputRequest.RepeatFromDate.IsZero() || inputRequest.RepeatFromDate.Equal(inputRequest.ReminderDate) {
			if inputRequest.RepeatEvery < 1 {
				inputRequest.RepeatEvery = 1
			}
			app.ApplicationAttribute.CronScheduler.Every(int(inputRequest.RepeatEvery))
			switch inputRequest.RepeatType {
			case "S":
				app.ApplicationAttribute.CronScheduler.Second()
				break
			case "D":
				app.ApplicationAttribute.CronScheduler.Day()
				break
			case "W":
				app.ApplicationAttribute.CronScheduler.Week()
				break
			case "M":
				app.ApplicationAttribute.CronScheduler.Month(inputRequest.RepeatFromDate.Day())
				break
			case "WD":
				app.ApplicationAttribute.CronScheduler.Week().
					Weekday(time.Monday).Weekday(time.Tuesday).Weekday(time.Wednesday).
					Weekday(time.Thursday).Weekday(time.Friday)
				break
			}
			_, err := app.ApplicationAttribute.CronScheduler.Tag(inputRequest.UUIDKey).Do(input.reminderJob, inputRequest.UUIDKey)
			if err != nil {
				helper.PanicIfError(err)
			}
		}
	}
}

func (input reminderService) reminderJob(value string) {
	fmt.Println(value)
}
