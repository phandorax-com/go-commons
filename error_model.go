package resttemplate

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

type ErrorStructResponse struct {
	Details       []detail `json:"details"`
	Path          string   `json:"path"`
	TimeStamp     string   `json:"time_stamp"`
	responseWrite http.ResponseWriter
}

type detail struct {
	ErrorMessage string   `json:"error_message"`
	MetaData     metaData `json:"meta_data"`
}

type metaData struct {
	Type  string      `json:"type"`
	Field string      `json:"fields"`
	Rule  string      `json:"rule"`
	Value interface{} `json:"value"`
}

func (errorModelResponse *ErrorStructResponse) createResponse(responseWrite http.ResponseWriter, request *http.Request) ErrorStructResponse {
	if len(errorModelResponse.TimeStamp) == 0 {
		errorModelResponse.TimeStamp = "2006-01-02T15:04:05.000000000"
	}
	responseWrite.Header().Set("content-type", "application/json")
	return ErrorStructResponse{
		Details:       errorModelResponse.Details,
		Path:          request.RequestURI,
		TimeStamp:     time.Now().UTC().Format(errorModelResponse.TimeStamp),
		responseWrite: responseWrite,
	}
}

func (errorModelResponse *ErrorStructResponse) generate(httpCode int) {
	errorModelResponse.responseWrite.WriteHeader(httpCode)
	json, _ := json.Marshal(&errorModelResponse)
	errorModelResponse.responseWrite.Write(json)
}

func (errorModelResponse *ErrorStructResponse) Exception(responseWrite http.ResponseWriter, request *http.Request, httpCode int, detailsErrors error) {

	if detailsErrors != nil {
		for _, err := range detailsErrors.(validator.ValidationErrors) {
			metaData := metaData{
				Type:  err.Kind().String(),
				Field: err.Field(),
				Rule:  err.Tag(),
				Value: err.Value(),
			}
			detailError := detail{
				ErrorMessage: err.Error(),
				MetaData:     metaData,
			}
			errorModelResponse.Details = append(errorModelResponse.Details, detailError)
		}
	} else {
		errorModelResponse.Details = []detail{}
	}

	errorModel := errorModelResponse.createResponse(responseWrite, request)
	errorModel.generate(httpCode)
}
