package socialnetwork

import (
	"encoding/json"
	"net/http"
	"strings"

	h "socialnetwork/pkg/api/_http"
	token "socialnetwork/pkg/api/jwt"
	ws "socialnetwork/pkg/api/websocket"
	request "socialnetwork/pkg/db/requests"
	model "socialnetwork/pkg/structs"

	"github.com/golang-jwt/jwt/v5"
)

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("id")
	eventReturn := model.EventReturn{}

	userId, err := token.GetUserIdFromToken(r)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	if r.Method == "POST" {

		err = json.NewDecoder(r.Body).Decode(&eventReturn)
		if err != nil {
			h.Http400(w, "wrong message request of client")
			return
		}

		groupId = eventReturn.GroupId

		groupMembers, err := request.EventInserter(eventReturn, userId)
		if err != nil {
			h.Http400(w, err.Error())
			return
		}
		for _, v := range groupMembers {
			err = ws.SendNewNotification(v.UserId)
			if err != nil {
				h.Http400(w, err.Error())
				return
			}
		}
	}

	eventList, err := request.EventFetcher(groupId)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(eventList)
}

func GroupListHandler(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("id")

	_, err := token.GetUserIdFromToken(r)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	groupList, err := request.GroupList(groupId)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(groupList)
}

func GroupCreationHandler(w http.ResponseWriter, r *http.Request) {
	var claims jwt.MapClaims
	var err error
	groupCreationReturn := model.GroupCreationReturn{}

	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	claims, err = token.VerifyToken(tokenString)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	err = json.NewDecoder(r.Body).Decode(&groupCreationReturn)
	if err != nil {
		h.Http400(w, "wrong message request of client")
		return
	}

	groupId, err := request.GroupCreation(groupCreationReturn, claims["id"])
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(groupId)
}

func GroupAskMembershipHandler(w http.ResponseWriter, r *http.Request) {

	groupAskMembershipReturn := model.GroupAskMembershipReturn{}

	userId, err := token.GetUserIdFromToken(r)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	err = json.NewDecoder(r.Body).Decode(&groupAskMembershipReturn)
	if err != nil {
		h.Http400(w, "wrong message request of client")
		return
	}

	err = request.GroupAskRegister(userId, groupAskMembershipReturn)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	adminId, err := request.GroupAdminFetcher(groupAskMembershipReturn.GroupId)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	err = ws.SendNewNotification(adminId)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(true)
}

func GroupInviteMembershipHandler(w http.ResponseWriter, r *http.Request) {

	groupInviteMembershipReturn := model.GroupInviteMembershipReturn{}

	userId, err := token.GetUserIdFromToken(r)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	err = json.NewDecoder(r.Body).Decode(&groupInviteMembershipReturn)
	if err != nil {
		h.Http400(w, "wrong message request of client")
		return
	}

	err = request.GroupInviteRegister(userId, groupInviteMembershipReturn)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	err = ws.SendNewNotification(groupInviteMembershipReturn.Target)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}
}

func GroupFetchAskMembershipHandler(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("id")

	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	claims, err := token.VerifyToken(tokenString)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	exist, err := request.GroupFetchAskRegister(claims["id"], groupId)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(exist)
}

func GroupFetchExistHandler(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("id")

	group, err := request.GroupFetcher(groupId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(group)
}

func GroupFetchMemberHandler(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("id")

	var claims jwt.MapClaims
	var err error

	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]
	claims, err = token.VerifyToken(tokenString)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	exist, err := request.GroupMemberFetcher(claims, groupId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(exist)
}

func GroupFetchMembersHandler(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("id")

	client, err := token.GetUserIdFromToken(r)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	members, err := request.GroupMembersFetcher(client, groupId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(members)
}
