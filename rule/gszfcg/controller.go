// controller
package gszfcg

import (
	"github.com/gin-gonic/gin"
)

func GetInfoApi(c *gin.Context) {
	title := c.DefaultQuery("title", "")
	city := c.DefaultQuery("city", "0")
	sort := c.DefaultQuery("sort", "128")
	if title == "" {
		c.JSON(500, "标题不能为空！")
		return
	}
	params := Params{
		Title:       title,
		City:        city,
		BiddingSort: sort,
	}
	data, err := Crawls(&params)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, data)
}
