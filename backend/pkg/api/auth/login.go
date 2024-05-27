package socialnetwork

import (
	"encoding/json"
	"net/http"

	h "socialnetwork/pkg/api/_http"
	token "socialnetwork/pkg/api/jwt"
	ut "socialnetwork/pkg/api/utils"
	req "socialnetwork/pkg/db/requests"
	st "socialnetwork/pkg/structs"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	request := st.LoginRequest{}
	// Decode the json request into the request struct
	err := json.NewDecoder(r.Body).Decode(&request)
	// Client error
	if err != nil {
		h.Http400(w, err.Error())
		return
	}
	// Send the email to get the user
	user, err := req.GetUserByEmail(request.Email)
	if err != nil {
		h.Http401(w, map[string]string{"email": "Wrong email"})
		return
	}
	// compare password
	if ut.CheckPasswordHash(request.Password, user.Password) {
		// Generate and sign the token
		jwt, err := token.CreateToken(int(user.ID), user.Email, user.NickName)
		if err != nil {
			h.Http500(w, err.Error())
			return
		}
		// Response to the client
		h.Http200(w, st.LoginResponse{Token: jwt})
		return

	}
	h.Http401(w, map[string]string{"password": "Wrong password"})
}
