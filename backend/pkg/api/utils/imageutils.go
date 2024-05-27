package socialnetwork

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func Uploadimage(avatarFile multipart.File, path string) (string, error) {
	format, err := Imageformat(avatarFile)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	err = os.MkdirAll(path, 0o777)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	avatarName := uuid.NewString() + format
	avatarPath := path + "/" + avatarName

	avatarFile.Seek(0, 0)
	avatarData, err := io.ReadAll(avatarFile)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	err = os.WriteFile(avatarPath, avatarData, 0o666)
	if err != nil {
		fmt.Print(err)

		return "", err
	}
	return avatarName, nil
}

func Imageformat(imagefile multipart.File) (string, error) {
	buff := make([]byte, 512)
	if _, err := imagefile.Read(buff); err != nil {
		fmt.Println(err)
		return "", err
	}
	imagetype := http.DetectContentType(buff)
	if imagetype == "image/png" {
		return ".png", nil
	} else if imagetype == "image/jpeg" {
		return ".jpeg", nil
	} else if imagetype == "image/gif" {
		return ".gif", nil
	}
	return "", errors.New("wrong image format")
}
