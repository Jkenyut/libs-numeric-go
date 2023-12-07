package model_response

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
)

type ResponseDefault struct {
	Success    bool   `json:"success"`
	Message    string `json:"message"`
	StatusCode string `json:"statusCode"`
}
type DefaultResponse struct {
	Message ResponseDefault `json:"message,omitempty"`
	Data    any             `json:"data,omitempty"`
}

func DefaultErrorResponseWithMessage(msg string, status int) DefaultResponse {
	return DefaultResponse{
		Message: ResponseDefault{
			Success:    false,
			Message:    cases.Title(language.Und, cases.NoLower).String(msg),
			StatusCode: strconv.Itoa(status),
		},
	}
}

func DefaultSuccessResponseWithMessage(msg string, status int, data any) DefaultResponse {
	return DefaultResponse{
		Message: ResponseDefault{
			Success:    true,
			Message:    cases.Title(language.Und, cases.NoLower).String(msg),
			StatusCode: strconv.Itoa(status),
		},
		Data: data,
	}
}
