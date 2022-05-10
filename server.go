package main

import (
	"github.com/PanYuHaa/go-project-community/controller"
	"github.com/PanYuHaa/go-project-community/repository"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})

	r.POST("/community/post/do", func(c *gin.Context) {
		topicId, _ := c.GetPostForm("topic_id")
		content, _ := c.GetPostForm("content")
		data := cotroller.PublishPost(topicId, content)
		c.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
