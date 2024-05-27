package socialnetwork

import (
	"net/http"

	h "socialnetwork/pkg/api/_http"
	ut "socialnetwork/pkg/api/utils"
	req "socialnetwork/pkg/db/requests"
	st "socialnetwork/pkg/structs"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// utils
	var err error
	var errmessage interface{}

	// verif request value
	user, errmessage := VerifRegisterValue(r)
	if errmessage != nil {
		h.Http401(w, errmessage)
		return
	}

	// verif and upload ipage
	if user.AvatarFile != nil {
		user.AvatarName, err = ut.Uploadimage(user.AvatarFile, "./pkg/db/avatars")
		if err != nil {
			h.Http401(w, map[string]string{"avatar": "wrong format or size image"})
			return
		}
	}

	// hash password
	user.Password, err = ut.HashPassword(user.Password)
	if err != nil {
		h.Http500(w, err.Error())
		return
	}
	// insert user in DB
	err = req.InsertUser(user)
	if err != nil {
		h.Http500(w, err.Error())
		return
	}
	// return succes message
	h.Http200(w, map[string]string{"message": "succes register!"})
}

func VerifRegisterValue(r *http.Request) (st.SingUpRequest, interface{}) {
	// utils
	user := st.SingUpRequest{}
	var err error

	// verif email
	user.Email = r.FormValue("email")
	if user.Email == "" {
		return user, map[string]string{"email": "email invalid"}
	}
	if !req.CheckEmail(user.Email) {
		return user, map[string]string{"email": "email already exist"}
	}

	// verif  firstname
	user.FirstName = r.FormValue("firstName")
	if user.FirstName == "" {
		return user, map[string]string{"fristname": "firstname invalid"}
	}

	// verif lastname
	user.LastName = r.FormValue("lastName")
	if user.LastName == "" {
		return user, map[string]string{"lastname": "lastname invalid"}
	}

	// verif date of birth
	user.DateOfBirth = r.FormValue("dateOfBirth")
	if user.DateOfBirth == "" {
		return user, map[string]string{"dateofbirth": "date of birth invalid"}
	}

	// verif nickanme
	user.NickName = r.FormValue("nickname")
	if user.NickName == "" {
		return user, map[string]string{"nickname": "nickname invalid"}
	}
	if !req.CheckNickname(user.NickName) {
		return user, map[string]string{"nickname": "nickname already exist"}
	}

	// verif about me
	user.AboutMe = r.FormValue("aboutMe")

	// verif password
	user.Password = r.FormValue("password")
	if user.Password == "" {
		return user, map[string]string{"password": "password invalid"}
	}

	// verif avatar
	user.AvatarFile, _, err = r.FormFile("avatar")
	if err != nil {
		user.AvatarFile = nil
	}

	// return valid userrequest
	return user, nil
}
