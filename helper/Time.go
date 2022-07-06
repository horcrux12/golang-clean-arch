package helper

import "time"

func TimeStrToTimeWithTimeFormat(timeStr, format string) (output time.Time) {
	output, err := time.Parse(format, timeStr)
	if err != nil {
		PanicIfError(err)
	}
	return
}
