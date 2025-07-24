package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getLimit(c *gin.Context) (int, error) {
	limit := 10
	var err error

	limitParam := c.Query("limit")
	if limitParam != "" {
		limit, err = strconv.Atoi(limitParam)
		if err != nil {
			return 0, err
		}
	}

	return limit, nil
}

func getPage(c *gin.Context) (int, error) {
	page := 1
	var err error

	pageParam := c.Query("page")
	if pageParam != "" {
		page, err = strconv.Atoi(pageParam)
		if err != nil {
			return 0, err
		}
	}

	return page, nil
}
