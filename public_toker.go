package toker

import (
	"crypto/rsa"
	"encoding/json"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
)

type PublicToker struct {
	public *rsa.PublicKey
}

func NewPublicToker(publicPem []byte) (*PublicToker, error) {
	public, err := jwt.ParseRSAPublicKeyFromPEM(publicPem)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &PublicToker{
		public: public,
	}, nil
}

func (t *PublicToker) Parse(token string, data any) error {
	parsed, err := t.parseJWT(token)
	if err != nil {
		return err
	}
	claims, ok := parsed.Claims.(jwt.MapClaims)
	if !ok {
		return errors.New("not valid token")
	}
	marshal, err := json.Marshal(claims)
	if err != nil {
		return errors.WithStack(err)
	}
	var ret claim
	err = json.Unmarshal(marshal, &ret)
	if err != nil {
		return errors.WithStack(err)
	}
	err = json.Unmarshal(ret.Payload, &data)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (t *PublicToker) parseJWT(token string) (*jwt.Token, error) {
	parsed, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return t.public, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, errors.WithStack(err)
	}
	return parsed, nil
}
