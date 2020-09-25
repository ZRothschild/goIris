package middleware

import (
	"github.com/dgrijalva/jwt-go"
	jwtMiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/spf13/viper"
)

func Jwt(newViper *viper.Viper, keyStr string) *jwtMiddleware.Middleware {
	jwtKey := newViper.GetString(keyStr)
	return jwtMiddleware.New(jwtMiddleware.Config{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
		// ErrorHandler: AuthEchoError,
		// When set, the middleware verifies that tokens are signed with the specific signing algorithm
		// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
		// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
		SigningMethod: jwt.SigningMethodHS256,
	})
}

/***
default handler
*/
func AuthEchoError(ctx iris.Context, err error) {
	ctx.StatusCode(iris.StatusUnauthorized)
	_, _ = ctx.JSON(map[string]interface{}{"status": 2, "result": "", "message": err.Error()})
	return
}

/**
when not need login should use this handler
*/
func AuthNullError(ctx context.Context, err string) {
	return
}
