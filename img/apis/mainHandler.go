package main

import(
	"net/http"

	"github.com/gin-gonic/gin"
)

func mainHandler( c *gin.Context) {
	htmlStr := `<html>
	<head></head>
	<body>
	<center>
	<h1>Hey MKCL!</h1>
	<h3>Demo of Go Modules</h3>
	<p><a href="/kitty"> Cats </p>
	<p><a href="/pup"> Dogs </p>
	</center>
	</body>
	</html>`

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(htmlStr))
}