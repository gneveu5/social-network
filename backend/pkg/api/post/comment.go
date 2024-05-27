package socialnetwork

import (
	"net/http"
	"encoding/json"
	"strconv"
	"fmt"

	h "socialnetwork/pkg/api/_http"
	token "socialnetwork/pkg/api/jwt"
	model "socialnetwork/pkg/structs"
	request "socialnetwork/pkg/db/requests"
	utilities "socialnetwork/pkg/api/utils"
)

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	
	commentFetch := model.CommentFetchReturn{}

	_, err := token.GetUserIdFromToken(r)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	err = json.NewDecoder(r.Body).Decode(&commentFetch)
	if err != nil {
		h.Http400(w, "wrong message request of client")
		return
	}

	commentList, err := request.CommentFetcher(strconv.Itoa(commentFetch.PostId))
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(commentList)
}

func CommentingHandler(w http.ResponseWriter, r *http.Request) {

	commentPosting := model.CommentPosting{}
	client, err := token.GetUserIdFromToken(r)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	commentPosting.PostId = r.FormValue("postId")
	commentPosting.Message = r.FormValue("message")

	commentPosting.ImgFile, _, err = r.FormFile("picture")
	if err != nil {
		fmt.Println(err)
		commentPosting.ImgFile = nil
	}
	if commentPosting.ImgFile != nil {
		commentPosting.ImgName, err = utilities.Uploadimage(commentPosting.ImgFile, "./pkg/db/img")
		if err != nil {
			h.Http401(w, map[string]string{"image": "wrong format or size image"})
			return
		}
	}

	request.CommentInserter(commentPosting, client)
	commentList, err := request.CommentFetcher(commentPosting.PostId)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(commentList)
}
