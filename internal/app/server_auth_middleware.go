package app

import (
	"errors"
	"net/http"
	"strings"

	"github.com/assetto-corsa-web/accweb/internal/pkg/cfg"
	"github.com/gin-gonic/gin"
)

var ErrForbidden = errors.New("access denied")

type ACCWebAuthLevel int

type StaticAuthHandler struct {
	Config *cfg.Config
}

const (
	ACCWebAuthLevel_Mod ACCWebAuthLevel = iota
	ACCWebAuthLevel_Adm
)

func ACCWebAuthMiddleware(lvl ACCWebAuthLevel) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := GetUserFromClaims(c)

		if lvl == ACCWebAuthLevel_Mod && (!u.Mod && !u.Admin) {
			c.JSON(http.StatusForbidden, gin.H{"msg": ErrForbidden})
			c.Abort()
			return
		}

		if lvl == ACCWebAuthLevel_Adm && !u.Admin {
			c.JSON(http.StatusForbidden, gin.H{"msg": ErrForbidden})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (h *StaticAuthHandler) AuthenticateM2M() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from client request
		authHeader := strings.Split(c.Request.Header.Get("Authorization"), "Bearer ")
		if len(authHeader) != 2 {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}
		token := authHeader[1]

		// Check token
		if token != h.Config.Auth.StaticToken {
			c.Status(http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()
	}
}
