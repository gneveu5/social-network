package socialnetwork

import (
	"encoding/json"
	"net/http"

	h "socialnetwork/pkg/api/_http"
	tk "socialnetwork/pkg/api/jwt"
	req "socialnetwork/pkg/db/requests"
)

func UserinfoHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	usersend, _ := tk.GetUserIdFromToken(r)

	nicknames, err := req.GetUserInfo(query, usersend)
	if err != nil {
		h.Http400(w, "GetUserNickname() : "+err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Convertir les nicknames en JSON
	nicknamesJSON, err := json.Marshal(nicknames)
	if err != nil {
		h.Http400(w, "SearchHandler() : "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(nicknamesJSON)
}

func PublictoprivateHandler(w http.ResponseWriter, r *http.Request) {
	usersend, _ := tk.GetUserIdFromToken(r)
	err := req.SwapPublictoprivate(usersend)
	if err != nil {
		h.Http400(w, "SwapPublictoprivate() : "+err.Error())
		return
	}
	h.Http200(w, "succes")
}

func PrivatetopublicHandler(w http.ResponseWriter, r *http.Request) {
	usersend, _ := tk.GetUserIdFromToken(r)
	err := req.SwapPrivatetopublic(usersend)
	if err != nil {
		h.Http400(w, "SwapPrivatetopublic() : "+err.Error())
		return
	}
	h.Http200(w, "succes")
}

func NicknameHandler(w http.ResponseWriter, r *http.Request) {
	userid, _ := tk.GetUserIdFromToken(r)
	userinfo, err := req.GetAvatarAndUsernameByID(userid)
	if err != nil {
		h.Http401(w, err)
		return
	}
	h.Http200(w, userinfo)
}
