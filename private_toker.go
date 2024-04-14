package toker

import (
	"crypto/rsa"
	"encoding/json"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type PrivateToker struct {
	private *rsa.PrivateKey
}

func NewPrivateToker(privatePem []byte) (*PrivateToker, error) {
	private, err := jwt.ParseRSAPrivateKeyFromPEM(privatePem)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &PrivateToker{
		private: private,
	}, nil
}

func (t *PrivateToker) Generate(expired time.Time, data any) (string, error) {
	content, err := json.Marshal(data)
	if err != nil {
		return "", errors.WithStack(err)
	}
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, &claim{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "",
			Subject:   "",
			Audience:  nil,
			ExpiresAt: jwt.NewNumericDate(expired),
			NotBefore: nil,
			IssuedAt:  jwt.NewNumericDate(now),
			ID:        "",
		},
		Payload: content,
	})
	signedString, err := token.SignedString(t.private)
	if err != nil {
		return "", errors.WithStack(err)
	}
	return signedString, nil
}
