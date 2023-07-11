package validation

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	globalValidator    *validator.Validate
	enUniTrans         *ut.UniversalTranslator
	englishTranslator  ut.Translator
	enLocales          locales.Translator
	jaUniTrans         *ut.UniversalTranslator
	japaneseTranslator ut.Translator
	jpLocales          locales.Translator
)

const (
	japanese = "ja"
	english  = "en"
)

func init() {
	globalValidator = validator.New()

	enLocales = en.New()
	enUniTrans = ut.New(enLocales, enLocales)
	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)
	englishTranslator, _ = enUniTrans.GetTranslator(english)
	_ = en_translations.RegisterDefaultTranslations(globalValidator, englishTranslator)
}

// GetTranslator to get translator with language
func GetTranslator(lang string) ut.Translator {
	switch lang {
	case japanese:
		return japaneseTranslator
	default:
		return englishTranslator
	}
}

// GetInstance global validator
func GetInstance() *validator.Validate {
	return globalValidator
}
