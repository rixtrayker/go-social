package auth

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user"] = username
	claims["aud"] = "go-social.mydomain.io"
	claims["iss"] = "mydomain.io"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	newToken, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		err = errors.New("Something went wrong")
		return "", err
	}
	return newToken, nil
}

func CheckJWT(token string) error {
	t, err := jwt.Parse(token, func(tkn *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		return err
	} else if !t.Valid {
		err = errors.New("Invalid token")
		return err
	} else if t.Claims.(jwt.MapClaims)["aud"] != "go-social.mydomain.io" {
		err = errors.New("Invalid aud")
		return err
	} else if t.Claims.(jwt.MapClaims)["iss"] != "mydomain.io" {
		err = errors.New("Invalid iss")
		return err
	}
	return nil
}
