package socialnetwork

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	request "socialnetwork/pkg/db/requests"

	"github.com/golang-jwt/jwt/v5"
)

// JWT creation, can add as much stuff as possible in the claims
func CreateToken(userid int, email, nickname string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"id":       userid,
			"email":    email,
			"nickname": nickname,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})
	tokenString, err := token.SignedString([]byte(os.Getenv("SecretKey")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// check the token and its key

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SecretKey")), nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// retrieve claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("JWT claims retrieval error")
	}

	// check token expiration
	expirationTime := time.Unix(int64(claims["exp"].(float64)), 0)
	if time.Now().After(expirationTime) {
		return nil, fmt.Errorf("token Expired")
	}

	//check user exist
	tmp := fmt.Sprintf("%v", claims["nickname"])
	if request.CheckNickname(tmp) {
		return nil, fmt.Errorf("user doesn't exist")
	}

	return claims, nil
}

func GetUserIdFromToken(r *http.Request) (int, error) {
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]

	claims, _ := VerifyToken(tokenString)
	x := fmt.Sprintf("%v", claims["id"])
	userId, err := strconv.Atoi(x)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func GetNicknameFromToken(r *http.Request) string {
	tokenString := strings.Split(r.Header.Get("Authorization"), " ")[1]

	claims, _ := VerifyToken(tokenString)
	x := fmt.Sprintf("%v", claims["nickname"])

	return x
}
