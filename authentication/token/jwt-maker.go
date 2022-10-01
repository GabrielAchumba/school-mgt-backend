package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSecretKeySize = 32

type Maker interface {
	// CreateToken creates a new token for a specific username and duration
	CreateToken(userId, email string) (string, *Payload, error)
	CreateRefeshToken(userId, email string) (string, *Payload, error)

	// VerifyToken checks if the token is valid or not
	VerifyToken(token string) (*Payload, error)
	VerifyRefreshToken(token string) (*Payload, error)
}

// JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	tokenSecretKey        string
	refreshKokenSecretKey string
	duration              time.Duration
	refreshDuration       time.Duration
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(tokenSecretKey string, refreshTokenSecretKey string, duration time.Duration, refreshDuration time.Duration) (Maker, error) {
	if len(tokenSecretKey) < minSecretKeySize {
		return nil, fmt.Errorf("Invalid key size: must be at least %d characters", minSecretKeySize)
	}
	if len(refreshTokenSecretKey) < minSecretKeySize {
		return nil, fmt.Errorf("Invalid key size: must be at least %d characters", minSecretKeySize)
	}
	return &JWTMaker{tokenSecretKey, refreshTokenSecretKey, duration, refreshDuration}, nil
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateToken(userId, email string) (string, *Payload, error) {

	payload, err := NewPayload(userId, email, maker.duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.tokenSecretKey))
	return token, payload, err
}

// CreateToken creates a new token for a specific username and duration
func (maker *JWTMaker) CreateRefeshToken(userId, email string) (string, *Payload, error) {

	payload, err := NewPayload(userId, email, maker.duration)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.refreshKokenSecretKey))
	return token, payload, err
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errInvalidToken
		}
		return []byte(maker.tokenSecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errExpiredToken) {
			return nil, errExpiredToken
		}
		return nil, errInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errInvalidToken
	}

	return payload, nil
}

// VerifyToken checks if the token is valid or not
func (maker *JWTMaker) VerifyRefreshToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errInvalidToken
		}
		return []byte(maker.refreshKokenSecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errExpiredToken) {
			return nil, errExpiredToken
		}
		return nil, errInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, errInvalidToken
	}

	return payload, nil
}
