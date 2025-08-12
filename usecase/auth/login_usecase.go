package auth

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
	"your.module/name/domain/dto"
	"your.module/name/domain/repository"
	"your.module/name/domain/service"
)

type LoginUsecase struct {
	Users      repository.UserRepository
	Sessions   repository.SessionRepository
	Token      service.TokenService
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}

func NewLoginUsecase(users repository.UserRepository, sessions repository.SessionRepository, token service.TokenService, accessTTL, refreshTTL time.Duration) *LoginUsecase {
	return &LoginUsecase{Users: users, Sessions: sessions, Token: token, AccessTTL: accessTTL, RefreshTTL: refreshTTL}
}

var ErrInvalidCredentials = errors.New("invalid email or password")

func (uc *LoginUsecase) Execute(ctx context.Context, in dto.LoginRequest, ua, ip *string) (*dto.LoginResponse, error) {
	u, err := uc.Users.FindByEmail(ctx, in.Email)
	if err != nil || u == nil {
		return nil, ErrInvalidCredentials
	}
	if bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.Password)) != nil {
		return nil, ErrInvalidCredentials
	}

	// access token (JWT)
	access, err := uc.Token.SignAccessToken(u.ID, uc.AccessTTL)
	if err != nil {
		return nil, err
	}

	// refresh token (opaque, ngẫu nhiên)
	buf := make([]byte, 32)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	refresh := base64.RawURLEncoding.EncodeToString(buf)
	h := sha256.Sum256([]byte(refresh))
	hash := base64.RawURLEncoding.EncodeToString(h[:])

	// lưu session
	expires := time.Now().Add(uc.RefreshTTL)
	if _, err := uc.Sessions.Create(ctx, &repository.Session{
		UserID:           u.ID,
		RefreshTokenHash: hash,
		UserAgent:        ua,
		IP:               ip,
		ExpiresAt:        expires,
	}); err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil

	log.Printf("[DEBUG] refresh_plain=%s", refresh)
	log.Printf("[DEBUG] refresh_hash_saved=%s", hash)
}
