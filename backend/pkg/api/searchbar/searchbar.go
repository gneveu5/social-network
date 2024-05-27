package socialnetwork

import (
	"net/http"

	h "socialnetwork/pkg/api/_http"
	req "socialnetwork/pkg/db/requests"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")

	nicknames, err := req.GetUserNickname(query)
	if err != nil {
		h.Http400(w, "GetUserNickname() : "+err.Error())
		return
	}

	h.Http200(w, nicknames)
}
