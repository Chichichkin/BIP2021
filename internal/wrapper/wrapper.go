package wrapper

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"polybay/internal/responses"
)

func ErrorResponse(w http.ResponseWriter, err error) {
	raw, err := json.Marshal(responses.Error{Message: err.Error()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Error(w, string(raw), http.StatusBadRequest)
}

func SuccessResponse(w http.ResponseWriter, response *http.Response) {
	raw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	http.Error(w, string(raw), response.StatusCode)
}
