package errorModel

import "net/http"

func GenerateDataNotFound(fileName, funcName string) ErrorExceptionModel {
	return GenerateErrorModel(400, fileName, funcName, "ERR-4-DBS-001", nil)
}

func GenerateUnknownDataError(fileName, funcName, fieldName string) ErrorExceptionModel {
	errorParam := make([]ErrorParameter, 1)
	errorParam[0].ParamKey = "FieldName"
	errorParam[0].ParamValue = fieldName
	return GenerateErrorModelWithParam(400, fileName, funcName, "ERR-4-DBS-002", errorParam)
}

func GenerateForbiddenAccessError(fileName, funcName string) ErrorExceptionModel {
	return GenerateErrorModel(http.StatusForbidden, funcName, fileName, "ERR-4-SRV-002", nil)
}
