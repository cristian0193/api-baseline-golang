package utils

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

type Claim struct {
	IdUser string `json:"idUser"`
	User   string `json:"user"`
	jwt.StandardClaims
}

//Verifica segun el tipo de token que se envie
func VerifyToken(r string, lenguaje string) (string, int) {
	var description string
	var code int

	publicBytes, err := ioutil.ReadFile(os.Getenv("PUBLIC_RSA"))
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)

	token, err := jwt.Parse(r, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	fmt.Println(token)
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				description = Lenguage(lenguaje, "TOKEN_EXPIRED")
				code = http.StatusNonAuthoritativeInfo
				return description, code
			case jwt.ValidationErrorSignatureInvalid:
				description = Lenguage(lenguaje, "SIGNATURE_MATCH")
				code = http.StatusUnauthorized
				return description, code
			default:
				description = Lenguage(lenguaje, "INVALID_TOKEN")
				code = http.StatusUnauthorized
				return description, code
			}
		default:
			description = Lenguage(lenguaje, "INVALID_TOKEN")
			code = http.StatusUnauthorized
			return description, code
		}
	}
	if token != nil {
		code = http.StatusOK
	}
	return description, code
}

func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func ExtractIdUser(tokenStr string) string {
	publicBytes, err := ioutil.ReadFile(os.Getenv("PUBLIC_RSA"))
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(err)
	return claims["idUser"].(string)
}
