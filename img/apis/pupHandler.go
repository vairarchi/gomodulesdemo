package main

import(
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func pupHandler( c *gin.Context) {
	dogs := []string{
		"/img/pups/1.jpg",
		"/img/pups/2.jpg",
		"/img/pups/3.jpg",
		"/img/pups/4.jpg",
	}


	htmlStr := fmt.Sprintf(`<html>
	<head></head>
	<body>
	<center>
	<img src="%s" width="1000"
	</center>
	</body>
	</html>`, randStr(dogs))

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(htmlStr))
}