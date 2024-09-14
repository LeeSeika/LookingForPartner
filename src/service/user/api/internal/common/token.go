package common

import (
	"errors"
	"lookingforpartner/pkg/jwtx"
	"time"
)

func CreateTokenAndRefreshToken(uid string, accessExpire, refreshExpire int64, accessSecret string) (string, string, error) {
	now := time.Now().Unix()
	accessToken, err := jwtx.GetToken(accessSecret, now, accessExpire, uid)
	if err != nil {
		return "", "", errors.New("failed to create access token")
	}
	refreshToken, err := jwtx.GetToken(accessSecret, now, refreshExpire, uid)
	if err != nil {
		return "", "", errors.New("failed to create refresh token")
	}

	return accessToken, refreshToken, nil
}
