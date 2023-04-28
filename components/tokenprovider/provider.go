package tokenprovider

import "time"

type TokenProvider interface {
	Generate(payload *TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

type Token struct {
	Token     string     `json:"token"`
	CreatedAt *time.Time `json:"createdAt"`
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
}

type TokenPayload struct {
	Id string `json:"id"`
}
