package presenter

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	appErrors "pearshop_backend/app/errors"
	appLog "pearshop_backend/pkg/log"
)

type response struct {
	Data   interface{}     `json:"data,omitempty"`
	Paging interface{}     `json:"paging,omitempty"`
	Errors []responseError `json:"errors,omitempty"`
}

// ResponsePaging holds paging response information
type ResponsePaging struct {
	Total uint32 `json:"total"`
}

type responseError struct {
	Type    string      `json:"type"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Param   interface{} `json:"param,omitempty"`
}

// RenderErrors returns error response
func RenderErrors(ctx *gin.Context, err error) {
	// ignore context cancalled error
	if errors.Is(err, appErrors.ErrContextCancelled) {
		return
	}

	if errs, ok := err.(appErrors.SystemErrors); ok {
		// if err is a list of errors, we assume they are validation errors,
		// so we will always return http.StatusBadRequest
		ctx.JSON(http.StatusBadRequest, response{
			Errors: fromSystemErrors(errs),
		})

		return
	}

	if e, ok := err.(appErrors.SystemError); ok {
		ctx.JSON(e.StatusCode(), response{
			Errors: fromSystemErrors(appErrors.SystemErrors{e}),
		})

		return
	}

	appLog.
		WithField("url", ctx.Request.URL.String()).
		WithField("method", ctx.Request.Method).
		WithField("user_agent", ctx.Request.UserAgent()).
		WithField("ip", ctx.ClientIP()).
		WithError(err).
		Errorln("internal server error")

	ctx.JSON(http.StatusInternalServerError, response{
		Errors: []responseError{{
			Type:    string(appErrors.TypeInternal),
			Code:    string(appErrors.CodeInternal),
			Message: "internal server error",
			Param:   nil,
		}},
	})
}

// RenderData returns data response
func RenderData(ctx *gin.Context, data, paging interface{}) {
	ctx.JSON(http.StatusOK, response{
		Data:   data,
		Paging: paging,
	})
}

// parse system errors to response errors
func fromSystemErrors(errs appErrors.SystemErrors) []responseError {
	if len(errs) == 0 {
		return nil
	}

	respErrors := make([]responseError, len(errs))

	for i, e := range errs {
		respErrors[i] = responseError{
			Type:    string(e.Type()),
			Code:    string(e.Code()),
			Message: e.Message(),
			Param:   e.Param(),
		}
	}

	return respErrors
}
