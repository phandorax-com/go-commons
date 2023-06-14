package resttemplate

import (
	"encoding/json"
	"net/http"
)

type structResponse struct {
	responseWrite http.ResponseWriter
}

func (modelResponse *structResponse) createResponse(responseWrite http.ResponseWriter) structResponse {
	responseWrite.Header().Set("content-type", "application/json")
	return structResponse{
		responseWrite: responseWrite,
	}
}

func (modelResponse *structResponse) generate(httpCode int, value interface{}) {
	modelResponse.responseWrite.WriteHeader(httpCode)
	json.NewEncoder(modelResponse.responseWrite).Encode(value)
}

func (modelResponse *structResponse) RestResponse(responseWrite http.ResponseWriter, value interface{}, httpCode int) {
	response := modelResponse.createResponse(responseWrite)
	response.generate(httpCode, value)
}

func (errorModelResponse *ErrorStructResponse) BadRequestException(request *http.Request, code string, detailsErrors error) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusBadRequest, code, detailsErrors)
}
func (errorModelResponse *ErrorStructResponse) PaymentRequiredException(request *http.Request, code string) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusPaymentRequired, code, nil)
}
func (errorModelResponse *ErrorStructResponse) ForbiddenException(request *http.Request, code string) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusForbidden, code, nil)
}
func (errorModelResponse *ErrorStructResponse) NotFoundException(request *http.Request, code string) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusNotFound, code, nil)
}
func (errorModelResponse *ErrorStructResponse) MethodNotAllowedException(request *http.Request, code string) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusMethodNotAllowed, code, nil)
}
func (errorModelResponse *ErrorStructResponse) NotAcceptableException(request *http.Request, code string) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusNotAcceptable, code, nil)
}
func (errorModelResponse *ErrorStructResponse) ConflictException(request *http.Request, code string, detailsErrors error) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusConflict, code, detailsErrors)
}
func (errorModelResponse *ErrorStructResponse) PreconditionFailedException(request *http.Request, code string) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusPreconditionFailed, code, nil)
}
func (errorModelResponse *ErrorStructResponse) UnsuportedMediaTypeException(request *http.Request, code string) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusUnsupportedMediaType, code, nil)
}
func (errorModelResponse *ErrorStructResponse) UnprocessableEntityException(request *http.Request, code string) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusUnprocessableEntity, code, nil)
}
func (errorModelResponse *ErrorStructResponse) InternalServerErrorException(request *http.Request, code string) {
	errorModelResponse.Exception(errorModelResponse.responseWrite, request, http.StatusInternalServerError, code, nil)
}
