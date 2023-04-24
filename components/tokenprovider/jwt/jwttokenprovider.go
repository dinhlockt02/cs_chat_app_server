package jwt

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/tokenprovider"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type jwtTokenProvider struct {
	secret string
}

type AppClaims struct {
	jwt.RegisteredClaims
	Payload *tokenprovider.TokenPayload `json:"payload"`
}

func (j *jwtTokenProvider) Generate(payload *tokenprovider.TokenPayload, expiry int) (*tokenprovider.Token, error) {
	now := time.Now()
	expire := now.Add(time.Second * time.Duration(expiry))
	claims := AppClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			IssuedAt:  jwt.NewNumericDate(now),
		},
		Payload: payload,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	return &tokenprovider.Token{
		Token:     tokenString,
		CreatedAt: &now,
		ExpiredAt: &expire,
	}, nil
}

func (j *jwtTokenProvider) Validate(token string) (*tokenprovider.TokenPayload, error) {
	tk, err := jwt.ParseWithClaims(token, &AppClaims{Payload: &tokenprovider.TokenPayload{}}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	if claims, ok := tk.Claims.(*AppClaims); ok && tk.Valid {
		return claims.Payload, nil
	} else {
		return nil, common.ErrInvalidRequest(nil)
	}
}

func NewJwtTokenProvider(secret string) *jwtTokenProvider {
	return &jwtTokenProvider{secret: secret}
}
