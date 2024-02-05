package config

import "github.com/golang-jwt/jwt/v5"

type JwtClaim struct {
	Email string
	jwt.RegisteredClaims
}
