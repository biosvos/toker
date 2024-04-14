package toker

import "github.com/pkg/errors"

var (
	ErrExpiredToken = errors.New("expired token")
)
