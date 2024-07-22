package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	cache := NewLRUCache(100)

	r.GET("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		value, found := cache.Get(key)
		if !found {
			c.JSON(http.StatusNotFound, gin.H{"message": "Key not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"value": value})
	})

	r.POST("/cache", func(c *gin.Context) {
		var json struct {
			Key        string `json:"key"`
			Value      string `json:"value"`
			Expiration int64  `json:"expiration"`
		}
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		cache.Set(json.Key, json.Value, time.Now().Unix()+json.Expiration)
		c.JSON(http.StatusOK, gin.H{"message": "Key set"})
	})

	r.DELETE("/cache/:key", func(c *gin.Context) {
		key := c.Param("key")
		cache.Delete(key)
		c.JSON(http.StatusOK, gin.H{"message": "Key deleted"})
	})

	go func() {
		for {
			time.Sleep(1 * time.Second)
			cache.mutex.Lock()
			for key, item := range cache.items {
				if item.Expiration <= time.Now().Unix() {
					delete(cache.items, key)
				}
			}
			cache.mutex.Unlock()
		}
	}()

	r.Run(":8088")
}
