package structs

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type (
	JwtClaims struct {
		jwt.RegisteredClaims
		UserID uuid.UUID `json:"user_id"`
	}

	Env struct {
		DB_URL     string
		JWT_SECRET []byte
	}
)
