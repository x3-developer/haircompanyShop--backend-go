package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"serv_shop_haircompany/internal/shared/utils/response"
)

func ValidateDTO(DTO interface{}) []response.ErrorField {
	var errorFields []response.ErrorField
	validate := validator.New()

	err := validate.Struct(DTO)
	if err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			for _, ve := range validationErrors {
				fieldName := ve.Field()
				tag := ve.Tag()
				errorCode := response.GetErrorCodeByTag(tag)

				errorFields = append(errorFields, response.NewErrorField(fieldName, string(errorCode)))
			}
		}
	}

	return errorFields
}
