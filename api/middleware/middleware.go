package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)

        userId := session.Get("user_id")
        if userId == nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not logged in"})
            c.Abort()
            return
        }
        c.Set("user_id", userId)
        c.Next()
    }
}
