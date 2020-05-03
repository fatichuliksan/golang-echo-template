package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

type jwtHelper struct {
}

// Method ...
type Method interface {
	CreateJwtToken(secret string, ID string) (string, error)
	GetJwtClaims(c echo.Context) jwt.MapClaims
	GetJwtClaim(c echo.Context, key string) interface{}
}

// NewJWT ...
func NewJWT() Method {
	return &jwtHelper{}
}

// CreateJwtToken ...
func (t *jwtHelper) CreateJwtToken(secret string, ID string) (string, error) {
	claims := jwt.StandardClaims{}
	claims.Id = ID
	claims.ExpiresAt = time.Now().Add(30 * time.Minute).Unix()
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return token, err
}

func (t *jwtHelper) GetJwtClaims(c echo.Context) jwt.MapClaims {
	user := c.Get("ID")
	if user == nil {
		return nil
	}

	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	return claims
}

func (t *jwtHelper) GetJwtClaim(c echo.Context, key string) interface{} {
	claims := t.GetJwtClaims(c)

	return claims[key]
}
