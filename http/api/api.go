package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/CyrivlClth/snowserver/config"
)

func NextID(c *gin.Context) {
	id, err := config.IDGenerator().NextID()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": strconv.Itoa(int(id))})
	return
}

func Stats(c *gin.Context) {
	panic("implement this api")
}

func GetIDs(c *gin.Context) {
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil || count <= 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "param count must be integer and greater than 0"})
		return
	}
	ids := make([]string, 0)
	for i := 0; i < count; i++ {
		id, err := config.IDGenerator().NextID()
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ids = append(ids, strconv.Itoa(int(id)))
	}
	c.JSON(http.StatusOK, gin.H{"count": count, "ids": ids})
	return
}
