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

func (modelResponse *StructResponse) generate(httpCode int) {
	modelResponse.responseWrite.WriteHeader(httpCode)
	json, _ := json.Marshal(&modelResponse)
	modelResponse.responseWrite.Write(json)
}

func (modelResponse *StructResponse) RestResponse(responseWrite http.ResponseWriter, httpCode int) {
	response := modelResponse.createResponse(responseWrite)
	response.generate(httpCode)
}
