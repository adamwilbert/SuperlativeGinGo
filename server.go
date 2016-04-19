package main

import (
	"github.com/gin-gonic/gin"
	"net/http"

)




func main() {
    router := gin.Default()

    router.LoadHTMLGlob("templates/*")
    //router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
    router.GET("/main", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "title": "Main website",
        })
    })

    router.GET("/test", func(c *gin.Context) {
    	c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")

	})

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world\n")
	})

    router.Run(":8080")
}
