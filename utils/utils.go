package utils

import (
	"case-management/appcore/appcore_cache"
	"case-management/model"
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

// var taskProgress = make(map[string]int)
// var mu sync.Mutex

const bcryptCost = 12

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	return string(bytes), err
}

// func SSEProgress(c *gin.Context) {
// 	c.Writer.Header().Set("Content-Type", "text/event-stream")
// 	c.Writer.Header().Set("Cache-Control", "no-cache")
// 	c.Writer.Header().Set("Connection", "keep-alive")
// 	c.Writer.Flush()

// 	taskID := c.Query("taskID")

// 	for {
// 		progress := GetProgress(taskID)

// 		fmt.Fprintf(c.Writer, "data: %d\n\n", progress)
// 		c.Writer.Flush()

// 		if progress >= 100 {
// 			break
// 		}
// 		time.Sleep(1 * time.Second)
// 	}
// }

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

// func GetProgress(taskID string) int {
// 	ctx := context.Background()
// 	val, err := appcore_cache.Cache.Get(ctx, taskID).Result()

// 	if err != nil {
// 		return 0
// 	}
// 	i, err := strconv.Atoi(val)
// 	if err != nil {
// 		return 0
// 	}
// 	return i
// }

func ParseUint(s string) uint {
	val, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0
	}
	return uint(val)
}

func ParseUUID(s string) (uuid.UUID, error) {
	return uuid.Parse(strings.TrimSpace(s))
}

func SetImportStatus(taskID string, status model.ImportStatus) {
	ctx := context.Background()
	jsonData, err := json.Marshal(status)
	if err != nil {
		fmt.Printf("Error marshaling ImportStatus for taskID %s: %v\n", taskID, err)
		return
	}
	appcore_cache.Cache.Set(ctx, taskID, jsonData, 10*time.Minute)
}

func GetImportStatus(taskID string) model.ImportStatus {
	ctx := context.Background()
	val, err := appcore_cache.Cache.Get(ctx, taskID).Result()
	if err != nil {
		return model.ImportStatus{}
	}
	var status model.ImportStatus
	if err := json.Unmarshal([]byte(val), &status); err != nil {
		fmt.Printf("Error unmarshaling ImportStatus for taskID %s: %v\n", taskID, err)
		return model.ImportStatus{}
	}
	return status
}
