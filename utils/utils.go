package utils

import (
	"case-management/appcore/appcore_cache"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// var taskProgress = make(map[string]int)
// var mu sync.Mutex

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func SSEProgress(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Flush()

	taskID := c.Query("taskID")

	for {
		progress := GetProgress(taskID)

		fmt.Fprintf(c.Writer, "data: %d\n\n", progress)
		c.Writer.Flush()

		if progress >= 100 {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

// func SetProgress(taskID string, value int) {
// 	mu.Lock()
// 	taskProgress[taskID] = value
// 	mu.Unlock()
// }

// func GetProgress(taskID string) int {
// 	mu.Lock()
// 	defer mu.Unlock()
// 	return taskProgress[taskID]
// }

func SetProgress(taskID string, value int) {
	ctx := context.Background()
	appcore_cache.Cache.Set(ctx, taskID, value, 10*time.Minute)
}

func GetProgress(taskID string) int {
	ctx := context.Background()
	val, err := appcore_cache.Cache.Get(ctx, taskID).Result()
	if err != nil {
		return 0
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return i
}

func ParseUint(s string) uint {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint(val)
}
