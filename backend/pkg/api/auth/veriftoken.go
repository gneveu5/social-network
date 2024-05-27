package socialnetwork

import (
	"net/http"

	h "socialnetwork/pkg/api/_http"
)

func VeriftokenHandler(w http.ResponseWriter, r *http.Request) {
	h.Http200(w, "token is valid")
}
