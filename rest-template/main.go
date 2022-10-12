package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_validation/response/"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/rest_validation/test/not_found", defaultValidation).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func defaultErrorHandler(response http.ResponseWriter, request *http.Request) {
	handler := ErrorStructResponse{TimeStamp: "2006-02-01"}
	handler.Exception(response, request, http.StatusConflict, nil)
}

type Address struct {
	Street string `validate:"required,max=4,min=1"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func defaultValidation(response http.ResponseWriter, request *http.Request) {

	var validate *validator.Validate
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	validate = validator.New()
	err := validate.Struct(address)

	handler := ErrorStructResponse{}
	handler.Exception(response, request, http.StatusOK, err)
}
