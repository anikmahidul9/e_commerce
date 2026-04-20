package util

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type JwtHeader struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type JwtPayload struct {
	Sub         string `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJwt(secret string, data JwtPayload) (string, error) {
	header := JwtHeader{
		Alg: "HS256",
		Typ: "JWT",
	}
	//

	byteArrHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	headerBase64 := base64UrlEncode(byteArrHeader)

	byteArrData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	payloadBase64 := base64UrlEncode(byteArrData)

	message := headerBase64 + "." + payloadBase64

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)
	signature := h.Sum(nil)
	signatureBase64 := base64UrlEncode(signature)

	jwt := message + "." + signatureBase64
	return jwt, nil
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
