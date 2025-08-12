package auth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"your.module/name/domain/dto"
	"your.module/name/domain/repository"
	"your.module/name/domain/service"
)

var ErrInvalidRefresh = errors.New("invalid refresh token")

type RefreshTokenUsecase struct {
	Sessions         repository.SessionRepository
	Token            service.TokenService
	AccessTTLSeconds int64 // hoặc dùng time.Duration tuỳ bạn
}

func NewRefreshTokenUsecase(sessions repository.SessionRepository, token service.TokenService, accessTTLSeconds int64) *RefreshTokenUsecase {
	return &RefreshTokenUsecase{Sessions: sessions, Token: token, AccessTTLSeconds: accessTTLSeconds}
}

func (uc *RefreshTokenUsecase) Execute(ctx context.Context, in dto.RefreshRequest) (*dto.RefreshResponse, error) {
	// Chuẩn hoá: bỏ space/newline vô tình dính vào
	rt := strings.TrimSpace(in.RefreshToken) // <-- thêm

	// hash đúng chuẩn (URL-safe, no padding)
	h := sha256.Sum256([]byte(rt))
	hash := base64.RawURLEncoding.EncodeToString(h[:])

	s, err := uc.Sessions.FindActiveByHash(ctx, hash)
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, ErrInvalidRefresh
	}

	dur := time.Duration(uc.AccessTTLSeconds) * time.Second
	access, err := uc.Token.SignAccessToken(s.UserID, dur)
	if err != nil {
		return nil, err
	}
	return &dto.RefreshResponse{AccessToken: access}, nil
}
