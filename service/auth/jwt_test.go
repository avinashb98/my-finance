package auth_test

import (
	"github.com/avinashb98/myfin/config"
	"github.com/avinashb98/myfin/repository/user"
	"github.com/avinashb98/myfin/service/auth"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_GenerateToken(t *testing.T) {
	var userRepo user.Repository
	secret := "test_secret"
	issuer := "myfin"

	conf := config.JWT{
		Secret:                    secret,
		Issuer:                    issuer,
		JWTUserTokenExpiryInHours: 48,
	}

	service := auth.NewService(conf, userRepo)

	token, err := service.GenerateToken("test_handle", true)
	assert.Empty(t, err)
	assert.NotEmpty(t, token)

	token, err = service.GenerateToken("", true)
	assert.Empty(t, err)
	assert.NotEmpty(t, token)

	token, err = service.GenerateToken("", false)
	assert.Empty(t, err)
	assert.NotEmpty(t, token)
}

func TestService_ValidateToken(t *testing.T) {
	var userRepo user.Repository
	var _token string
	var token *jwt.Token
	var err error
	secret := "test_secret"
	issuer := "myfin"

	conf := config.JWT{
		Secret:                    secret,
		Issuer:                    issuer,
		JWTUserTokenExpiryInHours: 48,
	}
	service := auth.NewService(conf, userRepo)
	_, err = service.ValidateToken("")
	assert.NotEmpty(t, err)

	_token, _ = service.GenerateToken("test_handle", true)
	token, err = service.ValidateToken(_token)
	assert.Empty(t, err)
	assert.Equal(t, true, token.Valid)
	assert.Equal(t, "test_handle", token.Claims.(*auth.Payload).Handle)
	assert.Equal(t, "myfin", token.Claims.(*auth.Payload).Issuer)

	conf = config.JWT{
		Secret:                    secret,
		Issuer:                    issuer,
		JWTUserTokenExpiryInHours: -2,
	}
	service = auth.NewService(conf, userRepo)
	_token, err = service.GenerateToken("test_handle", true)
	assert.Empty(t, err)
	token, err = service.ValidateToken(_token)
	assert.NotEmpty(t, err)
	assert.Equal(t, false, token.Valid)
}
