package utils_test

import (
	"github.com/avinashb98/myfin/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	passwords := []string{
		"thisIsPassword", "",
	}

	for _, password := range passwords {
		passwordHash, err := utils.HashPassword(password, 14)
		assert.Empty(t, err)
		assert.NotEmpty(t, passwordHash)
		assert.Equal(t, len(passwordHash), 60)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	suite := []struct {
		password, hash string
	}{
		{"secret", "$2a$14$IuYesx1flV958YhNayYa4O590R3PXUpxzvELMgOg/gL7eIDs5e8Vi"},
		{"thisIsPassword", "$2a$14$wvlLDN.V.04N4sMK0mSU.u2kk7pqQ04zZAvDQsyxBX1X6PFbSttsW"},
		{"", "$2a$14$Iaqs05N.XMjs3VIXGCkrs.hGMqAD9Id7lUBAh8PfzHl5y3DWrJAX6"},
	}

	for _, testCase := range suite {
		passwordHash, err := utils.HashPassword(testCase.password, 14)
		assert.Empty(t, err)
		assert.Equal(t, utils.CheckPasswordHash(testCase.password, passwordHash), true)
	}
}
