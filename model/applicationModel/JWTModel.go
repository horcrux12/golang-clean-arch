package applicationModel

import "github.com/golang-jwt/jwt/v4"

type MyClaims struct {
	*jwt.StandardClaims
}

type JWTDetailAccount struct {
	ID int64
}
