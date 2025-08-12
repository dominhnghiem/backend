package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type HS256TokenService struct {
	Secret []byte
}

func NewHS256TokenService(secret string) *HS256TokenService {
	return &HS256TokenService{Secret: []byte(secret)}
}

func (s *HS256TokenService) SignAccessToken(userID int64, ttl time.Duration) (string, error) {
	now := time.Now()
	claims := jwt.MapClaims{
		"sub": userID,
		"iat": now.Unix(),
		"exp": now.Add(ttl).Unix(),
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(s.Secret)
}
