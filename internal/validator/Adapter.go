package validator

import (
	"fmt"

	localsEn "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"github.com/kosatnkn/catalyst/v3/app/adapters"
)

// Adapter is used to validate structures and variables
// against a set of validation rules.
type Adapter struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
}

// NewAdapter creates a new validator adapter instance.
func NewAdapter() (adapters.ValidatorAdapterInterface, error) {
	a := &Adapter{}

	a.validate = validator.New()

	en := localsEn.New()
	a.uni = ut.New(en, en)

	return a, nil
}

// Validate validates structs and slices of structs.
func (a *Adapter) Validate(data interface{}) map[string]string {
	if isSlice(data) {
		return a.validateSliceOfStructs(convertToSlice(data))
	}

	return a.validateStruct(data)
}

// ValidateField validates a single variable.
func (a *Adapter) ValidateField(name string, value interface{}, rules string) map[string]string {
	// returns nil or ValidationErrors ( []FieldError )
	err := a.validate.Var(value, rules)
	if err == nil {
		return nil
	}

	// from here you can create your own error messages in whatever language you wish
	errs := err.(validator.ValidationErrors)
	result := errs.Translate(a.getTranslator("en"))

	result[name] = result[""]
	delete(result, "")

	return result
}

// validateStruct validates a struct.
func (a *Adapter) validateStruct(data interface{}) map[string]string {
	// returns nil or ValidationErrors ( []FieldError )
	err := a.validate.Struct(data)
	if err == nil {
		return nil
	}

	// from here you can create your own error messages in whatever language you wish
	errs := err.(validator.ValidationErrors)

	return errs.Translate(a.getTranslator("en"))
}

// validateStruct validates an array of structs.
func (a *Adapter) validateSliceOfStructs(data []interface{}) map[string]string {
	e := make(map[string]string)

	for i, d := range data {
		res := a.validateStruct(d)
		// prepend the key name of the struct field with the index position of it in the array
		for k, v := range res {
			e[fmt.Sprintf(`%d_%s`, i, k)] = v
		}
	}

	// NOTE: need to send 'nil' when there are no errors
	if len(e) == 0 {
		return nil
	}

	return e
}

// getTranslator returns a translator for the given locale.
func (a *Adapter) getTranslator(locale string) ut.Translator {
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := a.uni.GetTranslator(locale)

	enTranslations.RegisterDefaultTranslations(a.validate, trans)

	return trans
}
