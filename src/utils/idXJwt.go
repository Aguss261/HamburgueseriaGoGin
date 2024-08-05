package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func IdXJwt(c *gin.Context) int {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No user ID found in context"})
		return -1
	}

	userIDStr, ok := userID.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID type"})
		return -1
	}

	if userIDStr == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID is empty"})
		return -1
	}

	id, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return -1
	}
	return int(id)
}
