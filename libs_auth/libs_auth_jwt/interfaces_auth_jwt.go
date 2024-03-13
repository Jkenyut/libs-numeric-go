package libs_auth_jwt

import (
	"context"
	"github.com/Jkenyut/libs-numeric-go/libs_models/libs_model_jwt"
)

// InterfacesAuthJWT method NewClientAuthJWT
type InterfacesAuthJWT interface {
	GenerateJWTAccessCustom(ctx context.Context, issuer string, audience []string, subject string, id string, data any) (tokenJWTAccess string, claims *libs_model_jwt.CustomClaims, err error)
	GenerateJWTRefreshCustom(ctx context.Context, issuer string, audience []string, subject string, id string, data any) (tokenJWTAccess string, claims *libs_model_jwt.CustomClaims, err error)
}
