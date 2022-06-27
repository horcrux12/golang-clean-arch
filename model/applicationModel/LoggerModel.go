package applicationModel

import (
	"go.uber.org/zap"
	"os"
)

type LoggerModel struct {
	IP          string `json:"ip"`
	PID         int    `json:"pid"`
	Time        int64  `json:"time"`
	Status      int    `json:"status"`
	Location    string `json:"location"`
	Application string `json:"application"`
	Version     string `json:"version"`
	Code        string `json:"code"`
	Message     string `json:"message"`
}

func GenerateLogModel(version string, application string) (output LoggerModel) {
	output.IP = "-"
	output.PID = os.Getpid()
	output.Application = application
	output.Version = version
	output.Code = "-"
	output.Message = "-"
	output.Location = "-"
	return output
}

func (object LoggerModel) ToLoggerObject() (output []zap.Field) {
	output = append(output, zap.String("ip", object.IP))
	output = append(output, zap.Int("pid", object.PID))
	output = append(output, zap.String("application", object.Application))
	output = append(output, zap.String("version", object.Version))
	output = append(output, zap.Int64("time", object.Time))
	output = append(output, zap.String("location", object.Location))
	output = append(output, zap.Int("status", object.Status))
	output = append(output, zap.String("code", object.Code))
	output = append(output, zap.String("message", object.Message))

	return output
}
