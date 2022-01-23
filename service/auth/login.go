package auth

import (
	"context"
	"github.com/avinashb98/myfin/utils"
)

func (s *service) IsAuthenticated(ctx context.Context, handle, password string) (bool, error) {
	userAuth, err := s.userRepo.GetUserAuthByHandle(ctx, handle)
	if err != nil {
		return false, err
	}
	return utils.CheckPasswordHash(password, userAuth.PasswordHash), nil
}
