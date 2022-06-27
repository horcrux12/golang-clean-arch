package errorModel

import "errors"

type ErrorExceptionModel struct {
	Code            int
	Error           error
	ErrorCause      error
	ErrorMessages   string
	ErrorParameters []ErrorParameter
	FileName        string
	FuncName        string
}

type ErrorParameter struct {
	ParamKey   string
	ParamValue string
}

func GenerateErrorModel(code int, fileName, funcName, err string, causedBY error) ErrorExceptionModel {
	var errModel ErrorExceptionModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.ErrorCause = causedBY
	errModel.FileName = fileName
	errModel.FuncName = funcName

	return errModel
}

func GenerateErrorModelWithParam(code int, fileName, funcName, err string, errorParam []ErrorParameter) ErrorExceptionModel {
	var errModel ErrorExceptionModel
	errModel.Code = code
	errModel.Error = errors.New(err)
	errModel.ErrorParameters = errorParam
	errModel.FileName = fileName
	errModel.FuncName = funcName

	return errModel
}
