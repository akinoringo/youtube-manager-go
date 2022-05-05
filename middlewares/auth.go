package middlewares

import (
	"context"
	"firebase.google.com/go/auth"
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"strings"
)

/**
*	HeaderからFirebaseIDトークンを検証
*/
func VerifyFirebaseIdToken(ctx echo.Context, auth *auth.Client) (*auth.Token, error) {
	headerAuth := ctx.Request().Header.Get("Authorization")
	token := strings.Replace(headerAuth, "Bearer ", "", 1)
	jwtToken, err := auth.VerifyIDToken(context.Background(), token)

	return jwtToken, err
}

/**
* Firebaseの認証（ログイン時の認証に使用）
*/
func FirebaseGuard() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, err := VerifyFirebaseIdToken(c, authClient)
			if err != nil {
				c.JSON(fasthttp.StatusUnauthorized, "Not Authorized")
			}

			c.Set("auth", jwtToken)
			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}

/**
* Firebaseの認証（非ログイン時の認証に使用）
*/
func FirebaseAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authClient := c.Get("firebase").(*auth.Client)
			jwtToken, _ := VerifyFirebaseIdToken(c, authClient)

			c.Set("auth", jwtToken)

			if err := next(c); err != nil {
				return err
			}

			return nil
		}
	}
}
