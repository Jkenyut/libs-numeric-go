package libs_model_jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Data any `json:"data,omitempty"`
}
