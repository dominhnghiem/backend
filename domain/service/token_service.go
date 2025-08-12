package service

import "time"

type TokenService interface {
	SignAccessToken(userID int64, ttl time.Duration) (string, error)
}
