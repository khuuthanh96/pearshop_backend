package payload

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"

	appErrors "pearshop_backend/app/errors"
	"pearshop_backend/pkg/validation"
)

type RequestPayload interface {
	StructName() string
}

type fieldConfig struct {
	Code         appErrors.Code    `yaml:"Code"`
	MessageByTag map[string]string `yaml:"MessageByTag"`
}

type payloadConfig map[string]fieldConfig

type config map[string]payloadConfig

func defaultUndefinedField(payloadName, fieldName, tagName string) appErrors.SystemError {
	return appErrors.NewInvalidArgumentErr(
		appErrors.CodeInvalidPayload,
		fmt.Sprintf("missing defined code for payload:%s field:%s tag:%s", payloadName, fieldName, tagName),
		nil,
	)
}

func (c config) dispatch(payloadName, fieldName, tagName string, param interface{}) appErrors.SystemError {
	var (
		paylConfig payloadConfig
		fConfig    fieldConfig
		tagMessage string
		ok         bool
	)

	paylConfig, ok = c[payloadName]
	if !ok {
		return defaultUndefinedField(payloadName, "", "")
	}

	fConfig, ok = paylConfig[fieldName]
	if !ok {
		return defaultUndefinedField(payloadName, fieldName, "")
	}

	tagMessage, ok = fConfig.MessageByTag[tagName]
	if !ok {
		return defaultUndefinedField(payloadName, fieldName, tagName)
	}

	return appErrors.NewInvalidArgumentErr(
		fConfig.Code,
		tagMessage,
		param,
	)
}

func validate(p RequestPayload) error {
	if err := validation.GetInstance().Struct(p); err != nil {
		switch e := err.(type) {
		case validator.ValidationErrors:
			errs := make(appErrors.SystemErrors, 0, len(e))
			for _, ee := range e {
				errs = append(errs, globalPayloadConfig.dispatch(
					p.StructName(),
					normalizeFieldName(ee.Field()),
					strings.ToLower(ee.Tag()),
					ee.Value(),
				))
			}

			return errs
		default:
			return err
		}
	}

	return nil
}

// normalizeFieldName remove suffix [0] for field type is slice error. Ex: Users[0] -> Users
func normalizeFieldName(name string) string {
	i := strings.Index(name, "[")

	if i > 0 {
		return name[:i]
	}

	return name
}
