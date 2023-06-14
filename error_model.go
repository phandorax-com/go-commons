package resttemplate

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"time"
)

type ErrorStructResponse struct {
	Path          string   `json:"path"`
	Code          string   `json:"internal_code"`
	TimeStamp     string   `json:"time_stamp"`
	Errors        []errors `json:"errors"`
	responseWrite http.ResponseWriter
}

type errors struct {
	Message  string   `json:"message"`
	MetaData metaData `json:"meta_data"`
}

type metaData struct {
	Type  string      `json:"type"`
	Field string      `json:"field"`
	Rule  string      `json:"rule"`
	Value interface{} `json:"value"`
}

func NewErrorResponse(response http.ResponseWriter) *ErrorStructResponse {
	return &ErrorStructResponse{
		responseWrite: response,
	}
}
func (errorModelResponse *ErrorStructResponse) createResponse(responseWrite http.ResponseWriter, request *http.Request, code string) ErrorStructResponse {

	responseWrite.Header().Set("content-type", "application/json")
	return ErrorStructResponse{
		Errors:        errorModelResponse.Errors,
		Code:          code,
		Path:          request.RequestURI,
		TimeStamp:     time.Now().UTC().Format("2006-01-02T15:04:05.000000000"),
		responseWrite: responseWrite,
	}
}

func (errorModelResponse *ErrorStructResponse) generate(httpCode int) {
	errorModelResponse.responseWrite.WriteHeader(httpCode)
	json, _ := json.Marshal(&errorModelResponse)
	errorModelResponse.responseWrite.Write(json)
}

func (errorModelResponse *ErrorStructResponse) Exception(responseWrite http.ResponseWriter, request *http.Request, httpCode int, code string, detailsErrors error) {

	errorModelResponse.Errors = []errors{}
	if detailsErrors != nil {
		for _, err := range detailsErrors.(validator.ValidationErrors) {

			message := ""
			if len(strings.Split(err.Error(), ":")) >= 3 {
				message = strings.Split(err.Error(), ":")[2]
			}

			metaData := metaData{
				Type:  err.Kind().String(),
				Field: err.Field(),
				Rule:  err.Tag(),
				Value: err.Value(),
			}
			detailError := errors{
				Message:  message,
				MetaData: metaData,
			}
			errorModelResponse.Errors = append(errorModelResponse.Errors, detailError)
		}
	}

	errorModel := errorModelResponse.createResponse(responseWrite, request, code)
	errorModel.generate(httpCode)
}
