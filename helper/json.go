package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, result interface{}) {
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(result)
	PanicIfError(err)
}
