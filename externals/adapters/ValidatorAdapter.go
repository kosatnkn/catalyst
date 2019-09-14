package adapters

import (
	localsEn "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/kosatnkn/catalyst/domain/boundary/adapters"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
)

// ValidatorAdapter is used to validate data structures.
type ValidatorAdapter struct {
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
}

// NewValidatorAdapter creates a new instance of the adapter.
func NewValidatorAdapter() (adapters.ValidatorAdapterInterface, error) {

	a := ValidatorAdapter{}

	a.validate = validator.New()

	en := localsEn.New()
	a.uni = ut.New(en, en)

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	trans, _ := a.uni.GetTranslator("en")
	a.trans = trans

	enTranslations.RegisterDefaultTranslations(a.validate, a.trans)

	return &a, nil
}

// Validate validates bound values of an unpacker struct against validation rules defined in that unpacker struct.
func (a *ValidatorAdapter) Validate(data interface{}) map[string]string {

	// returns nil or ValidationErrors ( []FieldError )
	err := a.validate.Struct(data)
	if err == nil {
		return nil
	}

	// from here you can create your own error messages in whatever language you wish
	errs := err.(validator.ValidationErrors)

	return errs.Translate(a.trans)
}

// ValidateField validates a single variable.
func (a *ValidatorAdapter) ValidateField(field interface{}, rules string) map[string]string {

	// returns nil or ValidationErrors ( []FieldError )
	err := a.validate.Var(field, rules)
	if err == nil {
		return nil
	}

	// from here you can create your own error messages in whatever language you wish
	errs := err.(validator.ValidationErrors)

	return errs.Translate(a.trans)
}
