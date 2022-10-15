package resttemplate

import (
	"encoding/json"
	"net/http"
)

type StructResponse struct {
	responseWrite http.ResponseWriter
}

func (modelResponse *StructResponse) createResponse(responseWrite http.ResponseWriter) StructResponse {
	responseWrite.Header().Set("content-type", "application/json")
	return StructResponse{
		responseWrite: responseWrite,
	}
}

func (modelResponse *StructResponse) generate(httpCode int, value interface{}) {
	modelResponse.responseWrite.WriteHeader(httpCode)
	json.NewEncoder(modelResponse.responseWrite).Encode(value)
}

func (modelResponse *StructResponse) RestResponse(responseWrite http.ResponseWriter, value interface{}, httpCode int) {
	response := modelResponse.createResponse(responseWrite)
	response.generate(httpCode, value)
}
