package toker

import "github.com/golang-jwt/jwt/v5"

type claim struct {
	jwt.RegisteredClaims
	Payload []byte `json:"payload,omitempty"`
}
