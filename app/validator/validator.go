package validator

import (
	localsEn "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

// use a single instance , it caches struct info
var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
)

// Validate validates bound values of an unpacker struct against validation rules defined in that unpacker struct.
func Validate(data interface{}) map[string]string {

	validate = validator.New()

	en := localsEn.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	enTranslations.RegisterDefaultTranslations(validate, trans)

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	// from here you can create your own error messages in whatever language you wish
	errs := err.(validator.ValidationErrors)

	return errs.Translate(trans)
}

// ValidateField validates a single variable.
func ValidateField(field interface{}, rules string) map[string]string {

	validate = validator.New()

	en := localsEn.New()
	uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := uni.GetTranslator("en")

	enTranslations.RegisterDefaultTranslations(validate, trans)

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Var(field, rules)
	if err == nil {
		return nil
	}

	// from here you can create your own error messages in whatever language you wish
	errs := err.(validator.ValidationErrors)

	return errs.Translate(trans)
}