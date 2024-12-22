package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (m *Middleware) JwtTokenCheck(c *gin.Context) {
	if token := strings.Split(c.GetHeader("Authorization"), " "); len(token) == 2 {
		JWT, err := m.parseToken(c, token[1])
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !JWT.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Next()
}

func (m *Middleware) parseToken(c *gin.Context, token string) (*jwt.Token, error) {
	Token, err := jwt.Parse(token,
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("bad sign method")
			}
			return []byte(m.cfg.Secret), nil
		})
	if err != nil {
		return nil, err
	}

	for k, v := range Token.Claims.(jwt.MapClaims) {
		c.Set(k, v)
	}

	return Token, nil
}
