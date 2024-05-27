// fonction pour simplifier les reponses au clients
package socialnetwork

import (
	"encoding/json"
	"net/http"
)

func Http400(w http.ResponseWriter, errormessage string) {
	errorMessageJSON := `{"general": "` + errormessage + `"}`
	http.Error(w, errorMessageJSON, http.StatusBadRequest)
}

func Http500(w http.ResponseWriter, errormessage string) {
	errorMessageJSON := `{"general": "` + errormessage + `"}`
	http.Error(w, errorMessageJSON, http.StatusInternalServerError)
}

func Http200(w http.ResponseWriter, message interface{}) {
	responseJSON, err := json.Marshal(message)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Write(responseJSON)
}

func Http401(w http.ResponseWriter, errorMessage interface{}) {
	jsonResponse, err := json.Marshal(errorMessage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	w.Write(jsonResponse)
}
