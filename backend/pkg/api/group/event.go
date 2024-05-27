package socialnetwork

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	h "socialnetwork/pkg/api/_http"
	token "socialnetwork/pkg/api/jwt"
	request "socialnetwork/pkg/db/requests"
	model "socialnetwork/pkg/structs"
)

func EventAttendeeFetchHandler(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("id")

	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	claims, err := token.VerifyToken(tokenString)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	if r.Method == "POST" {

		eventRegistration := model.EventRegistrationReturn{}

		err = json.NewDecoder(r.Body).Decode(&eventRegistration)
		if err != nil {
			h.Http400(w, "wrong message request of client")
			return
		}

		// pas beau : enlever la notif (vraiment pas beau)
		userId, _ := token.GetUserIdFromToken(r)
		groupIdInt, _ := strconv.Atoi(groupId)
		request.RemoveNotificationFromDatas(3, userId, groupIdInt, eventRegistration.EventId)

		if eventRegistration.Status == "Register" {
			err = request.EventAttendeesRegister(eventRegistration.EventId, claims)
			if err != nil {
				h.Http400(w, err.Error())
				return
			}
			json.NewEncoder(w).Encode("Unregister")

		} else if eventRegistration.Status == "Unregister" {
			err = request.EventAttendeesUnregister(eventRegistration.EventId, claims)
			if err != nil {
				h.Http400(w, err.Error())
				return
			}
			json.NewEncoder(w).Encode("Register")

		} else {
			_, err := request.EventAttendeesStatus(eventRegistration.EventId, claims)
			if err != nil {
				if err.Error() != "sql: no rows in result set" {
					h.Http400(w, err.Error())
					return
				} else if err.Error() == "sql: no rows in result set" {
					json.NewEncoder(w).Encode("Register")
				}
			} else {
				json.NewEncoder(w).Encode("Unregister")
			}
		}

	} else {
		eventList, err := request.EventAttendeesFetcher(groupId)
		if err != nil {
			h.Http400(w, err.Error())
			return
		}
		json.NewEncoder(w).Encode(eventList)
	}
}
