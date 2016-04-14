package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"time"
)


func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()

        // Set example variable
        c.Set("example", "12345")

        // before request

        c.Next()

        // after request
        latency := time.Since(t)
        log.Println("latency = ", latency)

        // access the status we are sending
        status := c.Writer.Status()
        log.Println("statuscode = ", status)
    }
}


func main() {
    router := gin.Default()
    router.Use(Logger())


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
