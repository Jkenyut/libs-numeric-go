package auth_jwt

import (
	"context"
	"errors"
	"github.com/Jkenyut/libs-numeric-go/libs_config"
	"github.com/Jkenyut/libs-numeric-go/libs_models/model_jwt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type ClientAuth struct {
	conf libs_config.JWTConfig
}

func NewClientAuthJWT(conf libs_config.JWTConfig) InterfacesAuthJWT {
	return &ClientAuth{
		conf: conf,
	}
}
func (repo *ClientAuth) GenerateJWTAccessCustom(ctx context.Context, audience []string, activityId string, id string, data any) (tokenJWTAccess string, claims model_jwt.CustomClaims, err error) {
	claimsAccess := model_jwt.CustomClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "login",
			Subject:   activityId,
			Audience:  audience,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(repo.conf.ExpiredAccess) * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        id,
		},
	}

	// Create the token
	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)

	// Sign the token with the secret key
	tokenJWTAccess, err = tokenAccess.SignedString([]byte(repo.conf.Access))
	if err != nil {
		return tokenJWTAccess, model_jwt.CustomClaims{}, errors.New(err.Error())
	}
	return tokenJWTAccess, claimsAccess, nil
}

func (repo *ClientAuth) GenerateJWTRefreshCustom(ctx context.Context, audience []string, activityId string, id string, data any) (tokenJWTAccess string, claims model_jwt.CustomClaims, err error) {
	claimsAccess := model_jwt.CustomClaims{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "login",
			Subject:   activityId,
			Audience:  audience,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(repo.conf.ExpiredAccess) * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        id,
		},
	}

	// Create the token
	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)

	// Sign the token with the secret key
	tokenJWTAccess, err = tokenAccess.SignedString([]byte(repo.conf.Refresh))
	if err != nil {
		return tokenJWTAccess, model_jwt.CustomClaims{}, errors.New(err.Error())
	}
	return tokenJWTAccess, claimsAccess, nil
}
