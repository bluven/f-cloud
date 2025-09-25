package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
}

type UserClaims struct {
	UID      uint   `json:"uid"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"isAdmin"`
	jwt.RegisteredClaims
}

type JWTAuth struct {
	Issuer       string
	AccessSecret string
	AccessExpire string

	accessSecret []byte
	accessExpire time.Duration
}

func (a *JWTAuth) Validate() error {
	if a.AccessSecret == "" {
		return errors.New("JWTAuth.AccessSecret is empty")
	}

	accessExpire, err := time.ParseDuration(a.AccessExpire)
	if err != nil {
		return err
	}

	if accessExpire == 0 {
		return errors.New("JWTAuth.AccessExpire is too short")
	}

	return nil
}

func (a *JWTAuth) GetAccessSecret() []byte {
	if a.accessSecret == nil {
		a.accessSecret = []byte(a.AccessSecret)
	}
	return a.accessSecret
}

func (a *JWTAuth) GetAccessExpire() time.Duration {
	if a.accessExpire == 0 {
		a.accessExpire, _ = time.ParseDuration(a.AccessExpire)
	}

	return a.accessExpire
}
