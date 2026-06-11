package autentication

import (
	"errors"
	"fmt"
	"net/http"
	"overtime_system_menagement/src/config"
	"strconv"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(userId uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["autorized"] = true
	permissions["userId"] = userId
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(config.SecretKey))
}

func ExtractUserId(r *http.Request) (uint64, error) {
	tokenString := extract(r)
	token, erro := jwt.Parse(tokenString, returnKeyVerification)
	if erro != nil {
		return 0, erro
	}
	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return userId, nil
	}
	return 0, errors.New("token inválido ")
}

func ValidadeToken(r *http.Request) error {
	tokenString := extract(r)
	token, erro := jwt.Parse(tokenString, returnKeyVerification)
	if erro != nil {
		return erro
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("token inválido")
}

func extract(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnKeyVerification(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("metodo de assinatura inesperado  %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}
