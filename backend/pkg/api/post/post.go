package socialnetwork

import (
	"encoding/json"
	"net/http"

	h "socialnetwork/pkg/api/_http"
	token "socialnetwork/pkg/api/jwt"
	model "socialnetwork/pkg/structs"
	request "socialnetwork/pkg/db/requests"
	utilities "socialnetwork/pkg/api/utils"
)

func PostHandler(w http.ResponseWriter, r *http.Request) {

	groupId := r.URL.Query().Get("id")
	postReturn := model.PostReturn{}

	client, err := token.GetUserIdFromToken(r)
	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	if r.Method == "POST" {

		err := r.ParseMultipartForm(10 << 20) // 10 MB
        if err != nil {
            h.Http400(w, err.Error())
            return
        }

        err = json.Unmarshal([]byte(r.FormValue("json")), &postReturn)
        if err != nil {
            h.Http400(w, err.Error())
            return
        }

		if len(postReturn.GroupId) > 0 {
			groupId = postReturn.GroupId
		}

		postReturn.ImgFile, _, err = r.FormFile("file")
		if err != nil {
			postReturn.ImgFile = nil
		}
		if postReturn.ImgFile != nil {
			postReturn.ImgName, err = utilities.Uploadimage(postReturn.ImgFile, "./pkg/db/img")
			if err != nil {
				h.Http401(w, map[string]string{"image": "wrong format or size image"})
				return
			}
		}

		err = request.PostInserter(postReturn, groupId, client)
		if err != nil {
			//remove picture if error
			h.Http400(w, err.Error())
			return
		}

	}

	postList, err := request.PostFetcher(groupId, client)

	if err != nil {
		h.Http400(w, err.Error())
		return
	}

	json.NewEncoder(w).Encode(postList)
}
