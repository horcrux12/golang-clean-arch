package errorModel

import (
	"github.com/go-playground/validator/v10"
	"github.com/horcrux12/clean-rest-api-template/constanta"
)

type ErrorValidatorModel struct {
	Tag          string
	Messages     map[string]string
	ValidateFunc func(validator.FieldLevel) bool
}

func InitiateNewErrorsValidator() (result []ErrorValidatorModel) {
	result = append(result,
		ErrorValidatorModel{
			Tag: "is-username",
			Messages: map[string]string{
				constanta.ENLangConstanta: "Can only start with character, must write in lowercase, only contains [character, numeric, dot and underscore]",
				constanta.IDLangConstanta: "Harus diawali oleh huruf, harus huruf kecil, hanya boleh mangandung [huruf, angka, titik dan garis bawah]",
			},
			ValidateFunc: validateUsername,
		},
	)
	return
}
