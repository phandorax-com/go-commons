package handler

import (
	"encoding/json"
	"net/http"
	"os/exec"
	"time"

	"github.com/go-playground/validator/v10"
)

type ErrorModelResponse struct {
	Details       []detail `json:"details"`
	Path          string   `json:"path"`
	TimeStamp     string   `json:"time_stamp"`
	ResponseWrite http.ResponseWriter
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

func (errorModelResponse *ErrorModelResponse) createResponse(responseWrite http.ResponseWriter, request *http.Request) ErrorModelResponse {
	if len(errorModelResponse.TimeStamp) == 0 {
		errorModelResponse.TimeStamp = "2006-01-02T15:04:05.000000000"
	}
	responseWrite.Header().Set("content-type", "application/json")
	uuid, _ := exec.Command("uuidgen").Output()
	responseWrite.Header().Set("Trace-UUID", string(uuid))
	return ErrorModelResponse{
		Details:       errorModelResponse.Details,
		Path:          request.RequestURI,
		TimeStamp:     time.Now().UTC().Format(errorModelResponse.TimeStamp),
		ResponseWrite: responseWrite,
	}
}

func (errorModelResponse *ErrorModelResponse) generate(httpCode int) {
	errorModelResponse.ResponseWrite.WriteHeader(httpCode)
	json, _ := json.Marshal(&errorModelResponse)
	errorModelResponse.ResponseWrite.Write(json)
}

func (errorModelResponse *ErrorModelResponse) Exception(responseWrite http.ResponseWriter, request *http.Request, httpCode int, detailsErrors error) {

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
