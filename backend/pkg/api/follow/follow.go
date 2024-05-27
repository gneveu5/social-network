package socialnetwork

import (
	"encoding/json"
	"fmt"
	"net/http"

	h "socialnetwork/pkg/api/_http"
	tk "socialnetwork/pkg/api/jwt"
	ws "socialnetwork/pkg/api/websocket"
	req "socialnetwork/pkg/db/requests"
)

// send follower list
func FollowerHandler(w http.ResponseWriter, r *http.Request) {
	userselected := r.URL.Query().Get("query")
	if userselected == "undefined" {
		userselected = tk.GetNicknameFromToken(r)
	}

	usersend, _ := tk.GetUserIdFromToken(r)
	
	followers, err := req.GetFollowers(userselected, usersend)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")

	// Convertir les nicknames en JSON
	followersJSON, err := json.Marshal(followers)
	if err != nil {
		h.Http400(w, "FollowerHandler() : "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(followersJSON)
}

func FollowerNotInGroupHandler(w http.ResponseWriter, r *http.Request) {
	groupId := r.URL.Query().Get("id")

	userselected := tk.GetNicknameFromToken(r)
	usersend, err := tk.GetUserIdFromToken(r)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}
	
	followers, err := req.GetFollowers(userselected, usersend)
	if err != nil {
		fmt.Println(err)
		h.Http400(w, err.Error())
		return
	}

	followers = req.SortFollowersForGroup(followers, groupId)

	followersJSON, err := json.Marshal(followers)
	if err != nil {
		h.Http400(w, "FollowerHandler() : "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(followersJSON)
}

// send following list
func FollowingHandler(w http.ResponseWriter, r *http.Request) {
	userselected := r.URL.Query().Get("query")
	usersend, _ := tk.GetUserIdFromToken(r)
	following, err := req.GetFollowing(userselected, usersend)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// fmt.Println(following)

	// Convertir les nicknames en JSON
	followingJSON, err := json.Marshal(following)
	if err != nil {
		h.Http400(w, "FollowingHandler() : "+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(followingJSON)
}

// request following
// accepted if public, require confirm if private
func RequestFollowHandler(w http.ResponseWriter, r *http.Request) {
	targetUserName := r.URL.Query().Get("query")
	targetUserId, err := req.GetIdfromNickname(targetUserName)
	if err != nil {
		fmt.Println("RequestFollowHandler : failed getting targetUserId\n" + err.Error())
		return
	}

	askingUserId, _ := tk.GetUserIdFromToken(r)

	isPrivate, err := req.GetPublicPrivatefromId(targetUserId)
	if err != nil {
		fmt.Println("RequestFollowHandler : failed getting publicprivate status\n" + err.Error())
		return
	}

	if !isPrivate {
		err = req.Follow(targetUserName, askingUserId)
		if err != nil {
			h.Http400(w, "Follow() : "+err.Error())
			return
		}
		mutualFollow, err := req.IsFollowerOf(targetUserId, askingUserId)
		if mutualFollow && err == nil {
			err = ws.AddPrivateChat(targetUserId, askingUserId)
			if err != nil {
				fmt.Println("RequestFollowHandler : failed updating private chat\n" + err.Error())
			}
		} else if err != nil {
			fmt.Println("RequestFollowHandler : failed testing mutual follow\n" + err.Error())
		}
		h.Http200(w, "follow ok")
	} else {
		err = req.FollowPrivate(targetUserName, askingUserId)
		if err != nil {
			fmt.Println(err)
		}
		err = req.InsertNotif(targetUserId, 2, askingUserId, 0)
		if err != nil {
			fmt.Println(err)
		}
		ws.SendNewNotification(targetUserId)
		h.Http200(w, "follow ok")
	}
}

func RequnfollowHandler(w http.ResponseWriter, r *http.Request) {
	userselected := r.URL.Query().Get("query")
	usersend, _ := tk.GetUserIdFromToken(r)
	err := req.UnFollow(userselected, usersend)
	if err != nil {
		fmt.Print(err)
		h.Http400(w, "UnFollow() : "+err.Error())
		return
	}

	userselectedId, err := req.GetIdfromNickname(userselected)
	if err != nil {
		fmt.Println("ReqUnfollowHandler : getting id from nickname\n" + err.Error())
	}

	err = ws.RemovePrivateChat(userselectedId, usersend)
	if err != nil {
		fmt.Println("ReqUnfollowHandler : failed removing private chat\n" + err.Error())
	}

	h.Http200(w, "unfollow ok")
}

func ReqCancelFollowPrivateHandler(w http.ResponseWriter, r *http.Request) {
	targetUserName := r.URL.Query().Get("query")
	targetUserId, err := req.GetIdfromNickname(targetUserName)
	askingUserId, _ := tk.GetUserIdFromToken(r)

	if err != nil {
		fmt.Print(err)
		h.Http400(w, "UnFollow() : "+err.Error())
		return
	}

	err = req.CancelFollowPrivate(targetUserName, askingUserId)
	if err != nil {
		fmt.Print(err)
		h.Http400(w, "UnFollow() : "+err.Error())
		return
	}

	err = req.RemoveNotificationFromDatas(2, targetUserId, askingUserId, 0)
	if err != nil {
		fmt.Print(err)
		h.Http400(w, "UnFollow() : "+err.Error())
		return
	}

	h.Http200(w, "unfollow ok")
}
