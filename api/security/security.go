package security

import (
	"lab01/api/classes"
	"lab01/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("TeamM1croserv1ces2019..")
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GetToken(userName string) classes.ResponseSecurity {
	var creds Credentials
	var res classes.ResponseSecurity
	creds.Username = userName
	expirationTime := time.Now().Add(config.MINUTES_EXPIRE_TOKEN * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		res.Code = config.ERROR_CREATE_TOKEN
		res.Token = ""
		return res
	}
	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	/*http.SetCookie(w, &http.Cookie{
	      Name:    "token",
	      Value:   tokenString,
	      Expires: expirationTime,
	  })
	*/
	res.Code = config.SUSSES_CREATE_TOKEN
	res.Token = tokenString
	return res
}
func ValidateToken(tknStr string) classes.ResponseSecurity {
	var res classes.ResponseSecurity
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !tkn.Valid {
		res.Code = config.USER_NOT_AUTHORIZE
		return res
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			res.Code = config.USER_NOT_AUTHORIZE
			return res
		}
		// bad request
		res.Code = config.USER_NOT_AUTHORIZE
		return res
	}
	res.Code = config.OK_PROCESS
	println("USER TOKEN: " + claims.Username)
	return res
}
