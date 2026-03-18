package libraries

import(
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"peregerine/configs"
)

type Token struct {
	AccessToken string
	RefreshToken string
}

func GenerateToken(userID string) (*Token, error) {
	ttl := configs.JWTTTL
	secret := configs.JWTSecret

	claims := jwt.MapClaims{
		"userId": userID,
		"time": time.Now().Add(time.Duration(ttl) * time.Minute).Unix(),
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := at.SignedString(secret)
	if err != nil {
		return nil, err
	}

	refreshClaims := jwt.MapClaims{
		"userId": userID,
		"time": time.Now().Add(time.Duration(ttl) * time.Minute).Unix(),
		"jti": uuid.NewString(),
	}

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := refresh.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &Token {
		AccessToken: accessToken,
		RefreshToken: refreshToken,
	}, nil
}