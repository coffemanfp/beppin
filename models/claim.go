package models

import (
	"fmt"

	"github.com/coffemanfp/beppin-server/config"
	"github.com/dgrijalva/jwt-go"
)

// Claim - Token user data.
type Claim struct {
	User User `json:"user"`
	jwt.StandardClaims
}

// GenerateJWT - Generates a JSON Web Token.
func (c *Claim) GenerateJWT() (result string, err error) {
	settings, err := config.GetSettings()
	if err != nil {
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	result, err = token.SignedString([]byte(settings.SecretKey))
	if err != nil {
		err = fmt.Errorf("failed to sign token:\n%s", err)
	}
	return
}
