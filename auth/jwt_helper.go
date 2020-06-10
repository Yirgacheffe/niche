package main

import (
	"crypto/rsa"
	"io/ioutil"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
)

// Jwks - List all json web keys for application
type Jwks struct {
	Keys []JSONWebKey `json:"keys"`
}

// JSONWebKey - Public key to Json web key
type JSONWebKey struct {
	Kty string `json:"kty"`
	Kid string `json:"kid"`
	Use string `json:"use"`
	E   string `json:"e"`
	N   string `json:"n"`
	Alg string `json:"alg"`
}

// using asymmetric crypto/RSA keys, openssl genrsa
const (
	privKeyPath = "keys/niche.rsa"
	pubKeyPath  = "keys/niche.rsa.pub"
)

const (
	jwtIss = "nichesoft.io"
	jwtAud = "todox-app"
)

var (
	verifyKey *rsa.PublicKey
	signKey   *rsa.PrivateKey
)

func init() {
	var err error

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatalf("[initKeys]: %s\n", err)
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal("error")
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal("error")
	}
}

// GenerateJWT - Generate jwt token
func GenerateJWT(name, role string) (string, error) {

	claims := jwt.MapClaims{
		"iss":  jwtIss,
		"aud":  jwtAud,
		"exp":  time.Now().Add(time.Minute * 10).Unix(),
		"name": name,
		"role": role,
	}

	// create a signer for rsa 256
	t := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	tokenString, err := t.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}
