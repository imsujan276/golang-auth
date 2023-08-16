package utils

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type Error struct {
	Results struct {
		Errors []map[string]map[string]string `json:"errors"`
	} `json:"results"`
}

func GoValidator(s interface{}) (interface{}, int) {
	validate := validator.New()
	enLocale := en.New()                // Create an instance of the English locale
	uni := ut.New(enLocale, enLocale)   // Create a universal translator
	trans, _ := uni.GetTranslator("en") // Get the translator for the English locale

	// Attach the translator to the validator
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	err := validate.Struct(s)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []map[string]map[string]string
			for _, e := range validationErrors {
				fieldName := e.StructField()
				validationTag := e.Tag()
				validationMessage := e.Translate(trans)

				errorMap := map[string]map[string]string{
					fieldName: {
						"tag":     validationTag,
						"message": validationMessage,
					},
				}
				errors = append(errors, errorMap)
			}
			return errors, len(validationErrors)
		}
	}

	return nil, 0
}
