package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func world(c *gin.Context) {
	err := os.Getenv("HELLO_MESSAGE")
	if err == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, err)
}

func query(c *gin.Context) {
	message := c.Query("message")
	if message == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
		}
	c.String(http.StatusOK, "%s", message)
}

func param(c *gin.Context) {
	name := c.Param("message")
	if name == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "%s", name)
}

func body(c *gin.Context) {
	body := c.PostForm("message")
	if body == ""{
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "%s", body)
}

func header(c *gin.Context) {
	header := c.GetHeader("message")
	if header == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "%s", header)
}

func cookie(c *gin.Context) {
	cookie, err := c.Cookie("message")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.String(http.StatusOK, "%s", cookie)
}

func health(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}

func allQueries(c *gin.Context) {

}

func ApplyRoutes(r *gin.Engine) {
	r.GET("/hello", world)
	r.GET("/repeat-my-query", query)
	r.GET("/repeat-my-param/:message", param)
	r.POST("/repeat-my-body", body)
	r.GET("/repeat-my-header", header)
	r.GET("/repeat-my-cookie", cookie)
	r.GET("/health", health)
	r.GET("/repeat-all-my-queries", allQueries)
}