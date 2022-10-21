package jwt

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	appId  = "1873974"
	jwtIss = "nichesoft.io"
	jwtAud = appId
	jwtKid = "-HVXLi1zsLoESUVjWWJwPayM0p0TRhgBm0nBotvBm6s"

	privPath = "/etc/secret-keys/ssh-privatekey"
)

var privSecret []byte

func init() {
	if keyData, err := ioutil.ReadFile(privPath); err != nil {
		log.Fatal(err)
	} else {
		privSecret = keyData
	}
}

// GenerateJWT - Generate jwt token
func GenerateJWT(id int, name, email string) (string, error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM(privSecret)
	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	claims := jwt.MapClaims{
		"iss":   jwtIss,
		"aud":   jwtAud,
		"kid":   jwtKid,
		"exp":   time.Now().Add(time.Minute * 30).Unix(),
		"sub":   id,
		"email": email,
		"name":  name,
	}

	// create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := t.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil // jwt string witout error ......
}
