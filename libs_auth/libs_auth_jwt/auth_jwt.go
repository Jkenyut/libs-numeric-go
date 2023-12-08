package libs_auth_jwt

import (
	"context"
	"github.com/Jkenyut/libs-numeric-go/libs_config"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_jwt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

//jwt

type ClientAuth struct {
	conf libs_config.JWTConfig
}

func NewClientAuthJWT(conf libs_config.JWTConfig) InterfacesAuthJWT {
	return &ClientAuth{
		conf: conf,
	}
}

func (repo *ClientAuth) GenerateJWTAccessCustom(ctx context.Context, issuer string, audience []string, subject string, id string, data any) (tokenJWTAccess string, claims libs_model_jwt.CustomClaims, err error) {
	claimsAccess := libs_model_jwt.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   subject,
			Audience:  audience,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(repo.conf.ExpiredAccess) * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        id,
		},
		Data: data,
	}

	// Create the token
	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)

	// Sign the token with the secret key
	tokenJWTAccess, err = tokenAccess.SignedString([]byte(repo.conf.Access))
	if err != nil {
		return tokenJWTAccess, libs_model_jwt.CustomClaims{}, err
	}
	return tokenJWTAccess, claimsAccess, nil
}

func (repo *ClientAuth) GenerateJWTRefreshCustom(ctx context.Context, issuer string, audience []string, subject string, id string, data any) (tokenJWTAccess string, claims libs_model_jwt.CustomClaims, err error) {
	claimsAccess := libs_model_jwt.CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			Subject:   subject,
			Audience:  audience,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(repo.conf.ExpiredAccess) * time.Minute)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ID:        id,
		}, Data: data,
	}

	// Create the token
	tokenAccess := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsAccess)

	// Sign the token with the secret key
	tokenJWTAccess, err = tokenAccess.SignedString([]byte(repo.conf.Refresh))
	if err != nil {
		return tokenJWTAccess, libs_model_jwt.CustomClaims{}, err
	}
	return tokenJWTAccess, claimsAccess, nil
}
