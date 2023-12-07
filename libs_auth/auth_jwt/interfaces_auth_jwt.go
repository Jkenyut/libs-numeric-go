package auth_jwt

import (
	"context"
	"github.com/Jkenyut/libs-numeric-go/libs_models/model_jwt"
)

type InterfacesAuthJWT interface {
	GenerateJWTAccessCustom(ctx context.Context, audience []string, activityId string, id string, data any) (tokenJWTAccess string, claims model_jwt.CustomClaims, err error)
	GenerateJWTRefreshCustom(ctx context.Context, audience []string, activityId string, id string, data any) (tokenJWTAccess string, claims model_jwt.CustomClaims, err error)
}
