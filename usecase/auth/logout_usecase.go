package auth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"

	"your.module/name/domain/dto"
	"your.module/name/domain/repository"
)

type LogoutUsecase struct {
	Sessions repository.SessionRepository
}

func NewLogoutUsecase(sessions repository.SessionRepository) *LogoutUsecase {
	return &LogoutUsecase{Sessions: sessions}
}

func (uc *LogoutUsecase) Execute(ctx context.Context, in dto.LogoutRequest) error {
	h := sha256.Sum256([]byte(in.RefreshToken))
	hash := base64.RawURLEncoding.EncodeToString(h[:])
	return uc.Sessions.RevokeByHash(ctx, hash)
}
