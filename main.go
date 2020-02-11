package main

import(
	"math/rand"
	"time"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/Pallinder/go-randomdata"
)

func randStr(animals []string) string {
	return animals[rand.Intn(len(animals))]
}


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


func kittyHandler( c *gin.Context) {
	cats := []string{
		"/img/kitties/1.jpg",
		"/img/kitties/2.jpg",
		"/img/kitties/3.jpg",
		"/img/kitties/4.jpg",
	}


	htmlStr := fmt.Sprintf(`<html>
	<head></head>
	<body>
	<center>
	<img src="%s" alt="%s" width="1000">
	</center>
	</body>
	</html>`, randStr(cats), randomdata.SillyName())

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(htmlStr))
}


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
	<img src="%s" alt="%s" width="1000">
	</center>
	</body>
	</html>`, randStr(dogs), randomdata.SillyName())

	c.Writer.Header().Set("Content-Type", "text/html")
	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write([]byte(htmlStr))
}



func main() {
	rand.Seed(time.Now().UnixNano())
	router := gin.Default()

	router.Static("/img", "./img")
	router.GET("/", mainHandler)
	router.GET("/kitty", kittyHandler)
	router.GET("/pup", pupHandler)

	router.Run(":8085")
	// s := &http.Server{
	// 	Addr:    ":4700",
	// 	Handler: router,
	// }
	// s.ListenAndServe()
}
