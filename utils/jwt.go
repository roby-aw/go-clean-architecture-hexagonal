package utils

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	jose "github.com/dvsekhvalnov/jose2go"
)

type JwtTokenClaims struct {
	Sub  string `json:"sub"`
	Name string `json:"name"`
	Exp  int64  `json:"exp"`
	Role string `json:"role"`
}

type Token struct {
	AccessToken         string `json:"access_token"`
	AccessTokenExpired  int    `json:"access_token_expired"`
	RefreshToken        string `json:"refresh_token"`
	RefreshTokenExpired int    `json:"refresh_token_expired"`
}
type JwtTokenClaimsUser struct {
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	Exp           int64  `json:"exp"`
	Role          string `json:"role"`
	Authorization bool   `json:"authorization"`
}

type RefreshJwtTokenClaimsUser struct {
	Sub                  string `json:"sub"`
	Email                string `json:"email"`
	Exp                  int64  `json:"exp"`
	Role                 string `json:"role"`
	AuthorizationRefresh bool   `json:"authorization_refresh"`
}

func GenerateAccessTokenUser(id, email string) (int, string, error) {
	secret1 := os.Getenv("JWT_SECRET")
	expired := 7200
	claims := &JwtTokenClaimsUser{
		Sub:   id,
		Email: email,
		Exp:   time.Now().Add(time.Duration(expired) * time.Second).Unix(),
		Role:  "User",
	}
	key, err := Decode(secret1)
	if err != nil {
		return 0, "", err
	}
	e, err := json.Marshal(claims)
	if err != nil {
		return 0, "", err
	}
	str, err := jose.Sign(string(e), jose.HS256, key, jose.Header("typ", "JWT"))
	if err != nil {
		return 0, "", err
	}
	return expired, str, err
}

func GenerateRefreshTokenUser(id, email string) (int, string, error) {
	secret1 := os.Getenv("JWT_SECRET")
	expired := 14400
	claims := &RefreshJwtTokenClaimsUser{
		Sub:                  id,
		Email:                email,
		Exp:                  time.Now().Add(time.Duration(expired) * time.Second).Unix(),
		Role:                 "User",
		AuthorizationRefresh: true,
	}
	key, err := Decode(secret1)
	if err != nil {
		return 0, "", err
	}
	e, err := json.Marshal(claims)
	if err != nil {
		return 0, "", err
	}
	str, err := jose.Sign(string(e), jose.HS256, key, jose.Header("typ", "JWT"))
	if err != nil {
		return 0, "", err
	}
	return expired, str, err
}

func Decode(data string) ([]byte, error) {
	test := []byte(data)
	hashbyte := binary.BigEndian.Uint64(test)
	str := fmt.Sprintf("%d", hashbyte)
	str = strings.Replace(str, "-", "+", -1) // 62nd char of encoding
	str = strings.Replace(str, "_", "/", -1) // 63rd char of encoding

	switch len(str) % 4 { // Pad with trailing '='s
	case 0: // no padding
	case 2:
		str += "==" // 2 pad chars
	case 3:
		str += "=" // 1 pad char
	}

	return base64.StdEncoding.DecodeString(str)
}

func Encode(data []byte) string {
	result := base64.StdEncoding.EncodeToString(data)
	result = strings.Replace(result, "+", "-", -1) // 62nd char of encoding
	result = strings.Replace(result, "/", "_", -1) // 63rd char of encoding
	result = strings.Replace(result, "=", "", -1)  // Remove any trailing '='s

	return result
}

func DecodeToken(token string) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	key, err := Decode(secret)
	if err != nil {
		return "", err
	}

	payload, _, err := jose.Decode(token, key)
	if err != nil {
		return "", err
	}
	claims := &JwtTokenClaims{}
	err = json.Unmarshal([]byte(payload), claims)
	if err != nil {
		return "", err
	}
	if claims.Exp < time.Now().Unix() {
		return "", fmt.Errorf("Token expired")
	}
	return claims.Sub, nil
}
