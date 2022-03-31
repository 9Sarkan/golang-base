package oauth

import (
	"context"
	"errors"
	"log"

	"github.com/9sarkan/golang-base/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var (
	JWT oautInterface = &AccessDetails{}
)

type oautInterface interface {
	VerifyToken(ctx context.Context, request string) (*AccessDetails, error)
}

// AccessDetails struct
type AccessDetails struct {
	UserID string
	Dialer string
}

func (a *AccessDetails) VerifyToken(ctx context.Context, request string) (*AccessDetails, error) {
	token, err := jwt.Parse(request, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(config.Map.AuthCreditional.HMACSecret), nil
		}
		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			return jwt.ParseRSAPublicKeyFromPEM([]byte(config.Map.AuthCreditional.RSASecret))
		}
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); ok {
			return jwt.ParseECPublicKeyFromPEM([]byte(config.Map.AuthCreditional.ECDSASecret))
		}

		log.Default().Printf("unexpected signing method: %v", token.Header["alg"])
		return nil, errors.New("401 unauthorized")
	})
	if err != nil {
		return nil, errors.New("401 unauthorized")
	}
	if token.Claims == nil || !token.Valid {
		return nil, errors.New("401 unauthorized")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, ok := claims["sub"].(string)
		if !ok {
			log.Default().Printf("error in get user sub from token: %v", claims)
			return nil, errors.New("401 unauthorized")
		}

		if _, err := uuid.Parse(userID); err != nil {
			log.Default().Printf("error in sub typo: %v", err)
			return nil, errors.New("401 unauthorized")
		}

		dialer, ok := claims["mobile"].(string)
		if !ok {
			log.Default().Printf("error in get user mobile from token: %v", claims)
			return nil, errors.New("401 unauthorized")
		}

		return &AccessDetails{
			UserID: userID,
			Dialer: dialer,
		}, nil
	}
	return nil, err
}
