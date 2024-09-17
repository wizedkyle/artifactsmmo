package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func QueryLimit(c *gin.Context) int64 {
	limit := c.DefaultQuery("limit", "50")
	limitInt64, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return 0
	}
	return limitInt64
}

func QueryLevel(c *gin.Context) int {
	level := c.DefaultQuery("level", "0")
	levelInt, err := strconv.Atoi(level)
	if err != nil {
		return 0
	}
	return levelInt
}
